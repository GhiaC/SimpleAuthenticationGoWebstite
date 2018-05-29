package main

import (
	"html/template"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//"log/syslog"
	"fmt"
	"log/syslog"
	"github.com/gorilla/sessions"
	"time"
)

var engine *xorm.Engine

type PageVariables struct {
	PageTitle string
	Answer    string
}

type users struct {
	Id       int64
	Username string `xorm:"varchar(256) not null"`
	Password string `xorm:"varchar(256) not null"`
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key-test-1234567890")
	store = sessions.NewCookieStore(key)
)

func authentiated(r *http.Request) (bool, string) {
	session, _ := store.Get(r, "cookie-name")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		//http.Error(w, "Forbidden", http.StatusForbidden)
		return false, ""
	}
	return true, session.Values["username"].(string)
}

var u = &users{}

func main() {
	var errDB error
	engine, errDB = xorm.NewEngine("mysql", "root:mghiasi@/authGo?charset=utf8")
	fmt.Println(errDB)
	engine.CreateTables(u)
	engine.Sync2(new(users))

	//var id []users
	//err := engine.Table("users").Where("username = ? and password = ? ", user1.Username , user1.Password).Limit(10,0).Find(&id)
	//fmt.Println(id)

	//results, err := engine.Query("select * from users")
	//fmt.Println(results,err)

	logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	if err != nil {
		log.Fatalf("Fail to create xorm system logger: %v\n", err)
	}

	logger := xorm.NewSimpleLogger(logWriter)
	logger.ShowSQL(true)
	engine.SetLogger(logger)

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/status", status)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/admin", admin)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	submit := r.PostForm.Get("submit")

	vars := PageVariables{
		Answer:    "",
		PageTitle: "Login",
	}

	if submit != "" && (username == "" || password == "") {
		vars.Answer = "username or password is empty"
	} else {
		var id int
		has, err := engine.Table("users").Where("username = ? and password = ? ", username, password).Cols("id").Get(&id)
		println(has, id)
		if has && err == nil && id > 0 {
			session, _ := store.Get(r, "cookie-name")
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Values["id"] = id
			//id,setted := session.Values["id"].(int)
			session.Save(r, w)
		} else {
			vars.Answer = "username or password is wrong"
		}
	}

	if ok, _ := authentiated(r); !ok {
		t, err := template.ParseFiles("view/login.html") //parse the html file homepage.html
		if err != nil { // if there is an error
			log.Print("template parsing error: ", err) // log it
		}

		err = t.Execute(w, vars) //execute the template and pass it the HomePageVars struct to fill in the gaps
		if err != nil { // if there is an error
			log.Print("template executing error: ", err) //log it
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Values["username"] = "empty"
	session.Values["id"] = 0
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	submit := r.PostForm.Get("submit")

	vars := PageVariables{
		Answer:    "",
		PageTitle: "Register",
	}

	if submit != "" && (username == "" || password == "") {
		vars.Answer = "username or password is empty"
	} else if submit != "" {
		user1 := new(users)
		user1.Username = username
		user1.Password = password

		affected, err := engine.Insert(user1)

		if affected > 0 && err == nil {
			vars.Answer = "you are Registered. Go to Login Page"
		}

	}

	if ok, _ := authentiated(r); ok {
		t, err := template.ParseFiles("view/register.html") //parse the html file homepage.html
		if err != nil { // if there is an error
			log.Print("template parsing error: ", err) // log it
		}

		err = t.Execute(w, vars) //execute the template and pass it the HomePageVars struct to fill in the gaps
		if err != nil { // if there is an error
			log.Print("template executing error: ", err) //log it
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func status(w http.ResponseWriter, r *http.Request) {

}

func admin(w http.ResponseWriter, r *http.Request) {

}

type HomePageVariables struct {
	Date        string
	Time        string
	LoginStatus string
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now() // find the time right now
	HomePageVars := HomePageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
		LoginStatus: "you aren't logged in",
	}
	if ok, username := authentiated(r); ok {
		HomePageVars.LoginStatus = "dear " + username + ", you are logged in"
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
