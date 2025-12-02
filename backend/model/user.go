package model

import (
	"database/sql"
	"time"
)

// UserRole 用户身份枚举
type UserRole int

const (
	RoleSuperAdmin UserRole = 1 // 超级管理员
	RoleAdmin      UserRole = 2 // 管理员
	RoleUser       UserRole = 3 // 普通用户
)

// User 用户模型
type User struct {
	ID          int64     `json:"id" db:"id"`
	Username    string    `json:"username" db:"username"`
	AvatarID    string    `json:"avatar_id" db:"avatar_id"`
	Role        UserRole  `json:"role" db:"role"`
	Password    string    `json:"-" db:"password"` // 不返回给前端
	Status      int       `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	CreatedBy   *int64    `json:"created_by" db:"created_by"`
	LastLoginAt sql.NullTime `json:"last_login_at" db:"last_login_at"`
	LastLoginIP sql.NullString `json:"last_login_ip" db:"last_login_ip"`
}

// UserResponse 用户响应结构（不包含敏感信息）
type UserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	AvatarID  string    `json:"avatar_id"`
	Role      UserRole  `json:"role"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy *int64    `json:"created_by"`
}

// ToResponse 转换为响应结构
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		AvatarID:  u.AvatarID,
		Role:      u.Role,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy,
	}
}

// IsSuperAdmin 判断是否为超级管理员
func (u *User) IsSuperAdmin() bool {
	return u.Role == RoleSuperAdmin
}

