// Package utils 提供JWT令牌生成与验证工具
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	// 密钥和有效期配置
	accessSecret      = []byte("jank-blog-secret")         // Access Token 使用的密钥
	refreshSecret     = []byte("jank-blog-refresh-secret") // Refresh Token 使用的密钥
	accessExpireTime  = time.Hour * 2                      // Access Token 有效期
	refreshExpireTime = time.Hour * 48                     // Refresh Token 有效期
	clockSkew         = 5 * time.Second                    // 允许的时间偏差量
)

// GenerateJWT 生成 Access Token 和 Refresh Token
// 参数：
//   - accountID: 账户ID
//
// 返回值：
//   - string: Access Token
//   - string: Refresh Token
//   - error: 生成过程中的错误
func GenerateJWT(accountID int64) (string, string, error) {
	accessTokenString, err := generateToken(accountID, accessSecret, accessExpireTime)
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := generateToken(accountID, refreshSecret, refreshExpireTime)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// ValidateJWTToken 验证 Access Token 或 Refresh Token
// 参数：
//   - tokenString: 令牌字符串
//   - isRefreshToken: 是否为刷新令牌
//
// 返回值：
//   - *jwt.Token: 验证通过的令牌
//   - error: 验证过程中的错误
func ValidateJWTToken(tokenString string, isRefreshToken bool) (*jwt.Token, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	secret := accessSecret
	if isRefreshToken {
		secret = refreshSecret
	}

	token, err := validateToken(tokenString, secret)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, fmt.Errorf("无效 token")
	} else {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().UTC().Add(clockSkew).Unix() > int64(exp) {
				if isRefreshToken {
					return nil, fmt.Errorf("refresh token 已过期，请重新登录")
				}
				return nil, fmt.Errorf("access token 已过期，请使用 refresh token 获取新的 access token")
			}
		} else {
			return nil, fmt.Errorf("缺少 exp 字段")
		}
	}

	return token, nil
}

// RefreshTokenLogic 负责刷新 Token
// 参数：
//   - refreshTokenString: 刷新令牌字符串
//
// 返回值：
//   - map[string]string: 包含新的 Access Token 和 Refresh Token 的映射
//   - error: 刷新过程中的错误
func RefreshTokenLogic(refreshTokenString string) (map[string]string, error) {
	token, err := ValidateJWTToken(refreshTokenString, true)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accountID := int64(claims["account_id"].(float64))

		newAccessToken, newRefreshToken, err := GenerateJWT(accountID)
		if err != nil {
			return nil, err
		}

		return map[string]string{
			"accessToken":  newAccessToken,
			"refreshToken": newRefreshToken,
		}, nil
	}

	return nil, fmt.Errorf("refresh token 验证失败")
}

// ParseAccountAndRoleIDFromJWT 从 JWT 中提取 accountID 和 roleID
// 参数：
//   - tokenString: 令牌字符串
//
// 返回值：
//   - int64: 账户ID
//   - error: 解析过程中的错误
func ParseAccountAndRoleIDFromJWT(tokenString string) (int64, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := ValidateJWTToken(tokenString, false)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("无法解析 access token 中的 claims")
	}

	accountID, ok := claims["account_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("access token 中缺少 account_id")
	}

	return int64(accountID), nil
}

// generateToken 通用的 token 生成函数
// 参数：
//   - accountID: 账户ID
//   - secret: 密钥
//   - expireTime: 过期时间
//
// 返回值：
//   - string: 生成的令牌
//   - error: 生成过程中的错误
func generateToken(accountID int64, secret []byte, expireTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"account_id": accountID,
		"exp":        time.Now().UTC().Add(expireTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// validateToken 验证 token 是否有效
// 参数：
//   - tokenString: 令牌字符串
//   - secret: 密钥
//
// 返回值：
//   - *jwt.Token: 验证通过的令牌
//   - error: 验证过程中的错误
func validateToken(tokenString string, secret []byte) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}
