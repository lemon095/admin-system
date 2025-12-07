package main

import (
	"admin-system/model"
	"log"
	"os"
	"os/signal"
	"syscall"

	"admin-system/config"
	"admin-system/database"
	"admin-system/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	config.Init()

	// 初始化数据库
	if err := database.InitMySQL(); err != nil {
		log.Fatalf("MySQL初始化失败: %v", err)
	}
	defer database.CloseMySQL()

	// 初始化Redis
	if err := database.InitRedis(); err != nil {
		log.Fatalf("Redis初始化失败: %v", err)
	}
	defer database.CloseRedis()

	// 初始化数据表
	if err := model.InitTables(); err != nil {
		log.Fatalf("数据表初始化失败: %v", err)
	}

	// 设置Gin模式
	if config.ServerMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	port := ":" + config.ServerPort
	log.Printf("服务器启动在端口 %s", port)

	// 优雅关闭
	go func() {
		if err := r.Run(port); err != nil {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("服务器正在关闭...")
}
