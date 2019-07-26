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

	db.DropTableIfExists(model.User{}, model.Post{}, "follower")
	db.CreateTable(model.User{}, model.Post{})

	model.AddUser("Chuck", "123@abc", "i@chenteng.me")
	model.AddUser("Jay", "123@jay", "i@jay.com")

	u1, _ := model.GetUserByUsername("Chuck")
	u1.CreatePost("Jay is my idol.")
	model.UpdateAboutMe(u1.Username, `Coll Web developer`)

	u2, _:= model.GetUserByUsername("Jay")
	u2.CreatePost("Hello, I am JayChou!")
	u2.CreatePost("Chuck is a cool Programer.")

	u1.Follow(u2.Username)

	//users := []model.User{
	//	{
	//		Username:     "Chuck",
	//		PasswordHash: model.GeneratePasswordHash("123@abc"),
	//		Email:        "i@example.com",
	//		Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=retro", model.Md5("i@example.com")),
	//		Posts: []model.Post{
	//			{Body: "This my first post.Thank you!"},
	//		},
	//	},
	//	{
	//		Username:     "Jay",
	//		PasswordHash: model.GeneratePasswordHash("666@jay"),
	//		Email:        "666@jay.com",
	//		Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=retro", model.Md5("666@jay.com")),
	//		Posts: []model.Post{
	//			{Body: "Hello,I'm Jay Chou!"},
	//			{Body: "I think Chuck is a good boy!"},
	//		},
	//	},
	//}
	//
	//for _, u := range users {
	//	db.Debug().Create(&u)
	//}
}
