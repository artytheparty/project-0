package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	_ "github.com/artytheparty/project-0/acc"
	_ "github.com/artytheparty/project-0/dbhandling"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	// fmt.Println("Welcome to the Banking App")
	// admin := acc.CreateAccount("kjonah", "password", "Kyle", "Johna", 1, 5000)
	// admin2 := acc.CreateAccount("jjohna", "password", "John", "Johna", 2, 15000)
	// admin3 := acc.CreateAccount("ukra", "password", "Ursula", "Kung", 3, 5000)
	// admin4 := acc.CreateAccount("shword", "password", "Shelly", "Wordinson", 4, 5000)
	// admin5 := acc.CreateAccount("uthough", "password", "Umelia", "Thoughtful", 5, 5000)
	// admin6 := acc.CreateAccount("kps13", "password", "Kong", "Pierce", 6, 5000)
	// admin7 := acc.CreateAccount("bhide", "password", "Robert", "Hide", 7, 5000)
	// admin8 := acc.CreateAccount("dlecious", "password", "David", "Lecious", 8, 5000)
	// admin9 := acc.CreateAccount("hparker", "password", "Hong", "Parker", 9, 5000)
	// admin10 := acc.CreateAccount("golang", "password", "Godin", "Lang", 10, 5000)

	//fmt.Println(admin.Info())
	// accounts := make(map[int]acc.AccountHolder)
	// accounts[1] = admin
	// accounts[2] = admin2
	// accounts[3] = admin3
	// accounts[4] = admin4
	// accounts[5] = admin5
	// accounts[6] = admin6
	// accounts[7] = admin7
	// accounts[8] = admin8
	// accounts[9] = admin9
	// accounts[10] = admin10
	//initializing accounts list
	//dbhandling.WriteToFile(accounts)
	// accounts2 := dbhandling.ReadFile()
	// for k := range accounts2 {
	// 	accounts2[k].Info()
	// }
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	getAll(db)
}
func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CUSTOMERS")
	for rows.Next() {
		var username string
		var pass string
		var fName string
		var lName string
		var aNum string
		var accBal float64
		rows.Scan(&username, &pass, &fName, &lName, &aNum, &accBal)
		fmt.Println(username, pass, fName, lName, aNum, accBal)
	}
}
