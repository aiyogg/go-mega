package model

import "time"

// Post model
type Post struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	User      User
	Body      string     `gorm:"type: varchar(522)"`
	Timestamp *time.Time `sql:"DEFAULT: current_timestamp"`
}
