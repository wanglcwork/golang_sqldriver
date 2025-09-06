// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

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
	PostCount int    `gorm:"default:0"`         // 文章数量统计字段
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post 模型
type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"size:200;not null"`
	Content       string    `gorm:"type:text;not null"`
	UserID        uint      `gorm:"not null"`              // 外键，关联 User
	Comments      []Comment `gorm:"foreignKey:PostID"`     // 一对多关系
	CommentStatus string    `gorm:"size:50;default:'有评论'"` // 评论状态字段
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Comment 模型
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	PostID    uint   `gorm:"not null"` // 外键，关联 Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 在 Post 创建前的钩子函数，更新用户的文章数量统计字段
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	var user User
	if err := tx.First(&user, p.UserID).Error; err != nil {
		return err
	}
	user.PostCount++
	if err := tx.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// 在 Comment 删除后的钩子函数，检查文章的评论数量并更新评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	if err := tx.First(&post, c.PostID).Error; err != nil {
	}
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err
	}
	if commentCount == 0 {
		post.CommentStatus = "无评论"
		if err := tx.Save(&post).Error; err != nil {
			return err
		}
	}
	return nil
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
		First(&postWithMostComments).Error
	if err != nil {
		fmt.Println("Error fetching post with most comments:", err)
	}
	fmt.Printf("Post with most comments: %s\n", postWithMostComments.Title)
}
