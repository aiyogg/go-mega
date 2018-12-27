package controller

import (
	"net/http"

	"github.com/dota2mm/go-mega/vm"
)

type home struct{}

func (h home) registerRouters() {
	http.HandleFunc("/", indexHander)
	http.HandleFunc("/login", loginHander)
	http.HandleFunc("/logout", logoutHandler)
}

func indexHander(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}

func loginHander(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetVM()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if len(username) < 3 {
			v.AddError("username must longer than 3 characters")
		}
		if len(password) < 6 {
			v.AddError("username must longer than 6 characters")
		}

		if !vm.CheckLogin(username, password) {
			v.AddError("username password not correct, please input agin")
		}

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			// fmt.Fprintf(w, "Username: %s Password: %s", username, password)
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
