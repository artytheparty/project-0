package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/artytheparty/project-0/acc"
	"github.com/artytheparty/project-0/bank"
	_ "github.com/lib/pq"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/login", login)
	http.ListenAndServe(":6060", nil)
}

//Userholder Struct holds all the information which should be passed to the template
type Userholder struct {
	ID, Username, Password, Fname, Lname string
	Accounts                             []acc.Account
	IncorrectUnOrPass                    bool
	Loggedin                             bool
	Notsignedin                          bool
}

var globalUser Userholder = Userholder{}

func depwit(response http.ResponseWriter, request *http.Request) {
	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataB, err := sql.Open("postgres", connecDB)
	defer dataB.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to DB")
	temp, _ := template.ParseFiles("template/login.html")

	//executes the template last thing
	fmt.Println(temp.Execute(response, v))
}

func login(response http.ResponseWriter, request *http.Request) {
	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataB, err := sql.Open("postgres", connecDB)
	defer dataB.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to DB")
	v := Userholder{}
	//navigation variables
	v.IncorrectUnOrPass = false
	v.Loggedin = false
	v.Notsignedin = true
	temp, _ := template.ParseFiles("template/login.html")
	uhold := request.FormValue("username")
	phold := request.FormValue("pw")
	fmt.Println("userinput:", uhold, "userpass:", phold)
	usrhold := bank.GetUsrInfo(uhold, dataB)
	if phold == usrhold.GetUsrPassword() {
		v.Loggedin = true
		v.IncorrectUnOrPass = false
		v.Notsignedin = false
		v.ID = usrhold.GetUsrID()
		v.Username = uhold
		v.Password = phold
		v.Fname = usrhold.GetUsrFName()
		v.Lname = usrhold.GetUsrLName()
		v.Accounts = usrhold.GetAccounts()
		globalUser = v
	} else {
		v.Loggedin = false
		v.IncorrectUnOrPass = true
		v.Notsignedin = false
	}

	//executes the template last thing
	fmt.Println(temp.Execute(response, v))
}

// func connect() *sql.DB {
// 	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	dataB, err := sql.Open("postgres", connecDB)
// 	defer dataB.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("connected to DB")
// 	return dataB
// }

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)
