package main

import (
	"net/http"

	"github.com/dota2mm/go-mega/controller"
)

func main() {
	controller.Startup()

	http.ListenAndServe(":5000", nil)
}
