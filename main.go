package main

import (
	"html/template"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	)

var engine *xorm.Engine

type PageVariables struct {
	PageTitle        string
	Answer           string
}


func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/status", status)
	http.HandleFunc("/admin", admin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	var errDB error
	engine, errDB = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")

	print(errDB)
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	submit := r.PostForm.Get("submit")

	vars := PageVariables{
		Answer : "",
		PageTitle: "Login",
	}

	if submit != "" && (username == "" || password == "" ){
		vars.Answer = "username or password is empty"
	}

	t, err := template.ParseFiles("view/login.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, vars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func register(w http.ResponseWriter, r *http.Request) {

}

func status(w http.ResponseWriter, r *http.Request) {

}

func admin(w http.ResponseWriter, r *http.Request) {

}

func HomePage(w http.ResponseWriter, r *http.Request) {

	//now := time.Now() // find the time right now
	//HomePageVars := PageVariables{ //store the date and time in a struct
	//	Date: now.Format("02-01-2006"),
	//	Time: now.Format("15:04:05"),
	//}
	//
	//t, err := template.ParseFiles("view/login.html") //parse the html file homepage.html
	//if err != nil { // if there is an error
	//	log.Print("template parsing error: ", err) // log it
	//}
	//err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	//if err != nil { // if there is an error
	//	log.Print("template executing error: ", err) //log it
	//}
}
