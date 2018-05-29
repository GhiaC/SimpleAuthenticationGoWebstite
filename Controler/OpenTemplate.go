package Controler

import (
	"net/http"
	"log"
	"html/template"
)

func OpenTemplate(w http.ResponseWriter , vars interface{}, filename string)  {
	t, err := template.ParseFiles("view/"+filename)
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, vars)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

