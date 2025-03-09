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
	UserCache           = "User_Cache"
	UserCacheExpireTime = time.Hour * 2 // Access Token 有效期
)

// GetAccount 获取用户信息逻辑
func GetAccount(req *dto.GetAccountRequest, c echo.Context) (*account.GetAccountVo, error) {
	userInfo, err := mapper.GetAccountByEmail(req.Email)
	if err != nil {
		utils.BizLogger(c).Errorf("「%s」邮箱不存在", req.Email)
		return nil, fmt.Errorf("「%s」邮箱不存在", req.Email)
	}

	vo, err := utils.MapModelToVO(userInfo, &account.GetAccountVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("获取用户信息时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("获取用户信息时映射 vo 失败: %v", err)
	}

	return vo.(*account.GetAccountVo), nil
}

// RegisterAcc 用户注册逻辑
func RegisterAcc(req *dto.RegisterRequest, c echo.Context) (*account.RegisterAccountVo, error) {
	registerLock.Lock()
	defer registerLock.Unlock()

	totalAccounts, err := mapper.GetTotalAccounts()
	if err != nil {
		utils.BizLogger(c).Errorf("获取用户总数失败: %v", err)
		return nil, fmt.Errorf("获取用户总数失败: %v", err)
	}

	if totalAccounts > 0 {
		utils.BizLogger(c).Error("系统限制: 当前为单用户独立部署版本，已达到账户数量上限 (1/1)")
		return nil, fmt.Errorf("系统限制: 当前为单用户独立部署版本，已达到账户数量上限 (1/1)")
	}

	existingUser, _ := mapper.GetAccountByEmail(req.Email)
	if existingUser != nil {
		utils.BizLogger(c).Errorf("「%s」邮箱已被注册", req.Email)
		return nil, fmt.Errorf("「%s」邮箱已被注册", req.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.BizLogger(c).Errorf("哈希加密失败: %v", err)
		return nil, fmt.Errorf("哈希加密失败: %v", err)
	}

	acc := &model.Account{
		Email:    req.Email,
		Password: string(hashedPassword),
		Nickname: req.Nickname,
		Phone:    req.Phone,
	}

	if err := mapper.CreateAccount(acc); err != nil {
		utils.BizLogger(c).Errorf("「%s」用户注册失败: %v", req.Email, err)
		return nil, fmt.Errorf("「%s」用户注册失败: %v", req.Email, err)
	}

	vo, err := utils.MapModelToVO(acc, &account.RegisterAccountVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("用户注册时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("用户注册时映射 vo 失败: %v", err)
	}

	return vo.(*account.RegisterAccountVo), nil
}

// LoginAcc 登录用户逻辑
func LoginAcc(req *dto.LoginRequest, c echo.Context) (*account.LoginVo, error) {
	acc, err := mapper.GetAccountByEmail(req.Email)
	if err != nil {
		utils.BizLogger(c).Errorf("「%s」用户不存在: %v", req.Email, err)
		return nil, fmt.Errorf("「%s」用户不存在: %v", req.Email, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(req.Password))
	if err != nil {
		utils.BizLogger(c).Errorf("密码输入错误: %v", err)
		return nil, fmt.Errorf("密码输入错误: %v", err)
	}

	accessTokenString, refreshTokenString, err := utils.GenerateJWT(acc.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("token 生成失败: %v", err)
		return nil, fmt.Errorf("token 生成失败: %v", err)
	}

	cacheKey := fmt.Sprintf("%s:%d", UserCache, acc.ID)

	err = global.RedisClient.Set(context.Background(), cacheKey, accessTokenString, UserCacheExpireTime).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("登录时设置缓存失败: %v", err)
		return nil, fmt.Errorf("登录时设置缓存失败: %v", err)
	}

	token := &account.LoginVo{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}

	vo, err := utils.MapModelToVO(token, &account.LoginVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("用户登录时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("用户登陆时映射 vo 失败: %v", err)
	}

	return vo.(*account.LoginVo), nil
}

// LogoutAcc 处理用户登出逻辑
func LogoutAcc(c echo.Context) error {
	logoutLock.Lock()
	defer logoutLock.Unlock()

	accountID, err := utils.ParseAccountAndRoleIDFromJWT(c.Request().Header.Get("Authorization"))
	if err != nil {
		utils.BizLogger(c).Errorf("解析 access token 失败: %v", err)
		return fmt.Errorf("解析 access token 失败: %v", err)
	}

	cacheKey := fmt.Sprintf("%s:%d", UserCache, accountID)
	err = global.RedisClient.Del(c.Request().Context(), cacheKey).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("删除 Redis 缓存失败: %v", err)
		return fmt.Errorf("删除 Redis 缓存失败: %v", err)
	}

	return nil
}

// ResetPassword 重置密码逻辑
func ResetPassword(req *dto.ResetPwdRequest, c echo.Context) error {
	passwordResetLock.Lock()
	defer passwordResetLock.Unlock()

	if req.NewPassword != req.AgainNewPassword {
		utils.BizLogger(c).Errorf("两次密码输入不一致")
		return fmt.Errorf("两次密码输入不一致")
	}

	accountID, err := utils.ParseAccountAndRoleIDFromJWT(c.Request().Header.Get("Authorization"))
	if err != nil {
		utils.BizLogger(c).Errorf("解析 token 失败: %v", err)
		return fmt.Errorf("解析 token 失败: %v", err)
	}

	acc, err := mapper.GetAccountByAccountID(accountID)
	if err != nil {
		utils.BizLogger(c).Errorf("「%s」用户不存在: %v", req.Email, err)
		return fmt.Errorf("「%s」用户不存在: %v", req.Email, err)
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.BizLogger(c).Errorf("密码加密失败: %v", err)
		return fmt.Errorf("密码加密失败: %v", err)
	}
	acc.Password = string(newPassword)

	if err := mapper.UpdateAccount(acc); err != nil {
		utils.BizLogger(c).Errorf("密码修改失败: %v", err)
		return fmt.Errorf("密码修改失败: %v", err)
	}

	return nil
}
