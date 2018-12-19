package model

// User model
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `ogrem:"type: varchar(64)"`
	Email        string `gorm:"type: varchar(120)"`
	PasswordHash string `gorem:"type: varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many: followers; association_jointable_foreignkey: follower_id"`
}
