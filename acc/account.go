package acc

import (
	"fmt"
)

//Account Structure
type Account struct {
	accNum  string
	usrID   string
	accType string
	accBal  float64
}

//CreateAccount Simple Constructor which will return a new Account
func CreateAccount(aN string, aID string, typ string, initialDeposit float64) Account {
	return Account{aN, aID, typ, initialDeposit}
}

//AccountInfo Prints Account information straight to console | use case in printing the informatio on the screen for the user to view
func (a *Account) AccountInfo() {
	fmt.Printf("Acc. Number: %s User Associated ID: %s Acc. Type: %s Balance: %f\n",
		a.accNum, a.usrID, a.accType, a.accBal)
}

//Deposit deposits passed amount of money into the account the method has benn called on.
func (a *Account) Deposit(val float64) Account {
	var updatedBal float64
	updatedBal = a.accBal + val
	return Account{a.accNum, a.usrID, a.accType, updatedBal}
}

//Withdraw Withdraws the passed amount from the account the function is called on.
func (a *Account) Withdraw(val float64) Account {
	var updatedBal float64
	if a.accBal >= val {
		updatedBal = a.accBal - val
	} else {
		fmt.Println("NOT ENOUGH MONEY!!!")
		updatedBal = a.accBal
	}
	return Account{a.accNum, a.usrID, a.accType, updatedBal}
}

//UpdateAccountSlice Updates the Account slice with the latest information assosiated with a user
func UpdateAccountSlice(ac []Account, a Account, pos int) []Account {
	holder := ac
	holder[pos] = a
	return holder
}

//GetAccountNum return the account number
func (a *Account) GetAccountNum() string {
	return a.accNum
}

//GetUsrID returns the accounts userid
func (a *Account) GetUsrID() string {
	return a.usrID
}

//GetAccountType returns the account tye
func (a *Account) GetAccountType() string {
	return a.accType
}

//GetAccountBal returns the account balance
func (a *Account) GetAccountBal() float64 {
	return a.accBal
}
