package Controler

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"fmt"
	"strings"
	"Website/Models"
)

func OpenTemplate(w http.ResponseWriter, r *http.Request, vars interface{}, filename string,headerVar Models.HeaderVariables) {

	var allFiles []string
	files, err := ioutil.ReadDir("./view")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, "./view/"+filename)
		}
	}

	templates, err := template.ParseFiles(allFiles...)

	s1 := templates.Lookup("header.html")
	s1.Execute(w, headerVar)

	s2 := templates.Lookup("navigation.html")
	loggedIn, _ := Authenticated(r)
	s2.Execute(w, Models.NavigationVariables{LoggedIn: loggedIn})

	s3 := templates.Lookup("jumbotron.html")
	s3.Execute(w, nil)

	data := templates.Lookup(filename)
	data.Execute(w, vars)

	//t, err := template.ParseFiles("view/" + filename)
	//if err != nil { // if there is an error
	//	log.Print("template parsing error: ", err) // log it
	//}
	//
	//err = t.Execute(w, vars)
	//if err != nil { // if there is an error
	//	log.Print("template executing error: ", err) //log it
	//}

	w.Write([]byte("</div></div>"))

	footer := templates.Lookup("footer.html")
	footer.Execute(w, headerVar)
}
