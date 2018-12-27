package controller

import (
	"log"
	"net/http"
)

// middleAuth func - session 校验用户身份
func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		log.Println("auth username: ", username)
		if err != nil {
			log.Println("middle auth get session err and redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
