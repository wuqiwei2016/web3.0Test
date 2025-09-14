package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 定义Student模型
type Student struct {
	gorm.Model
	Name  string `gorm:"size:100"`
	Age   int
	Grade string `gorm:"size:50"`
}

func main() {
	// 连接SQLite数据库
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败:", err)
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&Student{})
	if err != nil {
		fmt.Println("迁移表结构失败:", err)
	}

	// 1. 插入新记录
	zhangsan := Student{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.Create(&zhangsan)
	if result.Error != nil {
		fmt.Println("插入失败:", result.Error)
	}
	fmt.Printf("插入成功，ID: %d\n", zhangsan.ID)

	// 2. 查询年龄大于18的学生
	var adultStudents []Student
	result = db.Where("age > ?", 18).Find(&adultStudents)
	if result.Error != nil {
		fmt.Println("查询失败:", result.Error)
	}
	fmt.Println("年龄大于18的学生:")
	for _, student := range adultStudents {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n",
			student.ID, student.Name, student.Age, student.Grade)
	}

	// 3. 更新张三的年级
	result = db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	if result.Error != nil {
		fmt.Println("更新失败:", result.Error)
	}
	fmt.Printf("更新了 %d 条记录\n", result.RowsAffected)

	// 4. 删除年龄小于15的学生
	result = db.Where("age < ?", 15).Delete(&Student{})
	if result.Error != nil {
		fmt.Println("删除失败:", result.Error)
	}
	fmt.Printf("删除了 %d 条记录\n", result.RowsAffected)
}
