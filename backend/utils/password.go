package utils

import (
	"crypto/md5"
	"fmt"
)

// MD5Hash 生成MD5哈希
func MD5Hash(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}

// VerifyPassword 验证密码
func VerifyPassword(password, hashedPassword string) bool {
	return MD5Hash(password) == hashedPassword
}

