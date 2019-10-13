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

/*
tmm test functions
error handling make sure program runs even if there is an error
flags try to include them
create user id port the file handling to be able to save unapproved accounts and users
to implement finding whic id to use just  use this sql query
also create accounts for approval
Example
	SELECT COUNT(ProductID)
	FROM Products;

check out how log works and you can proabbly just return out of errors and continue the runnnin of your program.
research
create function to transfer funds between accounts probably be able to reuse the withdraw and deposit.
added things to examples
if there is time add http functionality
employee pkg not used as of now will be used at UI
checkout requirements and eexamples for comments and maybe more ideas
*/

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
