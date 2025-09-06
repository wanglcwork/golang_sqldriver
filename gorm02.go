// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

package main

import (
	"fmt"
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
	dsn := "user:password@tcp()(sqlite:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 查询某个用户发布的所有文章及其对应的评论信息
	var user User
	userID := 1 // 假设查询用户ID为1的用户
	err = db.Preload("Posts.Comments").First(&user, userID).Error
	if err != nil {
		fmt.Println("Error fetching user posts and comments:", err)
	}
	fmt.Printf("User: %s\n", user.Username)
	for _, post := range user.Posts {
		fmt.Printf("Post: %s\n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf(" - Comment: %s\n", comment.Content)
		}
	}
	// 查询评论数量最多的文章信息
	var postWithMostComments Post
	err = db.Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count desc").
		Limit(1).
		Scan(&postWithMostComments).Error
	if err != nil {
		fmt.Println("Error fetching post with most comments:", err)
	}
	fmt.Printf("Post with most comments: %s\n", postWithMostComments.Title)
}

// 注意：请根据实际情况修改数据库连接字符串 (dsn)。
// 运行此代码后，Gorm 会执行关联查询，获取指定用户的所有文章及其评论，以及评论数量最多的文章信息。
