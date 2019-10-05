package acc

import (
	"fmt"
	"strconv"
)

//AccountHolder Structure will be used to create a new account
type AccountHolder struct {
	accountNum                  int
	fName, lname, username, psw string
	accountBal                  float64
}

//Info will print out the infor of an account
func (a AccountHolder) Info() string {
	return ("Account Number: " + strconv.Itoa(a.accountNum) +
		"\nLast Name: " + a.lname +
		"\nFirst Name: " + a.fName +
		"username :" + a.username +
		"\nAccount balance: " + strconv.FormatFloat(a.accountBal, 'f', 2, 64))
}

//CreateAccount will create an account and return the
//created account
func CreateAccount(counter int, fname string, lname string, un string, pass string, initialdeposit float64) AccountHolder {
	return AccountHolder{counter, fname, lname, un, pass, initialdeposit}
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
