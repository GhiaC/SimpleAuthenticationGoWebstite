package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageVariables struct {
	Date         string
	Time         string
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
	t, err := template.ParseFiles("view/homepage.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, nil) //execute the template and pass it the HomePageVars struct to fill in the gaps
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

func HomePage(w http.ResponseWriter, r *http.Request){

	now := time.Now() // find the time right now
	HomePageVars := PageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("view/homepage.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

//func showHtml(page , vars string)  {
//
//	t, err := template.ParseFiles("view/"+page+".html") //parse the html file homepage.html
//	if err != nil { // if there is an error
//		log.Print("template parsing error: ", err) // log it
//	}
//	err = t.Execute(w, vars) //execute the template and pass it the HomePageVars struct to fill in the gaps
//	if err != nil { // if there is an error
//		log.Print("template executing error: ", err) //log it
//	}
//}