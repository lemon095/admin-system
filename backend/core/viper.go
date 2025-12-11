package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Viper 配置
func Viper() *viper.Viper {
	if err := godotenv.Load(); err != nil {
		log.Println("未找到.env文件，使用环境变量")
	}

	v := viper.New()
	v.SetConfigFile("env")
	v.AutomaticEnv()

	global.GVA_CONFIG.Mysql.Path = getEnv("DB_HOST", "rm-2ze72h0maxcrf682uno.mysql.rds.aliyuncs.com")
	global.GVA_CONFIG.Mysql.Port = getEnv("DB_PORT", "33301")
	global.GVA_CONFIG.Mysql.Username = getEnv("DB_USER", "game_user")
	global.GVA_CONFIG.Mysql.Password = getEnv("DB_PASSWORD", "*Chuchangkeji")
	global.GVA_CONFIG.Mysql.Dbname = getEnv("DB_NAME", "bunker")
	global.GVA_CONFIG.Mysql.Config = getEnv("DB_CONFIG", "charset=utf8mb4&parseTime=True&loc=Local")
	//DBCharset = getEnv("DB_CHARSET", "utf8mb4")

	global.GVA_CONFIG.Redis.Addr = fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "172.28.196.80"), getEnv("REDIS_PORT", "6379"))
	global.GVA_CONFIG.Redis.Password = getEnv("REDIS_PASSWORD", "")
	global.GVA_CONFIG.Redis.DB, _ = strconv.Atoi(getEnv("REDIS_DB", "0"))

	global.GVA_CONFIG.JWT.SigningKey = getEnv("JWT_SECRET", "your-secret-key-change-in-production")
	global.GVA_CONFIG.JWT.ExpiresTime = getEnv("JWT_EXPIRE_HOURS", "72") + "h"
	global.GVA_CONFIG.JWT.BufferTime = getEnv("JWT_REFRESH_THRESHOLD_HOURS", "24") + "h"

	global.GVA_CONFIG.Zap.Level = getEnv("ZAP_LEVEl", "info")
	global.GVA_CONFIG.Zap.Format = getEnv("ZAP_FORMAT", "console")
	global.GVA_CONFIG.Zap.Director = getEnv("ZAP_DIRECTOR", "log")

	global.GVA_CONFIG.System.Addr = getEnv("SERVER_PORT", "9001")
	//ServerMode = getEnv("SERVER_MODE", "release")

	global.GVA_CONFIG.Local.StorePath = getEnv("LOCAL_STORE_PATH", "uploads/file")

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")

	global.GVA_CONFIG.Captcha.KeyLong, _ = strconv.Atoi(getEnv("CAPTCHA_KEY_LONG", "6"))
	global.GVA_CONFIG.Captcha.ImgWidth, _ = strconv.Atoi(getEnv("CAPTCHA_IMG_WIDTH", "240"))
	global.GVA_CONFIG.Captcha.ImgHeight, _ = strconv.Atoi(getEnv("CAPTCHA_IMG_HEIGHT", "80"))
	global.GVA_CONFIG.Captcha.OpenCaptcha, _ = strconv.Atoi(getEnv("CAPTCHA_OPEN_CAPTCHA", "0"))
	global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut, _ = strconv.Atoi(getEnv("CAPTCHA_OPEN_CAPTCHA_TIMEOUT", "3600"))

	return v
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
