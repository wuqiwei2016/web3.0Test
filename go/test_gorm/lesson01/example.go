package lesson01

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
	ignored      string         // fields that aren't exported are ignored
}

type Member struct {
	gorm.Model
	Name string
	Age  uint8
}

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	Author
	ID      int
	Upvotes int32
}

type Blog2 struct {
	ID     int64
	Author Author `gorm:"embedded;embeddedPrefix:author_"`
	// Author  Author
	Upvotes int32
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	// db.AutoMigrate(&Member{})
	// db.AutoMigrate(&Blog{})
	db.AutoMigrate(&Blog2{})

	user := &User{}                // 创建一个 User 结构体指针实例
	user.MemberNumber.Valid = true // 设置 MemberNumber 字段为有效状态
	db.Create(user)                // 将 user 对象插入数据库

	// create传指针
	// mem := Member{}
	// db.Create(&mem)
	// fmt.Println(mem.ID)
	// db.Delete(&Member{}, 1)
}
