package model

import "time"

// Post model
type Post struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	User      User
	Body      string     `gorm:"type:varchar(522)"`
	Timestamp *time.Time `sql:"DEFAULT:current_timestamp"`
}

// GetPostByUserID 获取用户文章
func GetPostByUserID(id int) (*[]Post, error) {
	var posts []Post
	// Preload 相当于预先 join table，不然取到的 posts 就没有 User 信息
	if err := db.Preload("User").Where("user_id=?", id).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}
