package acc

import (
	"fmt"
	"testing"
)

func TestDeposit(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	hold2 := hold.Deposit(20)
	if hold2.accBal == 263.50 {
		fmt.Println("Success")
	} else {
		panic("BRUUUH")
	}
}
func TestWithdraw(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	hold2 := hold.Withdraw(200)
	if hold2.accBal == 43.50 {
		fmt.Println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
func TestCreateAccount(t *testing.T) {
	hold := CreateAccount("2", "1", "s", 243.50)
	if hold.accNum == "2" && hold.usrID == "1" && hold.accType == "s" && hold.accBal == 243.50 {
		fmt.Println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
func TestUpdateAccountSlice(t *testing.T) {
	var holder []Account
	hold := CreateAccount("2", "1", "s", 243.50)
	hold2 := CreateAccount("2", "1", "s", 243.50)
	tester := CreateAccount("2", "1", "s", 500)
	holder = append(holder, hold)
	holder = append(holder, hold2)
	holder = UpdateAccountSlice(holder, tester, 0)
	if holder[0].accBal == 500 {
		fmt.Println("Success")
	} else {
		fmt.Println("Error happened")
	}
}
