package main

import (
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"Website/Controler"
	"Website/Models"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("Resource"))
	http.Handle("/Resource/", http.StripPrefix("/Resource/", fs))
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/login", Controler.Login)
	http.HandleFunc("/register", Controler.Register)
	http.HandleFunc("/users", Controler.Status)
	http.HandleFunc("/logout", Controler.Logout)
	http.HandleFunc("/admin", Controler.Admin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now() // find the time right now
	HomePageVars := Models.HomePageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
		LoginStatus: "you aren't logged in",
	}
	if ok, username := Controler.Authenticated(r); ok {
		HomePageVars.LoginStatus = "dear " + username + ", you are logged in"
	}

	Controler.OpenTemplate(w,r,HomePageVars,"homepage.html",Models.HeaderVariables{Title:"Authentication GO"})
}
