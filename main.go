package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//database constants
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	//opening conncetion to the database
	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataB, err := sql.Open("postgres", connecDB)
	defer dataB.Close()
	if err != nil {
		panic(err)
	}
	//actual program running here.
	/* EVERYTHING RUNS AND BEEN TESTED HERE
	var tempUSR usr.User = bank.GetUsrInfo("akhv", dataB)
	fmt.Println("unaltered usr\n", tempUSR)
	bank.Deposit(tempUSR, 200, dataB)
	fmt.Println("User after deposit\n", tempUSR)
	tempUSR = bank.Withdraw(tempUSR, 200, dataB)
	fmt.Println("all changes to user", tempUSR)

	bank.UpdateUserDB(dataB, tempUSR)
	var tempUSR2 usr.User = bank.GetUsrInfo("akhv", dataB)
	fmt.Println("pulled account from db\n", tempUSR2)
	*/
}
