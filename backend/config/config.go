package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-admin-team/go-admin-core/sdk/config"
	"github.com/joho/godotenv"
)

func init() {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用环境变量")
	}
	config.ApplicationConfig.Port, _ = strconv.ParseInt(getEnv("SERVER_PORT", "9001"), 10, 64)
	config.ApplicationConfig.Mode = getEnv("SERVER_MODE", "release")

	config.LoggerConfig.Path = getEnv("logger.path", "temp/logs")
	config.LoggerConfig.Level = getEnv("logger.level", "trace")
	config.LoggerConfig.EnabledDB = getEnv("logger.enableddb", "false") == "true"

	config.JwtConfig.Secret = getEnv("JWT_SECRET", "your-secret-key-change-in-production")
	hour, _ := strconv.ParseInt(getEnv("JWT_EXPIRE_HOURS", "72"), 10, 64)
	config.JwtConfig.Timeout = hour * 3600

	config.DatabaseConfig.Driver = "mysql"
	config.DatabaseConfig.Source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		getEnv("DB_USER", "game_user"),
		getEnv("DB_PASSWORD", "*Chuchangkeji"),
		getEnv("DB_HOST", "rm-2ze72h0maxcrf682uno.mysql.rds.aliyuncs.com"),
		getEnv("DB_PORT", "33301"),
		getEnv("DB_NAME", "bunker"))
	config.DatabasesConfig["*"] = config.DatabaseConfig

	config.CacheConfig.Memory = struct {
		Addr     string
		Password string
		Db       int
	}{
		Addr:     "127.0.0.1:6379",
		Password: getEnv("REDIS_PASSWORD", ""),
		Db:       0,
	}

	config.GenConfig.DBName = getEnv("DB_NAME", "bunker")
	config.GenConfig.FrontPath = getEnv("frontpath", "../frontend/src")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
