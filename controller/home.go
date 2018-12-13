package controller

import (
	"net/http"

	"github.com/dota2mm/go-mega/vm"
)

type home struct{}

func (h home) registerRouters() {
	http.HandleFunc("/", indexHander)
	http.HandleFunc("/login", loginHander)
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

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			// fmt.Fprintf(w, "Username: %s Password: %s", username, password)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
