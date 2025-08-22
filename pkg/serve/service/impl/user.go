// Package impl 用户服务实现
// 创建者：Done-0
// 创建时间：2025-08-10
package impl

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/model/user"
	"github.com/Done-0/jank/internal/types/consts"
	"github.com/Done-0/jank/internal/utils/logger"
	"github.com/Done-0/jank/internal/utils/verification"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/mapper"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"
)

var (
	registerLock      sync.Mutex // 用户注册锁，保护并发用户注册的操作
	passwordResetLock sync.Mutex // 修改密码锁，保护并发修改用户密码的操作
	logoutLock        sync.Mutex // 用户登出锁，保护并发用户登出操作
)

// UserServiceImpl 用户服务实现
type UserServiceImpl struct {
	userMapper mapper.UserMapper
	rbacMapper mapper.RBACMapper
}

// NewUserService 创建用户服务实例
func NewUserService(userMapperImpl mapper.UserMapper, rbacMapperImpl mapper.RBACMapper) service.UserService {
	return &UserServiceImpl{
		userMapper: userMapperImpl,
		rbacMapper: rbacMapperImpl,
	}
}

// Register 用户注册逻辑
func (us *UserServiceImpl) Register(c *app.RequestContext, req *dto.RegisterRequest) (*vo.RegisterResponse, error) {
	registerLock.Lock()
	defer registerLock.Unlock()

	// 检查是否允许注册
	cfgs, err := configs.GetConfig()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get config: %v", err)
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	if !cfgs.AppConfig.User.AllowRegister {
		logger.BizLogger(c).Warn("registration attempt blocked: registration is disabled")
		return nil, fmt.Errorf("registration is currently disabled")
	}

	if !verification.VerifyEmailCode(c, req.EmailVerificationCode, req.Email) {
		return nil, fmt.Errorf("invalid email verification code")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.BizLogger(c).Errorf("password hashing failed: %v", err)
		return nil, fmt.Errorf("password hashing failed: %w", err)
	}

	u := &user.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Nickname: req.Nickname,
		Role:     cfgs.AppConfig.User.DefaultRole,
	}

	if err := us.userMapper.RegisterUser(c, u); err != nil {
		logger.BizLogger(c).Errorf("user registration failed for '%s': %v", req.Email, err)
		return nil, fmt.Errorf("user registration failed: %w", err)
	}

	userIDStr := strconv.FormatInt(u.ID, 10)
	if _, err := us.rbacMapper.AssignRole(c, userIDStr, cfgs.AppConfig.User.DefaultRole); err != nil {
		us.userMapper.DeleteUser(c, userIDStr)
		return nil, fmt.Errorf("user registration failed due to RBAC system error: %w", err)
	}

	roles, err := us.rbacMapper.GetUserRoles(c, userIDStr)
	userRoles := make([]string, 0)
	if err == nil {
		for _, role := range roles {
			userRoles = append(userRoles, role.V1)
		}
	}

	return &vo.RegisterResponse{
		ID:       strconv.FormatInt(u.ID, 10),
		Email:    u.Email,
		Nickname: u.Nickname,
		Roles:    userRoles,
	}, nil
}

// Login 登录用户逻辑
func (us *UserServiceImpl) Login(c *app.RequestContext, req *dto.LoginRequest) (*vo.LoginResponse, error) {
	u, err := us.userMapper.GetUserByEmail(c, req.Email)
	if err != nil {
		logger.BizLogger(c).Errorf("user not found for email '%s': %v", req.Email, err)
		return nil, fmt.Errorf("user not found: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		logger.BizLogger(c).Errorf("password verification failed for user '%s': %v", u.Email, err)
		return nil, fmt.Errorf("invalid password")
	}

	cfgs, err := configs.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	now := time.Now()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		consts.JWTSubjectClaim: u.ID,
		"exp":                  now.Add(time.Duration(cfgs.AppConfig.JWT.ExpireTime) * time.Hour).Unix(),
		"iat":                  now.Unix(),
	})
	accessTokenStr, err := accessToken.SignedString([]byte(cfgs.AppConfig.JWT.Secret))
	if err != nil {
		logger.BizLogger(c).Errorf("failed to sign access token for user %d: %v", u.ID, err)
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		consts.JWTSubjectClaim: u.ID,
		"exp":                  now.Add(time.Duration(cfgs.AppConfig.JWT.RefreshExpire) * time.Hour).Unix(),
		"iat":                  now.Unix(),
	})
	refreshTokenStr, err := refreshToken.SignedString([]byte(cfgs.AppConfig.JWT.Secret))
	if err != nil {
		logger.BizLogger(c).Errorf("failed to sign refresh token for user %d: %v", u.ID, err)
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	err = global.RedisClient.Set(context.Background(), fmt.Sprintf("%s:%d", consts.AuthAccessTokenKeyPrefix, u.ID), accessTokenStr, time.Duration(cfgs.AppConfig.JWT.ExpireTime)*time.Hour).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to cache access token for user %d: %v", u.ID, err)
		return nil, fmt.Errorf("failed to cache access token: %w", err)
	}

	err = global.RedisClient.Set(context.Background(), fmt.Sprintf("%s:%d", consts.AuthRefreshTokenKeyPrefix, u.ID), refreshTokenStr, time.Duration(cfgs.AppConfig.JWT.RefreshExpire)*time.Hour).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to cache refresh token for user %d: %v", u.ID, err)
		return nil, fmt.Errorf("failed to cache refresh token: %w", err)
	}

	logger.BizLogger(c).Infof("user %d logged in successfully", u.ID)

	return &vo.LoginResponse{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
	}, nil
}

