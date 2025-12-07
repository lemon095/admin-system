package database

import (
	"database/sql"
	"fmt"
	"log"

	"admin-system/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *sql.DB
var GormDB *gorm.DB

// InitMySQL 初始化MySQL连接
func InitMySQL() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBCharset,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("打开数据库连接失败: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("数据库连接失败: %v", err)
	}

	// 设置连接池参数
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)

	log.Println("MySQL连接成功")

	GormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: DB,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return nil
}

// CloseMySQL 关闭MySQL连接
func CloseMySQL() {
	if DB != nil {
		DB.Close()
		log.Println("MySQL连接已关闭")
	}
}
