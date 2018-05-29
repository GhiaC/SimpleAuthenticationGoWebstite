package Controler

import (
	"net/http"
	"Website/Models"
)

func Register(w http.ResponseWriter, r *http.Request) {

	if ok, _ := Authenticated(r); ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		r.ParseForm()
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		submit := r.PostForm.Get("submit")

		vars := Models.LoginPageVariables{
			Answer:      "",
			Url:         "/register",
			SubmitValue: "Register",
		}

		if submit != "" && (username == "" || password == "") {
			vars.Answer = "username or password is empty"
		} else if username != "" && password != "" {

			newUser := Models.NewUser(username, password)
			affected, err := GetEngine().Insert(newUser)

			if affected > 0 && err == nil {
				vars.Answer = "you are Registered. Go to Login Page"
			}
		}
		OpenTemplate(w, r, vars, "login.html",Models.HeaderVariables{Title:"Register"})
	}
}
