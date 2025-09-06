// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 模型
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:100;not null;unique"`
	Email     string `gorm:"size:100;not null;unique"`
	Password  string `gorm:"size:255;not null"`
	Posts     []Post `gorm:"foreignKey:UserID"` // 一对多关系
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post 模型
type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:200;not null"`
	Content   string    `gorm:"type:text;not null"`
	UserID    uint      `gorm:"not null"`          // 外键，关联 User
	Comments  []Comment `gorm:"foreignKey:PostID"` // 一对多关系
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Comment 模型
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   `gorm:"not null"` // 外键，关联 Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	// 连接数据库
	dsn := "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移，创建表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

// 注意：请根据实际情况修改数据库连接字符串 (dsn)。
// 运行此代码后，Gorm 会根据定义的模型自动创建对应的数据库表，并建立相应的外键关系。
