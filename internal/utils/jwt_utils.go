package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtSecret     = []byte("jank-blog-secret")         // Access Token 使用的密钥
	refreshSecret = []byte("jank-blog-refresh-secret") // Refresh Token 使用的密钥
)

// GenerateJWT 生成 Access Token 和 Refresh Token
func GenerateJWT(userId uint, email string) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// 生成 Refresh Token，有效期 48 小时
	refreshClaims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 48).Unix(),
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
	secret := jwtSecret
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
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("token has expired")
			}
		} else {
			return nil, fmt.Errorf("missing exp claim")
		}
		return token, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
