package middleware

import (
	"net/http"
	"strings"

	"admin-system/database"
	"admin-system/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头获取token
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "未提供认证token",
			})
			ctx.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token格式错误",
			})
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		// 解析token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "无效的token: " + err.Error(),
			})
			ctx.Abort()
			return
		}

		// 检查Redis中的token是否存在（可选，用于token撤销）
		exists, err := database.CheckTokenExists(claims.UserID)
		if err == nil && !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "token已失效",
			})
			ctx.Abort()
			return
		}

		// 检查是否需要刷新token
		if utils.ShouldRefreshToken(claims) {
			// 生成新token
			newToken, err := utils.GenerateToken(claims.UserID, claims.Username, claims.Role)
			if err == nil {
				// 将新token添加到响应头
				ctx.Header("X-New-Token", newToken)
			}
		}

		// 将用户信息存储到上下文
		ctx.Set("user_id", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}

