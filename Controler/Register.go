package Controler

import (
	"net/http"
	"Website/Models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	submit := r.PostForm.Get("submit")

	vars := Models.PageVariables{
		Answer:    "",
		PageTitle: "Register",
	}

	if submit != "" && (username == "" || password == "") {
		vars.Answer = "username or password is empty"
	} else if submit != "" {

		newUser := Models.NewUser(username,password)
		affected, err := GetEngine().Insert(newUser)

		if affected > 0 && err == nil {
			vars.Answer = "you are Registered. Go to Login Page"
		}
	}

	if ok, _ := Authenticated(r); ok {
		OpenTemplate(w,vars,"register.html")
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}