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

	//results, err := engine.Query("select * from users")
	//fmt.Println(results,err)


	http.HandleFunc("/", HomePage)
	http.HandleFunc("/login", Controler.Login)
	http.HandleFunc("/register", Controler.Register)
	http.HandleFunc("/status", Controler.Status)
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

	Controler.OpenTemplate(w,HomePageVars,"homepage.html")
}
