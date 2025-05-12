// Package service 提供业务逻辑处理，处理账户相关业务
// 创建者：Done-0
// 创建时间：2025-05-10
package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"jank.com/jank_blog/internal/global"
	model "jank.com/jank_blog/internal/model/account"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/account/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/account"
)

var (
	registerLock      sync.Mutex // 用户注册锁，保护并发用户注册的操作
	passwordResetLock sync.Mutex // 修改密码锁，保护并发修改用户密码的操作
	logoutLock        sync.Mutex // 用户登出锁，保护并发用户登出操作
)

const (
	USER_CACHE             = "USER_CACHE"
	USER_CACHE_EXPIRE_TIME = time.Hour * 2 // Access Token 有效期
)

// GetAccount 获取用户信息逻辑
// 参数：
//   - c: Echo 上下文
//   - req: 获取账户请求
//
// 返回值：
//   - *account.GetAccountVO: 用户账户视图对象
//   - error: 操作过程中的错误
func GetAccount(c echo.Context, req *dto.GetAccountRequest) (*account.GetAccountVO, error) {
	userInfo, err := mapper.GetAccountByEmail(c, req.Email)
	if err != nil {
		utils.BizLogger(c).Errorf("「%s」邮箱不存在", req.Email)
		return nil, fmt.Errorf("「%s」邮箱不存在", req.Email)
	}

	vo, err := utils.MapModelToVO(userInfo, &account.GetAccountVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("获取用户信息时映射 VO 失败: %v", err)
		return nil, fmt.Errorf("获取用户信息时映射 VO 失败: %w", err)
	}

	return vo.(*account.GetAccountVO), nil
}

// RegisterAcc 用户注册逻辑
// 参数：
//   - c: Echo 上下文
//   - req: 注册账户请求
//
// 返回值：
//   - *account.RegisterAccountVO: 注册后的账户视图对象
//   - error: 操作过程中的错误
func RegisterAcc(c echo.Context, req *dto.RegisterRequest) (*account.RegisterAccountVO, error) {
	registerLock.Lock()
	defer registerLock.Unlock()

	var registerVO *account.RegisterAccountVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		totalAccounts, err := mapper.GetTotalAccounts(c)
		if err != nil {
			utils.BizLogger(c).Errorf("获取用户总数失败: %v", err)
			return fmt.Errorf("获取用户总数失败: %w", err)
		}

		if totalAccounts > 0 {
			utils.BizLogger(c).Error("系统限制: 当前为单用户独立部署版本，已达到账户数量上限 (1/1)")
			return fmt.Errorf("系统限制: 当前为单用户独立部署版本，已达到账户数量上限 (1/1)")
		}

		existingUser, _ := mapper.GetAccountByEmail(c, req.Email)
		if existingUser != nil {
			utils.BizLogger(c).Errorf("「%s」邮箱已被注册", req.Email)
			return fmt.Errorf("「%s」邮箱已被注册", req.Email)
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.BizLogger(c).Errorf("哈希加密失败: %v", err)
			return fmt.Errorf("哈希加密失败: %w", err)
		}

		acc := &model.Account{
			Email:    req.Email,
			Password: string(hashedPassword),
			Nickname: req.Nickname,
			Phone:    req.Phone,
		}

		if err := mapper.CreateAccount(c, acc); err != nil {
			utils.BizLogger(c).Errorf("「%s」用户注册失败: %v", req.Email, err)
			return fmt.Errorf("「%s」用户注册失败: %w", req.Email, err)
		}

		vo, err := utils.MapModelToVO(acc, &account.RegisterAccountVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("用户注册时映射 VO 失败: %v", err)
			return fmt.Errorf("用户注册时映射 VO 失败: %w", err)
		}

		registerVO = vo.(*account.RegisterAccountVO)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return registerVO, nil
}

