package applications

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/artytheparty/project-0/bank"
)

//AccountHolder will hold the necessary information to create an accoun after it is approved
type AccountHolder struct {
	username, password, fname, lname, types string
	bal                                     float64
}

//CreateAccountHolder will create the account holder
func CreateAccountHolder(username string, pass string, fname string, lname string, types string, bal float64) AccountHolder {
	return AccountHolder{username, pass, fname, lname, types, bal}
}

//AccountHolderInfo will just print the account.
func (a *AccountHolder) AccountHolderInfo() {
	fmt.Printf("USERNAME: %s PASSWORD: %s FNAME: %s LNAME: %s TYPE: %s BALANCE: %f\n",
		a.username, a.password, a.fname, a.lname, a.types, a.bal)
}

//AccountHolderInfo2 will just return the accountinfo as a string.
func (a *AccountHolder) AccountHolderInfo2() string {
	//var balance string = fmt.Sprintf("%f", a.bal)
	return ("USERNAME: " + a.username +
		" PASSWORD: " + a.password +
		" FNAME: " + a.fname +
		" LNAME: " + a.lname +
		" TYPE: " + a.types +
		" BALANCE: " + strconv.FormatFloat(a.bal, 'f', 2, 64) +
		"\n")
}

//WriteToFile will write the array of accounts to a txt file
func WriteToFile(c []AccountHolder) {
	f, err := os.Create("accounts.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range c {
		f.WriteString(c[i].AccountHolderInfo2())
	}
	err = f.Close()
}

//ReadFile will read the accounts.txt file so that we can recreate the list for the employee
func ReadFile() []AccountHolder {
	var holder []AccountHolder
	f, err := os.Open("accounts.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewReader(f)
	var txtline string
	for {
		txtline, err = scanner.ReadString('\n')
		//fmt.Println(txtline)

		if err != nil {
			break
		}
		uname := txtline[9:strings.Index(txtline, "PASSWORD: ")]
		pass := txtline[strings.Index(txtline, "PASSWORD: ")+9 : strings.Index(txtline, "FNAME: ")]
		fname := txtline[strings.Index(txtline, "FNAME: ")+6 : strings.Index(txtline, "LNAME: ")]
		lname := txtline[strings.Index(txtline, "LNAME: ")+6 : strings.Index(txtline, "TYPE: ")]
		types := txtline[strings.Index(txtline, "TYPE: ")+5 : strings.Index(txtline, "BALANCE: ")]
		aBal, _ := strconv.ParseFloat(txtline[strings.Index(txtline, "BALANCE: ")+8:len(txtline)-1], 64)
		//1fmt.Println(aBal)
		holder = append(holder, CreateAccountHolder(uname, pass, fname, lname, types, aBal))
	}
	if err != io.EOF {
		fmt.Printf("failed: %v\n", err)
	}
	return holder

}

//PrintHolder will just print out the list of acconts
func PrintHolder(a []AccountHolder) {
	for i := range a {
		fmt.Println("Temp. Acc Number", i)
		a[i].AccountHolderInfo()
	}
}

//AskNewAccount will write the map of accounts whic have aplied to open one
func AskNewAccount(db *sql.DB) AccountHolder {

	var username string
	var pass string
	var fname string
	var lname string
	var bal float64
	var types string
	fmt.Println("Enter your username: ")
	fmt.Scan(&username)
	fmt.Println("Enter your password: ")
	fmt.Scan(&pass)
	fmt.Println("Enter your First Name: ")
	fmt.Scan(&fname)
	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lname)
	fmt.Println("Enter your Account type \"c\" for checking \"s\" for savings: ")
	fmt.Scan(&types)
	fmt.Println("Enter your Initial deposit: ")
	fmt.Scan(&bal)
	return CreateAccountHolder(username, pass, fname, lname, types, bal)
}

//ApproveAndAddToDB willl ad a selected account to the db
func ApproveAndAddToDB(db *sql.DB, a []AccountHolder) {
	fmt.Println("Which account would you like to add/remove?")
	PrintHolder(a)
	fmt.Println("Enter temporary account number: ")
	var choice int
	fmt.Scan(&choice)
	fmt.Println("[r]emove? or [a]pprove]: ?")
	var choice2 string
	fmt.Scan(&choice2)
	holder := a[choice]
	if choice2 == "a" {
		bank.CreateNewUserEntry(holder.username, holder.password, holder.fname, holder.lname, db)
		usrholder := bank.GetUsrInfo(holder.username, db)
		bank.CreateNewAccountEntry(usrholder.GetUsrID(), holder.types, holder.bal, db)
		b := remove(a, choice)
		WriteToFile(b)
		fmt.Println("Successfully added the user and their account to Database")
	}
	if choice2 == "r" {
		b := remove(a, choice)
		WriteToFile(b)
		fmt.Println("Successfully removed Pending Account from file.")
	}
}
func remove(s []AccountHolder, i int) []AccountHolder {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
