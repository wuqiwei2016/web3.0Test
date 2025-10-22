package main

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 连接到SQLite数据库
func ConnectDatabase() {
	var err error
	// 连接到SQLite数据库
	DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// 自动迁移表结构
	DB.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 创建一个初始管理员用户（如果不存在）
	var admin User
	if err := DB.Where("username = ?", "admin").First(&admin).Error; err != nil {
		// 管理员不存在，创建一个
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		admin := User{
			Username: "admin",
			Password: string(passwordHash),
			Email:    "wuqiwei2002@sina.com",
		}
		DB.Create(&admin)
	}
}
