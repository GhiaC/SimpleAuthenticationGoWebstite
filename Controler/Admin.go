package Controler

import (
	"net/http"

	"Website/Models"
)

func Admin(w http.ResponseWriter, r *http.Request) {
}

func Status(w http.ResponseWriter, r *http.Request) {
	var users []Models.User
	err := GetEngine().Table("users").Cols("id","username").Find(&users)
	if err == nil {
		result := Models.StatusPageVariables{Users:users}
		OpenTemplate(w,result,"status.html")
	}
}
