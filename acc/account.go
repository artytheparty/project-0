package acc

import "fmt"

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
func (a *Account) Deposit(val float64) {
	a.accBal = a.accBal + val
}

//Withdraw Withdraws money from account
func (a *Account) Withdraw(val float64) {
	if a.accBal >= val {
		a.accBal = a.accBal - val
	} else {
		fmt.Println("NOT ENOUGH MONEY!!!")
	}
}
