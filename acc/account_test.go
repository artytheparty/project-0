package acc

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	hold.Deposit(20)
	if hold.accBal == 263.50 {
		fmt.Println("Success")
	} else {
		panic("BRUUUH")
	}
}
func TestWithdraw(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	hold.Withdraw(250)
	if hold.accBal == 43.50 {
		fmt.Println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
