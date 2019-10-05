package acc

import (
	"fmt"
	"log"
	"testing"
)

// func TestAccountHolder_Info(t *testing.T) {

// }

// func TestCreateAccount(t *testing.T) {
// 	type args struct {
// 		counter        int
// 		fname          string
// 		lname          string
// 		initialdeposit float64
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want AccountHolder
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := CreateAccount(tt.args.counter, tt.args.fname, tt.args.lname, tt.args.initialdeposit); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("CreateAccount() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestWithdraw(t *testing.T) {
	a := CreateAccount(5, "john", "johnson", "jjohnson", "password", 5000)
	a.accountBal = 100
	a.Withdraw(110)
	if a.accountBal == 100 {
		log.Println("withdrawing works")
		fmt.Println(a.accountBal)
	} else {
		t.Error("Expected 90 but got :", a.accountBal)
	}
}

func TestDeposit(t *testing.T) {
	a := CreateAccount(5, "john", "johnson", "jjohnson", "password", 5000)
	a.Deposit(50)
	if a.accountBal == 5050 {
		fmt.Println("deposit works well")
	} else {
		t.Error("Check your code")
	}
}

func TestTransferFunds(t *testing.T) {

}
