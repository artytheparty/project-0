package dbhandling

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/artytheparty/project-0/acc"
)

//WriteToFile creates the file
func WriteToFile(c map[int]acc.AccountHolder) {
	f, err := os.Create("accounts.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i < len(c)+1; i++ {
		f.WriteString((c[i]).Info())
	}
	err = f.Close()

}

//ReadFile will access the accounts db
func ReadFile() {
	//holder := make(map[string]acc.AccountHolder)
	f, err := os.Open("accounts.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewReader(f)
	var txtline string
	for {
		txtline, err = scanner.ReadString('\n')
		fmt.Println(txtline)
		uname := txtline[9:strings.Index(txtline, "Password: ")]
		pass := txtline[strings.Index(txtline, "Password: ")+9 : strings.Index(txtline, "Last Name: ")]
		lname := txtline[strings.Index(txtline, "Last Name: ")+10 : strings.Index(txtline, "First Name: ")]
		fname := txtline[strings.Index(txtline, "First Name: ")+11 : strings.Index(txtline, "Account Number: ")]
		aNum := txtline[strings.Index(txtline, "Account Number: ")+15 : strings.Index(txtline, "Account Balance: ")]
		aBal, _ := strconv.ParseFloat(txtline[strings.Index(txtline, "Account Balance: ")+17:len(txtline)-1], 64)
		fmt.Println(uname, pass, lname, fname, aNum, aBal)
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf("failed: %v\n", err)
	}
	return
	//fmt.Println("error2")
	//f.Close()
	//fmt.Println(txtline)

}

// func main() {
// 	fmt.Println("I handle files")
// }
