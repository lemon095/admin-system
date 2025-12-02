package controller

import (
	"net/http"

	"admin-system/service"
	"admin-system/utils"

	"github.com/gin-gonic/gin"
)

// AuthController 认证控制器
type AuthController struct {
	authService *service.AuthService
}

// NewAuthController 创建认证控制器
func NewAuthController() *AuthController {
	return &AuthController{
		authService: &service.AuthService{},
	}
}

// Login 登录接口
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body service.LoginRequest true "登录信息"
// @Success 200 {object} service.LoginResponse
// @Failure 400 {object} map[string]interface{}
// @Router /api/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req service.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 获取客户端IP
	clientIP := ctx.ClientIP()

	// 调用服务层
	resp, err := c.authService.Login(&req, clientIP)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data":    resp,
	})
}

// RefreshToken 刷新token接口
// @Summary 刷新token
// @Description 刷新JWT token
// @Tags 认证
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/auth/refresh [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	// 从上下文获取用户信息（由中间件设置）
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		ctx.Abort()
		return
	}

	username, _ := ctx.Get("username")
	role, _ := ctx.Get("role")

	// 生成新token
	newToken, err := utils.GenerateToken(
		userID.(int64),
		username.(string),
		role.(int),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "生成token失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "刷新成功",
		"data": gin.H{
			"token": newToken,
		},
	})
}

// GetUserInfo 获取当前用户信息
// @Summary 获取用户信息
// @Description 获取当前登录用户信息
// @Tags 认证
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/auth/userinfo [get]
func (c *AuthController) GetUserInfo(ctx *gin.Context) {
	// 从上下文获取用户信息（由中间件设置）
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未授权",
		})
		ctx.Abort()
		return
	}

	username, _ := ctx.Get("username")
	role, _ := ctx.Get("role")

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"user_id":  userID,
			"username": username,
			"role":     role,
		},
	})
}
