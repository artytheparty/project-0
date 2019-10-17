package applications

import "testing"

func TestCreateAccountHolder(t *testing.T) {
	hold := CreateAccountHolder("username", "password", "fname", "lname", "s", 5000)
	if hold.username == "username" && hold.password == "password" && hold.fname == "fname" && hold.lname == "lname" && hold.types == "s" && hold.bal == 5000 {
		println("Success")
	} else {
		println("Error Occured")
	}
}
func TestAccountHolderInfo2(t *testing.T) {
	hold := CreateAccountHolder("username", "password", "fname", "lname", "s", 5000)
	if hold.AccountHolderInfo2() == "USERNAME: username PASSWORD: password FNAME: fname LNAME: lname TYPE: s BALANCE: 5000.00" {
		println("Success")
	} else {
		println("Error Occured")
	}
}
