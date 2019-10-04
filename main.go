package main

import (
	"fmt"
	"strconv"
)

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

func main() {

	var admin accountHolder = accountHolder{21933, "Carol", "Jenkins", 15942.93}
	fmt.Println(admin.Info())

}