// LoginAcc 登录用户逻辑
// 参数：
//   - c: Echo 上下文
//   - req: 登录请求
//
// 返回值：
//   - *account.LoginVO: 登录成功后的令牌视图对象
//   - error: 操作过程中的错误
func LoginAcc(c echo.Context, req *dto.LoginRequest) (*account.LoginVO, error) {
	acc, err := mapper.GetAccountByEmail(c, req.Email)
	if err != nil {
		utils.BizLogger(c).Errorf("「%s」用户不存在: %v", req.Email, err)
		return nil, fmt.Errorf("「%s」用户不存在: %w", req.Email, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(req.Password))
	if err != nil {
		utils.BizLogger(c).Errorf("密码输入错误: %v", err)
		return nil, fmt.Errorf("密码输入错误: %w", err)
	}

	accessTokenString, refreshTokenString, err := utils.GenerateJWT(acc.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("token 生成失败: %v", err)
		return nil, fmt.Errorf("token 生成失败: %w", err)
	}

	cacheKey := fmt.Sprintf("%s:%d", USER_CACHE, acc.ID)

	err = global.RedisClient.Set(context.Background(), cacheKey, accessTokenString, USER_CACHE_EXPIRE_TIME).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("登录时设置缓存失败: %v", err)
		return nil, fmt.Errorf("登录时设置缓存失败: %w", err)
	}

	token := &account.LoginVO{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	vo, err := utils.MapModelToVO(token, &account.LoginVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("用户登录时映射 VO 失败: %v", err)
		return nil, fmt.Errorf("用户登陆时映射 VO 失败: %v", err)
	}

	return vo.(*account.LoginVO), nil
}

// LogoutAcc 处理用户登出逻辑
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - error: 操作过程中的错误
func LogoutAcc(c echo.Context) error {
	logoutLock.Lock()
	defer logoutLock.Unlock()

	accountID, err := utils.ParseAccountAndRoleIDFromJWT(c.Request().Header.Get("Authorization"))
	if err != nil {
		utils.BizLogger(c).Errorf("解析 access token 失败: %v", err)
		return fmt.Errorf("解析 access token 失败: %w", err)
	}

	cacheKey := fmt.Sprintf("%s:%d", USER_CACHE, accountID)
	err = global.RedisClient.Del(c.Request().Context(), cacheKey).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("删除 Redis 缓存失败: %v", err)
		return fmt.Errorf("删除 Redis 缓存失败: %w", err)
	}

	return nil
}

// ResetPassword 重置密码逻辑
// 参数：
//   - c: Echo 上下文
//   - req: 重置密码请求
//
// 返回值：
//   - error: 操作过程中的错误
func ResetPassword(c echo.Context, req *dto.ResetPwdRequest) error {
	passwordResetLock.Lock()
	defer passwordResetLock.Unlock()

	return utils.RunDBTransaction(c, func(tx error) error {
		if req.NewPassword != req.AgainNewPassword {
			utils.BizLogger(c).Errorf("两次密码输入不一致")
			return fmt.Errorf("两次密码输入不一致")
		}

		accountID, err := utils.ParseAccountAndRoleIDFromJWT(c.Request().Header.Get("Authorization"))
		if err != nil {
			utils.BizLogger(c).Errorf("解析 token 失败: %v", err)
			return fmt.Errorf("解析 token 失败: %w", err)
		}

		acc, err := mapper.GetAccountByAccountID(c, accountID)
		if err != nil {
			utils.BizLogger(c).Errorf("「%s」用户不存在: %v", req.Email, err)
			return fmt.Errorf("「%s」用户不存在: %w", req.Email, err)
		}

		newPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			utils.BizLogger(c).Errorf("密码加密失败: %v", err)
			return fmt.Errorf("密码加密失败: %w", err)
		}
		acc.Password = string(newPassword)

		if err := mapper.UpdateAccount(c, acc); err != nil {
			utils.BizLogger(c).Errorf("密码修改失败: %v", err)
			return fmt.Errorf("密码修改失败: %w", err)
		}

		return nil
	})
}
