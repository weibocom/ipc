package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func auth(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// TODO: removed this trick later
		fn(w, r, params)

		// user := getUserName(r)
		// if user == "" {
		// 	http.Redirect(w, r, "/", http.StatusFound)
		// 	return
		// }

		// w.Header().Set("Content-Type", "application/json")
		// fn(w, r, params)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.FormValue("name")
	pass := r.FormValue("password")
	redirectTarget := "/"
	if checkUser(name, pass) {
		setSession(name, w)
		redirectTarget = "/#/accounts"
	}
	http.Redirect(w, r, redirectTarget, http.StatusFound)
}

func checkUser(name, pass string) bool {
	if name == "" || pass == "" {
		return false
	}

	// TODO: 从数据库中校验用户
	return name == "admin" && pass == "admin"
}

func logoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}
