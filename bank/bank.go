package bank

import (
	"database/sql"
	"fmt"

	"github.com/artytheparty/project-0/acc"
	"github.com/artytheparty/project-0/usr"

	//"go/types: don't report errors referring to invalid types or operands"

	_ "github.com/lib/pq"
)

//both deposit and withdraw only edit the first account'

//Deposit deposits amount into selected account if many
 func Deposit(a usr.User, dAmt float64, db *sql.DB) {
 	var accHolder []acc.Account = a.GetAccounts()
 	//get user information holders
 	if len(a.GetAccounts()) > 1 {
 		var i int = 0
 		for k := range accHolder {
 			fmt.Printf("Option: %d\n", i)
 			accHolder[k].AccountInfo()
 			i++
 		}
 		fmt.Println("Enter the correct option coresponding to your accounts: ")
 		choice, _ := fmt.Scanln()
 		//here the account in the array doesnt get updated
 		//probably need a mehtod also that means withdraw doesnt work either
 		accHolder[choice] = accHolder[choice].Deposit(dAmt)
 	} else {
 		accHolder[0].Deposit(dAmt)
 	}

// }

//Withdraw will withdraw money from a user's account
func Withdraw(a usr.User, dAmt float64, db *sql.DB) usr.User {
	var accHolder []acc.Account = a.GetAccounts()
	//get user information holders
	if len(a.GetAccounts()) > 1 {
		var i int = 0
		for k := range accHolder {
			fmt.Printf("Option: %d\n", i)
			accHolder[k].AccountInfo()
			i++
		}
		fmt.Println("Enter the correct option coresponding to your accounts: ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Incorrect value")
		}
		modifiedAcc := accHolder[choice].Withdraw(dAmt)
		accHolder2 := acc.UpdateAccountSlice(accHolder, modifiedAcc, choice)
		updtUser := usr.UpdateUserAccounts(a, accHolder2)
		return updtUser

	} else {
		modifiedAcc := accHolder[0].Withdraw(dAmt)
		accHolder2 := acc.UpdateAccountSlice(accHolder, modifiedAcc, 0)
		updtUser := usr.UpdateUserAccounts(a, accHolder2)
		return updtUser
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
