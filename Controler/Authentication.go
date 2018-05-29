package Controler

import (
	"net/http"
	"Website/Models"
)

func Authenticated(r *http.Request) (bool, string) {
	session, _ := Store.Get(r, "cookie-name")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		return false, ""
	}
	return true, session.Values["username"].(string)
}


func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	submit := r.PostForm.Get("submit")

	vars := Models.PageVariables{
		Answer:    "",
		PageTitle: "Login",
	}

	if submit == "Login" && (username == "" || password == "") {
		vars.Answer = "username or password is empty"
	} else if username != "" && password != "" {
		var id int
		engine := GetEngine()
		has, err := engine.Table("users").Where("username = ? and password = ? ", username, password).Cols("id").Get(&id)
		println(has, id)
		if has && err == nil && id > 0 {
			session, _ := Store.Get(r, "cookie-name")
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Values["id"] = id
			//id,setted := session.Values["id"].(int)
			session.Save(r, w)
		} else {
			vars.Answer = "username or password is wrong"
		}
	}

	if ok, _ := Authenticated(r); !ok {
		OpenTemplate(w,vars,"login.html")
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Values["username"] = "empty"
	session.Values["id"] = 0
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
