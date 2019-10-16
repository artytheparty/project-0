package main

import (
	"database/sql"
	"fmt"

	"github.com/artytheparty/project-0/ui"
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
	fmt.Println("            WELCOME TO THE                                  ")
	fmt.Println(" ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄     ▄▄▄▄▄▄▄▄▄▄▄     ▄▄▄▄▄▄▄▄▄▄    ")
	fmt.Println("▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌   ▐░░░░░░░░░░░▌   ▐░░░░░░░░░░▌   ")
	fmt.Println("▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌   ▐░█▀▀▀▀▀▀▀▀▀    ▐░█▀▀▀▀▀▀▀█░▌  ")
	fmt.Println("▐░▌          ▐░▌       ▐░▌   ▐░▌             ▐░▌       ▐░▌  ")
	fmt.Println("▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄█░▌   ▐░█▄▄▄▄▄▄▄▄▄    ▐░▌       ▐░▌  ")
	fmt.Println("▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌   ▐░░░░░░░░░░░▌   ▐░▌       ▐░▌  ")
	fmt.Println("▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀█░█▀▀    ▐░█▀▀▀▀▀▀▀▀▀    ▐░▌       ▐░▌  ")
	fmt.Println("▐░▌          ▐░▌     ▐░▌     ▐░▌             ▐░▌       ▐░▌  ")
	fmt.Println("▐░▌ ▄        ▐░▌      ▐░▌  ▄ ▐░█▄▄▄▄▄▄▄▄▄  ▄ ▐░█▄▄▄▄▄▄▄█░▌▄ ")
	fmt.Println("▐░▌▐░▌       ▐░▌       ▐░▌▐░▌▐░░░░░░░░░░░▌▐░▌▐░░░░░░░░░░▌▐░▌")
	fmt.Println(" ▀  ▀         ▀         ▀  ▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀  ▀▀▀▀▀▀▀▀▀▀  ▀ ")
	fmt.Println("                          BANKING SYSTEM                    ")
	fmt.Println(" FRIENDSHIP    RELIABILITY     EXPERIENCE        DRIVE      ")

	ui.Menu(dataB)

}