// Logout 处理用户登出逻辑
func (us *UserServiceImpl) Logout(c *app.RequestContext) (*vo.LogoutResponse, error) {
	logoutLock.Lock()
	defer logoutLock.Unlock()

	userID, exists := c.Get(consts.JWTSubjectClaim)
	if !exists {
		logger.BizLogger(c).Errorf("user ID not found in context")
		return nil, fmt.Errorf("user ID not found in context")
	}

	err := global.RedisClient.Del(context.Background(), fmt.Sprintf("%s:%d", consts.AuthAccessTokenKeyPrefix, userID.(int64))).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to delete access token cache for user %d: %v", userID.(int64), err)
		return nil, fmt.Errorf("failed to delete access token cache: %w", err)
	}

	err = global.RedisClient.Del(context.Background(), fmt.Sprintf("%s:%d", consts.AuthRefreshTokenKeyPrefix, userID.(int64))).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to delete refresh token cache for user %d: %v", userID.(int64), err)
		return nil, fmt.Errorf("failed to delete refresh token cache: %w", err)
	}

	logger.BizLogger(c).Infof("user %d logged out successfully", userID.(int64))

	return &vo.LogoutResponse{
		Message: "Logged out successfully",
	}, nil
}

// RefreshToken 刷新token逻辑
func (us *UserServiceImpl) RefreshToken(c *app.RequestContext, req *dto.RefreshTokenRequest) (*vo.RefreshTokenResponse, error) {
	cfgs, err := configs.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfgs.AppConfig.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		logger.BizLogger(c).Errorf("invalid refresh token provided: %v", err)
		return nil, fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logger.BizLogger(c).Errorf("invalid token claims format")
		return nil, fmt.Errorf("invalid token claims")
	}

	userIDFloat, ok := claims[consts.JWTSubjectClaim].(float64)
	if !ok {
		logger.BizLogger(c).Errorf("invalid user ID in refresh token claims")
		return nil, fmt.Errorf("invalid refresh token")
	}
	userID := int64(userIDFloat)

	now := time.Now()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		consts.JWTSubjectClaim: userID,
		"exp":                  now.Add(time.Duration(cfgs.AppConfig.JWT.ExpireTime) * time.Hour).Unix(),
		"iat":                  now.Unix(),
	})
	accessTokenString, err := accessToken.SignedString([]byte(cfgs.AppConfig.JWT.Secret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		consts.JWTSubjectClaim: userID,
		"exp":                  now.Add(time.Duration(cfgs.AppConfig.JWT.RefreshExpire) * time.Hour).Unix(),
		"iat":                  now.Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(cfgs.AppConfig.JWT.Secret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	refreshCacheKey := fmt.Sprintf("%s:%d", consts.AuthRefreshTokenKeyPrefix, userID)
	cachedRefreshToken := global.RedisClient.Get(context.Background(), refreshCacheKey).Val()
	if cachedRefreshToken == "" {
		logger.BizLogger(c).Errorf("refresh token not found in cache for user %d", userID)
		return nil, fmt.Errorf("refresh token expired or not found")
	}
	if cachedRefreshToken != req.RefreshToken {
		logger.BizLogger(c).Errorf("refresh token mismatch for user %d", userID)
		return nil, fmt.Errorf("invalid refresh token")
	}

	cacheKey := fmt.Sprintf("%s:%d", consts.AuthAccessTokenKeyPrefix, userID)
	err = global.RedisClient.Set(context.Background(), cacheKey, accessTokenString, time.Duration(cfgs.AppConfig.JWT.ExpireTime)*time.Hour).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to update access token cache for user %d: %v", userID, err)
		return nil, fmt.Errorf("failed to update access token cache: %w", err)
	}

	err = global.RedisClient.Set(context.Background(), refreshCacheKey, refreshTokenString, time.Duration(cfgs.AppConfig.JWT.RefreshExpire)*time.Hour).Err()
	if err != nil {
		logger.BizLogger(c).Errorf("failed to update refresh token cache for user %d: %v", userID, err)
		return nil, fmt.Errorf("failed to update refresh token cache: %w", err)
	}

	logger.BizLogger(c).Infof("token refreshed successfully for user %d", userID)

	return &vo.RefreshTokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

// GetProfile 获取用户资料逻辑
func (us *UserServiceImpl) GetProfile(c *app.RequestContext) (*vo.GetProfileResponse, error) {
	userID, exists := c.Get(consts.JWTSubjectClaim)
	if !exists {
		logger.BizLogger(c).Errorf("unable to get current user ID from context")
		return nil, fmt.Errorf("authentication required")
	}

	u, err := us.userMapper.GetUserByID(c, userID.(int64))
	if err != nil {
		logger.BizLogger(c).Errorf("user not found: %v", err)
		return nil, fmt.Errorf("user not found: %w", err)
	}

	roles, err := us.rbacMapper.GetUserRoles(c, strconv.FormatInt(u.ID, 10))
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user roles: %v", err)
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	userRoles := make([]string, 0, len(roles))
	for _, role := range roles {
		userRoles = append(userRoles, role.V1)
	}

	return &vo.GetProfileResponse{
		ID:       strconv.FormatInt(u.ID, 10),
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Roles:    userRoles,
	}, nil
}

// Update 更新用户信息逻辑
func (us *UserServiceImpl) Update(c *app.RequestContext, req *dto.UpdateRequest) (*vo.UpdateResponse, error) {
	userID, exists := c.Get(consts.JWTSubjectClaim)
	if !exists {
		logger.BizLogger(c).Errorf("unable to get user ID from context")
		return nil, fmt.Errorf("unable to get user ID from context")
	}

	u, err := us.userMapper.GetUserByID(c, userID.(int64))
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user info: %v", err)
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	if req.Nickname != "" && u.Nickname != req.Nickname {
		existingUser, err := us.userMapper.GetUserByNickname(c, req.Nickname)
		if err == nil && existingUser.ID != u.ID {
			logger.BizLogger(c).Errorf("nickname '%s' is already in use", req.Nickname)
			return nil, fmt.Errorf("nickname '%s' is already in use", req.Nickname)
		}
	}

	if req.Nickname != "" {
		u.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		u.Avatar = req.Avatar
	}

	if err := us.userMapper.UpdateUser(c, u); err != nil {
		logger.BizLogger(c).Errorf("user update failed for '%s': %v", u.Email, err)
		return nil, fmt.Errorf("user update failed: %w", err)
	}

	roles, err := us.rbacMapper.GetUserRoles(c, strconv.FormatInt(u.ID, 10))
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user roles: %v", err)
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	userRoles := make([]string, 0, len(roles))
	for _, role := range roles {
		userRoles = append(userRoles, role.V1)
	}

	return &vo.UpdateResponse{
		ID:       strconv.FormatInt(u.ID, 10),
		Email:    u.Email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Roles:    userRoles,
	}, nil
}

// ResetPassword 重置密码逻辑
func (us *UserServiceImpl) ResetPassword(c *app.RequestContext, req *dto.ResetPasswordRequest) (*vo.ResetPasswordResponse, error) {
	passwordResetLock.Lock()
	defer passwordResetLock.Unlock()

	userID, exists := c.Get(consts.JWTSubjectClaim)
	if !exists {
		logger.BizLogger(c).Errorf("unable to get user ID from context")
		return nil, fmt.Errorf("unable to get user ID from context")
	}

	u, err := us.userMapper.GetUserByID(c, userID.(int64))
	if err != nil {
		logger.BizLogger(c).Errorf("user not found: %v", err)
		return nil, fmt.Errorf("user not found: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.OldPassword))
	if err != nil {
		logger.BizLogger(c).Errorf("old password verification failed for user '%s': %v", u.Email, err)
		return nil, fmt.Errorf("old password verification failed")
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.BizLogger(c).Errorf("password hashing failed for user '%s': %v", u.Email, err)
		return nil, fmt.Errorf("password hashing failed: %w", err)
	}

	u.Password = string(newPassword)

	if err := us.userMapper.UpdateUser(c, u); err != nil {
		logger.BizLogger(c).Errorf("password update failed for user '%s': %v", u.Email, err)
		return nil, fmt.Errorf("password update failed: %w", err)
	}

	return &vo.ResetPasswordResponse{
		Message: "密码重置成功",
	}, nil
}

// ListUsers 获取用户列表逻辑
func (us *UserServiceImpl) ListUsers(c *app.RequestContext, req *dto.ListUsersRequest) (*vo.ListUsersResponse, error) {
	users, total, err := us.userMapper.ListUsers(c, req.PageNo, req.PageSize, req.Keyword, req.Role)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get user list: %v", err)
		return nil, fmt.Errorf("failed to get user list: %w", err)
	}

	list := make([]*vo.UserItem, 0, len(users))
	for _, u := range users {
		userIDStr := strconv.FormatInt(u.ID, 10)
		roles, err := us.rbacMapper.GetUserRoles(c, userIDStr)
		userRoles := make([]string, 0)
		if err == nil {
			for _, role := range roles {
				userRoles = append(userRoles, role.V1)
			}
		}

		list = append(list, &vo.UserItem{
			ID:       strconv.FormatInt(u.ID, 10),
			Email:    u.Email,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
			Roles:    userRoles,
		})
	}

	return &vo.ListUsersResponse{
		Total:    total,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		List:     list,
	}, nil
}

// UpdateUserRole 管理员更新用户角色
func (us *UserServiceImpl) UpdateUserRole(c *app.RequestContext, req *dto.UpdateUserRoleRequest) (*vo.UpdateUserRoleResponse, error) {
	userID, exists := c.Get(consts.JWTSubjectClaim)
	if !exists {
		logger.BizLogger(c).Errorf("unable to get current user ID from context")
		return nil, fmt.Errorf("authentication required")
	}

	currentUserID, ok := userID.(int64)
	if !ok {
		logger.BizLogger(c).Errorf("invalid user ID type in context")
		return nil, fmt.Errorf("invalid user ID type")
	}

	currentUserIDStr := strconv.FormatInt(currentUserID, 10)
	currentUserRoles, err := us.rbacMapper.GetUserRoles(c, currentUserIDStr)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get current user roles: %v", err)
		return nil, fmt.Errorf("failed to get current user roles: %w", err)
	}

	var hasPermission bool
	for _, role := range currentUserRoles {
		permission, err := global.Enforcer.Enforce(role.V1, string(c.Path()), string(c.Method()))
		if err != nil {
			logger.BizLogger(c).Errorf("failed to check user permissions for role %s: %v", role.V1, err)
			continue
		}
		if permission {
			hasPermission = true
			break
		}
	}

	if !hasPermission {
		logger.BizLogger(c).Warnf("user ID %d attempted to update user role without permission", currentUserID)
		return nil, fmt.Errorf("insufficient permissions: you do not have permission to update user roles")
	}

	if fmt.Sprintf("%d", currentUserID) == req.ID {
		logger.BizLogger(c).Warnf("user ID %d attempted to modify their own role", currentUserID)
		return nil, fmt.Errorf("cannot modify your own role")
	}

	targetUserID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid target user ID format: %v", err)
		return nil, fmt.Errorf("invalid target user ID format: %w", err)
	}

	targetUser, err := us.userMapper.GetUserByID(c, targetUserID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get target user info: %v", err)
		return nil, fmt.Errorf("failed to get target user info: %w", err)
	}

	oldRoles, err := us.rbacMapper.GetUserRoles(c, req.ID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get target user roles: %v", err)
		return nil, fmt.Errorf("failed to get target user roles: %w", err)
	}

	for _, oldRole := range oldRoles {
		if _, err := us.rbacMapper.RevokeRole(c, req.ID, oldRole.V1); err != nil {
			logger.BizLogger(c).Errorf("failed to remove old role %s: %v", oldRole.V1, err)
			return nil, fmt.Errorf("failed to remove old role: %w", err)
		}
	}

	if _, err := us.rbacMapper.AssignRole(c, req.ID, req.Role); err != nil {
		logger.BizLogger(c).Errorf("failed to add new role %s: %v", req.Role, err)
		return nil, fmt.Errorf("failed to add new role: %w", err)
	}

	updatedRoles, err := us.rbacMapper.GetUserRoles(c, req.ID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get updated user roles: %v", err)
		return nil, fmt.Errorf("failed to get updated user roles: %w", err)
	}

	userRoles := make([]string, 0, len(updatedRoles))
	for _, role := range updatedRoles {
		userRoles = append(userRoles, role.V1)
	}

	return &vo.UpdateUserRoleResponse{
		ID:       req.ID,
		Email:    targetUser.Email,
		Nickname: targetUser.Nickname,
		Roles:    userRoles,
		Message:  fmt.Sprintf("User role successfully updated to %s", req.Role),
	}, nil
}
