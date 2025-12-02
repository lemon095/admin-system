package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBCharset           string
	RedisHost           string
	RedisPort           string
	RedisPassword       string
	RedisDB             int
	JWTSecret           string
	JWTExpireHours      int
	JWTRefreshThresholdHours int
	ServerPort          string
	ServerMode          string
)

func Init() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用环境变量")
	}

	DBHost = getEnv("DB_HOST", "mysql")
	DBPort = getEnv("DB_PORT", "3306")
	DBUser = getEnv("DB_USER", "root")
	DBPassword = getEnv("DB_PASSWORD", "root123456")
	DBName = getEnv("DB_NAME", "admin_system")
	DBCharset = getEnv("DB_CHARSET", "utf8mb4")

	RedisHost = getEnv("REDIS_HOST", "redis")
	RedisPort = getEnv("REDIS_PORT", "6379")
	RedisPassword = getEnv("REDIS_PASSWORD", "")
	RedisDB, _ = strconv.Atoi(getEnv("REDIS_DB", "0"))

	JWTSecret = getEnv("JWT_SECRET", "your-secret-key-change-in-production")
	JWTExpireHours, _ = strconv.Atoi(getEnv("JWT_EXPIRE_HOURS", "72"))
	JWTRefreshThresholdHours, _ = strconv.Atoi(getEnv("JWT_REFRESH_THRESHOLD_HOURS", "24"))

	ServerPort = getEnv("SERVER_PORT", "7701")
	ServerMode = getEnv("SERVER_MODE", "release")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

