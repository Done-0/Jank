package service

import (
	"context"
	"fmt"
	"strconv"
	"sync"

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
	AccAuthTokenCachePrefix     = "ACC_AUTH_TOKEN_CACHE_PREFIX"
	RefreshAuthTokenCachePrefix = "REFRESH_AUTH_TOKEN_CACHE_PREFIX"
)

// GetAccount 获取用户信息逻辑
func GetAccount(req *dto.GetAccountRequest, c echo.Context) (*account.GetAccountVo, error) {
	userInfo, err := mapper.GetAccountByEmail(req.Email)
	if err != nil {
		utils.BizLogger(c).Errorf("邮箱(%s)不存在", req.Email)
		return nil, fmt.Errorf("邮箱不存在")
	}

	vo, err := utils.MapModelToVO(userInfo, &account.GetAccountVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("获取用户信息时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("获取用户信息时映射 vo 失败: %v", err)
	}

	return vo.(*account.GetAccountVo), nil
}

// RegisterUser 用户注册逻辑
func RegisterUser(req *dto.RegisterRequest, c echo.Context) (*account.RegisterAccountVo, error) {
	registerLock.Lock()
	defer registerLock.Unlock()

	existingUser, _ := mapper.GetAccountByEmail(req.Email)
	if existingUser != nil {
		utils.BizLogger(c).Errorf("邮箱已被注册: %v", req.Email)
		return nil, fmt.Errorf("邮箱已被注册")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.BizLogger(c).Errorf("密码加密失败: %v", err)
		return nil, fmt.Errorf("密码加密失败: %v", err)
	}

	acc := &model.Account{
		Email:    req.Email,
		Password: string(hashedPassword),
		Nickname: req.Nickname,
		Phone:    req.Phone,
	}

	if err := mapper.CreateAccount(acc); err != nil {
		utils.BizLogger(c).Errorf("用户注册失败: %v", err)
		return nil, fmt.Errorf("用户注册失败: %v", err)
	}

	// 获取并分配默认角色，如果没有则自动创建
	role, err := mapper.GetRoleByCode("user")
	if err != nil {
		defaultRole := &model.Role{
			Code:        "user",
			Description: "普通用户",
		}
		if err := mapper.CreateRole(defaultRole); err != nil {
			utils.BizLogger(c).Errorf("创建默认角色失败: %v", err)
			return nil, fmt.Errorf("创建默认角色失败: %v", err)
		}
		role = defaultRole
	}

	if err := mapper.AssignRoleToUser(acc.ID, role.ID); err != nil {
		utils.BizLogger(c).Errorf("给用户分配角色失败: %v", err)
		return nil, fmt.Errorf("给用户分配角色失败: %v", err)
	}

	vo, err := utils.MapModelToVO(acc, &account.RegisterAccountVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("用户注册时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("用户注册时映射 vo 失败: %v", err)
	}

	return vo.(*account.RegisterAccountVo), nil
}

// LoginUser 登录用户逻辑
func LoginUser(req *dto.LoginRequest, c echo.Context) (*account.LoginVo, error) {
	user, err := mapper.GetAccountByEmail(req.Email)
	if err != nil {
		utils.BizLogger(c).Errorf("用户不存在: %v", err)
		return nil, fmt.Errorf("用户不存在: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		utils.BizLogger(c).Errorf("密码错误: %v", err)
		return nil, fmt.Errorf("密码错误: %v", err)
	}

	accessTokenString, refreshTokenString, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		utils.BizLogger(c).Errorf("生成 token 失败: %v", err)
		return nil, fmt.Errorf("生成 token 失败: %v", err)
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

// LogoutUser 刷新 token 逻辑
func LogoutUser(userId int64, c echo.Context) error {
	logoutLock.Lock()
	defer logoutLock.Unlock()

	accKey := AccAuthTokenCachePrefix + strconv.FormatInt(userId, 10)
	refreshKey := RefreshAuthTokenCachePrefix + strconv.FormatInt(userId, 10)

	ctx := context.Background()

	go func() {
		cmd := global.RedisClient.Do(ctx, global.DelCmd, accKey)
		if cmd.Err() != nil {
			utils.BizLogger(c).Errorf("删除鉴权 token 缓存失败: %v", cmd.Err())
		}
	}()

	go func() {
		cmd := global.RedisClient.Do(ctx, global.DelCmd, refreshKey)
		if cmd.Err() != nil {
			utils.BizLogger(c).Errorf("删除刷新 token 缓存失败: %v", cmd.Err())
		}
	}()

	return nil
}

// ResetPassword 重置密码逻辑
func ResetPassword(req *dto.ResetPwdRequest, userId int64, c echo.Context) error {
	passwordResetLock.Lock()
	defer passwordResetLock.Unlock()

	if req.NewPassword != req.AgainNewPassword {
		utils.BizLogger(c).Errorf("两次密码输入不一致")
		return fmt.Errorf("两次密码输入不一致")
	}

	acc, err := mapper.GetAccountByUserID(userId)
	if err != nil {
		utils.BizLogger(c).Errorf("用户不存在: %v", err)
		return fmt.Errorf("用户不存在: %v", err)
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.BizLogger(c).Errorf("密码加密失败")
		return fmt.Errorf("密码加密失败")
	}
	acc.Password = string(newPassword)

	if err := mapper.UpdateAccount(acc); err != nil {
		utils.BizLogger(c).Errorf("密码修改失败: %v", err)
		return fmt.Errorf("密码修改失败: %v", err)
	}

	go func() {
		global.BizLog.Infof("用户密码已重置: %s", acc.Email)
	}()

	return nil
}

// CreateRole 创建角色
func CreateRole(req *dto.CreateRoleRequest, c echo.Context) (*model.Role, error) {
	role := &model.Role{
		Code:        req.Code,
		Description: req.Description,
	}

	if err := mapper.CreateRole(role); err != nil {
		utils.BizLogger(c).Errorf("创建角色失败: %v", err)
		return nil, fmt.Errorf("创建角色失败: %v", err)
	}

	vo, err := utils.MapModelToVO(role, &account.RoleVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("角色创建时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("角色创建时映射 vo 失败: %v", err)
	}

	return vo.(*model.Role), nil
}

// UpdateRole 更新角色
func UpdateRole(req *dto.UpdateRoleRequest, c echo.Context) (*model.Role, error) {
	role, err := mapper.GetRoleByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取角色失败: %v", err)
		return nil, fmt.Errorf("获取角色失败: %v", err)
	}

	role.Code = req.Code
	role.Description = req.Description

	if err := mapper.UpdateRole(role); err != nil {
		utils.BizLogger(c).Errorf("更新角色失败: %v", err)
		return nil, fmt.Errorf("更新角色失败: %v", err)
	}

	vo, err := utils.MapModelToVO(role, &account.RoleVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("角色更新时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("角色更新时映射 vo 失败: %v", err)
	}

	return vo.(*model.Role), nil
}

// DeleteRole 删除角色
func DeleteRole(req *dto.DeleteRoleRequest, c echo.Context) error {
	if err := mapper.DeleteRole(req.ID); err != nil {
		utils.BizLogger(c).Errorf("删除角色失败: %v", err)
		return fmt.Errorf("删除角色失败: %v", err)
	}

	return nil
}

// ListRoles 获取所有角色
func ListRoles(c echo.Context) ([]*model.Role, error) {
	roles, err := mapper.GetAllRoles()
	if err != nil {
		utils.BizLogger(c).Errorf("获取所有角色失败: %v", err)
		return nil, fmt.Errorf("获取所有角色失败: %v", err)
	}

	return roles, nil
}

// CreatePermission 创建权限
func CreatePermission(req *dto.CreatePermissionRequest, c echo.Context) (*model.Permission, error) {
	permission := &model.Permission{
		Code:        req.Code,
		Description: req.Description,
	}

	if err := mapper.CreatePermission(permission); err != nil {
		utils.BizLogger(c).Errorf("创建权限失败: %v", err)
		return nil, fmt.Errorf("创建权限失败: %v", err)
	}

	vo, err := utils.MapModelToVO(permission, &account.PermissionVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("权限创建时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("权限创建时映射 vo 失败: %v", err)
	}

	return vo.(*model.Permission), nil
}

// UpdatePermission 更新权限
func UpdatePermission(req *dto.UpdatePermissionRequest, c echo.Context) (*model.Permission, error) {
	permission, err := mapper.GetPermissionByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取权限失败: %v", err)
		return nil, fmt.Errorf("获取权限失败: %v", err)
	}

	permission.Code = req.Code
	permission.Description = req.Description

	// 更新权限
	if err := mapper.UpdatePermission(permission); err != nil {
		utils.BizLogger(c).Errorf("更新权限失败: %v", err)
		return nil, fmt.Errorf("更新权限失败: %v", err)
	}

	vo, err := utils.MapModelToVO(permission, &account.PermissionVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("权限更新时映射 vo 失败: %v", err)
		return nil, fmt.Errorf("权限更新时映射 vo 失败: %v", err)
	}

	return vo.(*model.Permission), nil
}

// DeletePermission 删除权限
func DeletePermission(req *dto.DeletePermissionRequest, c echo.Context) error {
	if err := mapper.DeletePermission(req.ID); err != nil {
		utils.BizLogger(c).Errorf("删除权限失败: %v", err)
		return fmt.Errorf("删除权限失败: %v", err)
	}

	return nil
}

// ListPermissions 获取所有权限
func ListPermissions(c echo.Context) ([]*model.Permission, error) {
	permissions, err := mapper.GetAllPermissions()
	if err != nil {
		utils.BizLogger(c).Errorf("获取所有权限失败: %v", err)
		return nil, fmt.Errorf("获取所有权限失败: %v", err)
	}

	return permissions, nil
}

// AssignRoleToUser 为用户分配角色
func AssignRoleToUser(req *dto.AssignRoleRequest, c echo.Context) error {
	if err := mapper.AssignRoleToUser(req.RoleID, req.PermissionID); err != nil {
		utils.BizLogger(c).Errorf("为用户分配角色失败: %v", err)
		return fmt.Errorf("为用户分配角色失败: %v", err)
	}

	return nil
}

// AssignPermissionToRole 为角色分配权限
func AssignPermissionToRole(req *dto.AssignPermissionRequest, c echo.Context) error {
	if err := mapper.AssignPermissionToRole(req.RoleID, req.PermissionID); err != nil {
		utils.BizLogger(c).Errorf("为角色分配权限失败: %v", err)
		return fmt.Errorf("为角色分配权限失败: %v", err)
	}

	return nil
}
