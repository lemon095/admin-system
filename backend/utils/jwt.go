package utils

import (
	"errors"
	"time"

	"admin-system/config"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.JWTSecret)

// Claims JWT声明
type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID int64, username string, role int) (string, error) {
	expireTime := time.Now().Add(time.Duration(config.JWTExpireHours) * time.Hour)
	claims := &Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}

// ShouldRefreshToken 判断是否需要刷新token（距离过期时间小于阈值）
func ShouldRefreshToken(claims *Claims) bool {
	if claims.ExpiresAt == nil {
		return true
	}

	expireTime := claims.ExpiresAt.Time
	threshold := time.Duration(config.JWTRefreshThresholdHours) * time.Hour
	now := time.Now()
	return expireTime.Sub(now) < threshold
}

