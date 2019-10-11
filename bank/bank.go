package bank

import (
	"database/sql"
	"fmt"

	"github.com/artytheparty/project-0/acc"
	"github.com/artytheparty/project-0/usr"

	//"go/types: don't report errors referring to invalid types or operands"

	_ "github.com/lib/pq"
)

//Deposit deposits amount into selected account if many
func Deposit(a usr.User, dAmt float64, db *sql.DB) {
	var accHolder []acc.Account = a.GetAccounts()
	//get user information holders
	if len(a.GetAccounts()) > 1 {
		for k := range accHolder {
			accHolder[k].AccountInfo()
		}
		fmt.Println("Please input the Account Number to select an account to deposit to: ")
		choice, _ := fmt.Scanln()
		accHolder[choice].Deposit(dAmt)
	} else {
		accHolder[0].Deposit(dAmt)
	}

}

//GetUsrInfo bjhi
func GetUsrInfo(username string, db *sql.DB) usr.User {
	//user variabales
	var accs []acc.Account
	row := db.QueryRow("SELECT * FROM users WHERE username= $1", username)
	var id string
	var un string
	var pass string
	var fN string
	var lN string
	row.Scan(&id, &un, &pass, &fN, &lN)
	rows, _ := db.Query("SELECT * FROM accounts WHERE usrid=$1", id)
	for rows.Next() {
		var aN string
		var aT string
		var aB float64
		rows.Scan(&aN, &id, &aT, &aB)
		accs = append(accs, acc.CreateAccount(aN, id, aT, aB))
	}
	return usr.CreateUser(id, un, pass, fN, lN, accs)
}
