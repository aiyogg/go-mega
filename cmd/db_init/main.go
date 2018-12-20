package main

import (
	"log"

	"github.com/dota2mm/go-mega/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})

	users := []model.User{
		{
			Username:     "Chuck",
			PasswordHash: model.GeneratePasswordHash("123@abc"),
			Email:        "a@b.c",
			Posts: []model.Post{
				{Body: "This my first post.Thank you!"},
			},
		},
		{
			Username:     "Jay",
			PasswordHash: model.GeneratePasswordHash("666@jay"),
			Email:        "666@jay.com",
			Posts: []model.Post{
				{Body: "Hello,I'm Jay Chou!"},
				{Body: "I think Chuck is a good boy!"},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}
}
