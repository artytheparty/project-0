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

//CreateAccount fnskfnsod
func CreateAccount(aN string, aID string, typ string, initialDeposit float64) Account {
	return Account{aN, aID, typ, initialDeposit}
}

//AccountInfo prints out the account info
func (a *Account) AccountInfo() {
	fmt.Printf("Acc. Number: %s User Associated ID: %s Acc. Type: %s Balance: %f\n",
		a.accNum, a.usrID, a.accType, a.accBal)
}

//Deposit works i hope
func (a *Account) Deposit(val float64) Account {
	var updatedBal float64
	updatedBal = a.accBal + val
	return Account{a.accNum, a.usrID, a.accType, updatedBal}
}

//Withdraw Withdraws money from account
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

//UpdateAccountSlice updates old account with new account in the specified position
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
