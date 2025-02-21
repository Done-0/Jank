package utils

import (
	"fmt"
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
func GenerateJWT(accountID uint) (string, string, error) {
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
func ValidateJWTToken(tokenString string, isRefreshToken bool) (*jwt.Token, error) {
	// 移除 Bearer 前缀
	tokenString = removeBearerPrefix(tokenString)

	secret := accessSecret
	if isRefreshToken {
		secret = refreshSecret
	}

	token, err := validateToken(tokenString, secret)
	if err != nil {
		return nil, err
	}

	// 验证 token 的 claims 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, fmt.Errorf("无效 token")
	} else {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().UTC().Add(clockSkew).Unix() > int64(exp) {
				if isRefreshToken {
					return nil, fmt.Errorf("refresh token 已过期，请重新登录")
				}
				return nil, fmt.Errorf("access token 已过期，请重新登录")
			}
		} else {
			return nil, fmt.Errorf("缺少 exp 字段")
		}
	}

	return token, nil
}

// ParseAccountIDFromJWT 解析 JWT 并从中提取 accountID
func ParseAccountIDFromJWT(tokenString string) (uint, error) {
	// 移除 Bearer 前缀
	tokenString = removeBearerPrefix(tokenString)

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

	return uint(accountID), nil
}

// removeBearerPrefix 去除 Bearer 前缀
func removeBearerPrefix(tokenString string) string {
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		return tokenString[7:]
	}
	return tokenString
}

// generateToken 通用的 token 生成函数
func generateToken(accountID uint, secret []byte, expireTime time.Duration) (string, error) {
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
func validateToken(tokenString string, secret []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
