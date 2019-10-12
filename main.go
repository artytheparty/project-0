package main

import (
	"database/sql"
	"fmt"

	"github.com/artytheparty/project-0/bank"
	"github.com/artytheparty/project-0/usr"
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
	var tempUSR usr.User = bank.GetUsrInfo("akhv", dataB)
	fmt.Println(tempUSR)
	//bank.Deposit(tempUSR, 200, dataB)
	//fmt.Println(tempUSR)
	tempUSR = bank.Withdraw(tempUSR, 200, dataB)
	fmt.Println(tempUSR)
	//fmt.Print(len(tempUSR.GetAccounts()))

}
