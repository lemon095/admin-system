package database

import (
	"context"
	"fmt"
	"log"

	"admin-system/config"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

// InitRedis 初始化Redis连接
func InitRedis() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	// 测试连接
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redis连接失败: %v", err)
	}

	log.Println("Redis连接成功")
	return nil
}

// CloseRedis 关闭Redis连接
func CloseRedis() {
	if RDB != nil {
		RDB.Close()
		log.Println("Redis连接已关闭")
	}
}

