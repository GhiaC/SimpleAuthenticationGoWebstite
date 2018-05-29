package Controler

import (
	"net/http"

	"Website/Models"
)

func Admin(w http.ResponseWriter, r *http.Request) {
}

func Status(w http.ResponseWriter, r *http.Request) {
	if ok, _ := Authenticated(r); ok {
		var users []Models.User
		GetEngine().Table("users").Cols("id", "username").Find(&users)
		//if err == nil {
		result := Models.StatusPageVariables{Users: users}
		OpenTemplate(w, r, result, "status.html",Models.HeaderVariables{Title:"Users"})
		//}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
