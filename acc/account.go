package acc

import (
	"fmt"
)

//Account Structure
type Account struct {
	accNum  string
	accID   string
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
		a.accNum, a.accID, a.accType, a.accBal)
}

//Deposit works i hope
func (a *Account) Deposit(val float64) Account {
	var updatedBal float64
	updatedBal = a.accBal + val
	return Account{a.accNum, a.accID, a.accType, updatedBal}
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
	return Account{a.accNum, a.accID, a.accType, updatedBal}
}

//UpdateAccountSlice updates old account with new account in the specified position
func UpdateAccountSlice(ac []Account, a Account, pos int) []Account {
	holder := ac
	holder[pos] = a
	return holder
}
