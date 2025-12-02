package database

import (
	"fmt"
	"time"
)

// SaveTokenToRedis 保存token到Redis
func SaveTokenToRedis(userID int64, token string, expireTime time.Time) error {
	key := fmt.Sprintf("token:%d", userID)
	duration := time.Until(expireTime)
	if duration <= 0 {
		return fmt.Errorf("token已过期")
	}
	return RDB.Set(ctx, key, token, duration).Err()
}

// GetTokenFromRedis 从Redis获取token
func GetTokenFromRedis(userID int64) (string, error) {
	key := fmt.Sprintf("token:%d", userID)
	return RDB.Get(ctx, key).Result()
}

// DeleteTokenFromRedis 从Redis删除token
func DeleteTokenFromRedis(userID int64) error {
	key := fmt.Sprintf("token:%d", userID)
	return RDB.Del(ctx, key).Err()
}

// RefreshTokenInRedis 刷新Redis中的token
func RefreshTokenInRedis(userID int64, newToken string, expireTime time.Time) error {
	key := fmt.Sprintf("token:%d", userID)
	duration := time.Until(expireTime)
	if duration <= 0 {
		return fmt.Errorf("token已过期")
	}
	return RDB.Set(ctx, key, newToken, duration).Err()
}

// CheckTokenExists 检查token是否存在
func CheckTokenExists(userID int64) (bool, error) {
	key := fmt.Sprintf("token:%d", userID)
	count, err := RDB.Exists(ctx, key).Result()
	return count > 0, err
}
