package ui

import (
	"database/sql"
	"fmt"

	"github.com/artytheparty/project-0/bank"
	"github.com/artytheparty/project-0/emp"
	"github.com/artytheparty/project-0/usr"
)

//Menu Displays the main menu
func Menu(db *sql.DB) {
	var choice string
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
	fmt.Println("1) Customer LogIn")
	fmt.Println("2) Employee LogIn")
	fmt.Println("3) Apply for an account")
	fmt.Print("Make a choice: ")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		UserSignIn(db)
	case "2":
		EmployeeSignIn(db)
	case "3":

	}
}

//UserSignIn will ask the usr to sign in
func UserSignIn(db *sql.DB) {

	var uHolder string
	var pHolder string
	fmt.Println("Enter your username: ")
	fmt.Scan(&uHolder)
	fmt.Println("Enter your pasword: ")
	fmt.Scan(&pHolder)
	userHolder := bank.GetUsrInfo(uHolder, db)
	if userHolder.GetUsrPassword() == pHolder {
		fmt.Println("Success, Welcome!")
		UserMenu(bank.GetUsrInfo(uHolder, db), db)
	} else {
		fmt.Println("Wrong username or password!")
		Menu(db)
	}
	return
}

//UserMenu will display the menu forthe user
func UserMenu(a usr.User, db *sql.DB) {
	var choice string
	fmt.Println("1) Deposit money into account")
	fmt.Println("2) Withdraw money from account")
	//fmt.Println("3) Transfer money to another account(must know the acount number)")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		var amount float64
		fmt.Println("How much would youlike to deposit?: ")
		fmt.Scan(&amount)
		bank.Deposit(a, amount, db)
		fmt.Print("Deposit Successful!")
	case "2":
		var amount float64
		fmt.Println("How much would youlike to Withdraw: ")
		fmt.Scan(&amount)
		bank.Withdraw(a, amount, db)
		fmt.Print("Withdraw Successful!")
	}
}

//EmployeeSignIn wll show the option fo employees
func EmployeeSignIn(db *sql.DB) {

	var uHolder string
	var pHolder string
	fmt.Println("Enter your username: ")
	fmt.Scan(&uHolder)
	fmt.Println("Enter your pasword: ")
	fmt.Scan(&pHolder)
	employeeHolder := bank.GetEmployeeInfo(uHolder, db)
	if employeeHolder.GetEmployeePass() == pHolder {
		fmt.Println("Success, Welcome!")
		EmployeeMenu(bank.GetEmployeeInfo(uHolder, db), db)
	} else {
		fmt.Println("Wrong username or password!")
		Menu(db)
	}
	return
}

//EmployeeMenu Will print out the menu for employees
func EmployeeMenu(a emp.Employee, db *sql.DB) {
	fmt.Println("1) View Pending Accounts")
}
