package main

import (
	"net/http"

	"github.com/dota2mm/go-mega/model"
	"github.com/gorilla/context"

	"github.com/dota2mm/go-mega/controller"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Setup db
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup controller
	controller.Startup()

	http.ListenAndServe(":5000", context.ClearHandler(http.DefaultServeMux))
}
