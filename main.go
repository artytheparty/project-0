package main

import (
	"fmt"
	"strconv"
)

var counter int = 1

//accountHolder Struct
type accountHolder struct {
	accountNum   int
	fName, lname string
	accountBal   float64
}

func (a accountHolder) Info() string {
	return ("Account Number: " + strconv.Itoa(a.accountNum) +
		"\nLast Name: " + a.lname +
		"\nFirst Name: " + a.fName +
		"\nAccount balance: " + strconv.FormatFloat(a.accountBal, 'f', 2, 64))
}

func createAccount(fname string, lname string, initialdeposit float64) accountHolder {
	return accountHolder{counter, fname, lname, initialdeposit}
}

func main() {

	var admin accountHolder = accountHolder{21933, "Carol", "Jenkins", 15942.93}
	fmt.Println(admin.Info())
	var customerList = make(map[int]accountHolder)
	customerList[admin.accountNum] = admin
	fmt.Println(customerList)
}
