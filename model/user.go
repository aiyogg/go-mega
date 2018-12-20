package model

// User model
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `ogrem:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorem:"type:varchar(128)"`
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

// GetUserByUsername 根据用户名查用户
func (u *User) GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
