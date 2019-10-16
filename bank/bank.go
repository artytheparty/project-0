package bank

import (
	"database/sql"
	"fmt"

	"github.com/artytheparty/project-0/acc"
	"github.com/artytheparty/project-0/emp"
	"github.com/artytheparty/project-0/usr"

	//"go/types: don't report errors referring to invalid types or operands"

	_ "github.com/lib/pq"
)

//RecoverError recovers the program from the actual error
func RecoverError() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from Error: ", r)
	}
}

//both deposit and withdraw only edit the first account'

//Deposit deposits amount into selected account if many
func Deposit(a usr.User, dAmt float64, db *sql.DB) usr.User {

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
		modifiedAcc := accHolder[choice].Deposit(dAmt)
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

//PrintUserInfo will pring out te user information
func PrintUserInfo(db *sql.DB, username string) {
	user2 := GetUsrInfo(username, db)
	fmt.Println("Username: ", user2.GetUsrUsername())
	fmt.Println("Account Number: ", user2.GetUsrID())
	fmt.Println("First Name: ", user2.GetUsrFName())
	fmt.Println("Last Name: ", user2.GetUsrLName())
	user2.PrintAccounts()
}

//Withdraw will withdraw money from a user's account
func Withdraw(a usr.User, dAmt float64, db *sql.DB) usr.User {
	defer RecoverError()
	var accHolder []acc.Account = a.GetAccounts()
	//get user information holders
	if len(a.GetAccounts()) > 1 {
		fmt.Println("Which account would you like to withdraw from?")
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
			panic(err)
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

//GetEmployeeInfo returns the Employee struct from DB
func GetEmployeeInfo(usrname string, db *sql.DB) emp.Employee {
	row := db.QueryRow("SELECT * FROM employees WHERE username=$1", usrname)
	var id string
	var username string
	var pass string
	var fName string
	var lName string
	row.Scan(&id, &username, &pass, &fName, &lName)
	return emp.CreateNewEmployee(id, username, pass, fName, lName)

}

//UpdateUserDB pushes the updated user and user's account information to the database
func UpdateUserDB(db *sql.DB, a usr.User) {
	accountsHolder := a.GetAccounts()
	_, err := db.Exec("UPDATE users SET usrid=$1, username=$2, passwd=$3, fname=$4, lname=$5 WHERE usrid=$6",
		a.GetUsrID(), a.GetUsrUsername(), a.GetUsrPassword(), a.GetUsrFName(), a.GetUsrLName(), a.GetUsrID())
	if err != nil {
		fmt.Println("Not running")
	}
	for i := range accountsHolder {
		_, errAcc := db.Exec("UPDATE accounts SET accnum=$1, usrid=$2, acctype=$3, accbal=$4 WHERE accnum=$5",
			accountsHolder[i].GetAccountNum(), accountsHolder[i].GetUsrID(), accountsHolder[i].GetAccountType(), accountsHolder[i].GetAccountBal(), accountsHolder[i].GetAccountNum())
		if errAcc != nil {
			fmt.Println("Account DB update failed")
		}
	}

}

//CreateNewUserEntry creates a new useentry in the database
func CreateNewUserEntry(username string, password string, fname string, lname string, db *sql.DB) {
	var count int
	rows, _ := db.Query("SELECT COUNT(*) as count FROM users")
	for rows.Next() {
		rows.Scan(&count)
	}
	db.Exec("INSERT INTO users VALUES($1, $2, $3, $4, $5)",
		count, username, password, fname, lname)

}

//CreateNewAccountEntry creates a new account in the database
func CreateNewAccountEntry(usrid string, acctype string, bal float64, db *sql.DB) {
	var count int
	rows, _ := db.Query("SELECT COUNT(*) as count FROM accounts")
	for rows.Next() {
		rows.Scan(&count)
	}
	db.Exec("INSERT INTO users VALUES($1, $2, $3, $4)",
		count, usrid, acctype, bal)

}
