package router

import (
	"admin-system/controller"
	"admin-system/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 跨域中间件
	r.Use(middleware.CORSMiddleware())

	// 创建控制器
	authController := controller.NewAuthController()

	// 公开路由
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authController.Login)
		}
	}

	// 需要认证的路由
	apiAuth := api.Group("")
	apiAuth.Use(middleware.AuthMiddleware())
	{
		auth := apiAuth.Group("/auth")
		{
			auth.POST("/refresh", authController.RefreshToken)
			auth.GET("/userinfo", authController.GetUserInfo)
		}
	}

	return r
}

