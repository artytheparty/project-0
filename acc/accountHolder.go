package acc

import (
	"fmt"
	"strconv"
)

//AccountHolder Structure will be used to create a new account
type AccountHolder struct {
	username, psw, fName, lName string
	accountNum                  int
	accountBal                  float64
}

//Info will print out the infor of an account
func (a AccountHolder) Info() string {
	return ("\nUsername: " + a.username +
		"\nPassword: " + a.psw +
		"\nLast Name: " + a.lName +
		"\nFirst Name :" + a.fName +
		"\nAcccount Number" + strconv.Itoa(a.accountNum) +
		"\nAccount balance: " + strconv.FormatFloat(a.accountBal, 'f', 2, 64))
}

//CreateAccount will create an account and return the
//created account
func CreateAccount(un string, pass string, fname string, lname string, aCnum int, initialdeposit float64) AccountHolder {
	return AccountHolder{un, pass, fname, lname, aCnum, initialdeposit}
}

//Withdraw will take out the specified value
func (a *AccountHolder) Withdraw(w float64) {
	if a.accountBal > w {
		a.accountBal = a.accountBal - w
	} else {
		fmt.Println("Cannot withdraw more than your available balance!")
	}

}

//Deposit method
func (a *AccountHolder) Deposit(w float64) {
	a.accountBal += w
}
