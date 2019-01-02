package model

import (
	"fmt"
	"time"
)

// User model
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `ogrem:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorem:"type:varchar(128)"`
	LastSeen     *time.Time
	AboutMe      string `gorem:"type:varchar(140)"`
	Avatar       string `gorem:"type:varchar(200)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:followers;association_jointable_foreignkey:follower_id"`
}

// SetPassword 明文密码加密
func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	return u.PasswordHash == GeneratePasswordHash(password)
}

// SetAvatar 设置默认头像
func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}

// GetUserByUsername 根据用户名查用户
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// AddUser 添加一条用户记录
func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	return db.Create(&user).Error
}

// UpdateUserByUsername 更新用户信息
func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	item, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(item).Update(contents).Error
}

// UpdateLastSeen 更新最后在线时间
func UpdateLastSeen(username string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, contents)
}

// UpdateAboutMe 更新自我介绍
func UpdateAboutMe(username, text string) error {
	contents := map[string]interface{}{"about_me": text}
	return UpdateUserByUsername(username, contents)
}