package service

import (
	"fmt"
	"time"

	"admin-system/config"
	"admin-system/database"
	"admin-system/model"
	"admin-system/utils"
)

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token    string              `json:"token"`
	User     *model.UserResponse `json:"user"`
	ExpireAt int64               `json:"expire_at"`
}

// AuthService 认证服务
type AuthService struct{}

// Login 用户登录
func (s *AuthService) Login(req *LoginRequest, clientIP string) (*LoginResponse, error) {
	// 查询用户
	user, err := model.GetUserByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("用户名或密码错误")
	}

	// 验证密码
	if !utils.VerifyPassword(req.Password, user.Password) {
		return nil, fmt.Errorf("用户名或密码错误")
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Username, int(user.Role))
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %v", err)
	}

	// 更新登录信息
	if err := model.UpdateUserLoginInfo(user.ID, clientIP); err != nil {
		// 登录信息更新失败不影响登录流程
		fmt.Printf("更新登录信息失败: %v\n", err)
	}

	// 将token存储到Redis（用于token刷新和验证）
	expireTime := time.Now().Add(time.Duration(config.JWTExpireHours) * time.Hour)
	if err := database.SaveTokenToRedis(user.ID, token, expireTime); err != nil {
		fmt.Printf("保存token到Redis失败: %v\n", err)
	}

	return &LoginResponse{
		Token:    token,
		User:     user.ToResponse(),
		ExpireAt: expireTime.Unix(),
	}, nil
}
