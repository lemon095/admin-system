package model

import (
	"admin-system/database"
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// InitTables 初始化数据表
func InitTables() error {
	// 读取SQL文件并执行
	sqlContent := `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
		username VARCHAR(50) NOT NULL COMMENT '用户名',
		avatar_id VARCHAR(255) DEFAULT NULL COMMENT '头像ID/URL',
		role TINYINT NOT NULL DEFAULT 3 COMMENT '用户身份：1-超级管理员，2-管理员，3-普通用户',
		password VARCHAR(32) NOT NULL COMMENT '密码（MD5值）',
		status TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1-启用，0-禁用',
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
		created_by BIGINT UNSIGNED DEFAULT NULL COMMENT '创建者ID',
		last_login_at DATETIME DEFAULT NULL COMMENT '最后登录时间',
		last_login_ip VARCHAR(50) DEFAULT NULL COMMENT '最后登录IP',
		PRIMARY KEY (id),
		UNIQUE KEY uk_username (username),
		KEY idx_role (role),
		KEY idx_status (status),
		KEY idx_created_by (created_by)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';
	`

	_, err := database.DB.Exec(sqlContent)
	if err != nil {
		return fmt.Errorf("创建数据表失败: %v", err)
	}

	log.Println("数据表初始化成功")

	// 初始化默认超级管理员账号
	return initDefaultAdmin()
}

// initDefaultAdmin 初始化默认超级管理员账号
func initDefaultAdmin() error {
	// 检查是否已存在用户
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return fmt.Errorf("查询用户数量失败: %v", err)
	}

	// 如果已有用户，则不创建默认账号
	if count > 0 {
		log.Println("已存在用户，跳过默认账号创建")
		return nil
	}

	// 创建默认超级管理员账号
	username := "chuchangkeji"
	password := "chuchangkeji666"
	passwordMD5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	avatarID := "default-avatar-001"
	role := int(RoleSuperAdmin)
	status := 1
	createdBy := int64(1) // 自己创建自己

	query := `
		INSERT INTO users (username, avatar_id, role, password, status, created_by, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err = database.DB.Exec(query, username, avatarID, role, passwordMD5, status, createdBy, time.Now())
	if err != nil {
		return fmt.Errorf("创建默认管理员账号失败: %v", err)
	}

	// 更新created_by为自己
	_, err = database.DB.Exec("UPDATE users SET created_by = id WHERE username = ?", username)
	if err != nil {
		log.Printf("更新created_by失败: %v", err)
	}

	log.Printf("默认超级管理员账号创建成功: %s", username)
	return nil
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, username, avatar_id, role, password, status, created_at, updated_at, 
		       created_by, last_login_at, last_login_ip
		FROM users
		WHERE username = ? AND status = 1
	`

	err := database.DB.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.AvatarID,
		&user.Role,
		&user.Password,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.CreatedBy,
		&user.LastLoginAt,
		&user.LastLoginIP,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("用户不存在")
	}
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}

	return user, nil
}

// UpdateUserLoginInfo 更新用户登录信息
func UpdateUserLoginInfo(userID int64, ip string) error {
	query := `
		UPDATE users 
		SET last_login_at = ?, last_login_ip = ?
		WHERE id = ?
	`
	_, err := database.DB.Exec(query, time.Now(), ip, userID)
	return err
}
