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

// GetPostsByUserID 获取用户文章
func GetPostsByUserID(id int) (*[]Post, error) {
	var posts []Post
	// Preload 相当于预先 join table，不然取到的 posts 就没有 User 信息
	if err := db.Preload("User").Where("user_id=?", id).Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

// GetPostsByUserIDPageAndLimit 获取用户文章(带分页)
func GetPostsByUserIDPageAndLimit(id, page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit
	if err := db.Preload("User").Order("timestamp desc").Where("user_id=?", id).Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Where("user_id=?", id).Count(&total)
	return &posts, total, nil
}

// GetPostsByPageAndLimit 所有帖子
func GetPostsByPageAndLimit(page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit
	if err := db.Preload("User").Order("timestamp desc").Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, nil
	}
	db.Model(&Post{}).Count(&total)
	return &posts, total, nil
}
