package main

import (
	"log"
	"test_gorm/lesson01"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Parent struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Child struct {
	Parent
	Age int
}

//func InitDB(dst ...interface{}) *gorm.DB {
//	db, err := gorm.Open(mysql.Open("root:st123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
//	if err != nil {
//		panic(err)
//	}
//
//	db.AutoMigrate(dst...)
//
//	return db
//}

func main() {
	// 连接SQLite数据库
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&Parent{}, &Child{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// 调用lesson01
	lesson01.Run(db)
}
