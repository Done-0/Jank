package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	accessSecret      = []byte("jank-blog-secret")         // Access Token 使用的密钥
	refreshSecret     = []byte("jank-blog-refresh-secret") // Refresh Token 使用的密钥
	accessExpireTime  = time.Hour * 2                      // Access Token 有效期
	refreshExpireTime = time.Hour * 48                     // Refresh Token 有效期
	clockSkew         = 5 * time.Second                    // 允许的时间偏差量
)

// GenerateJWT 生成 Access Token 和 Refresh Token
func GenerateJWT(userId uint) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().UTC().Add(accessExpireTime).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(accessSecret)
	if err != nil {
		return "", "", err
	}

	// 生成 Refresh Token，有效期 48 小时
	refreshClaims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().UTC().Add(refreshExpireTime).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// ValidateJWTToken 验证 Access Token 或 Refresh Token
func ValidateJWTToken(tokenString string, isRefreshToken bool) (*jwt.Token, error) {
	secret := accessSecret
	if isRefreshToken {
		secret = refreshSecret
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return secret, nil
	})

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
