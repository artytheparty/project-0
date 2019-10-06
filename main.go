package main

import (
	"fmt"

	"github.com/artytheparty/project-0/acc"
	"github.com/artytheparty/project-0/dbhandling"
)

func main() {
	fmt.Println("Welcome to the Banking App")
	admin := acc.CreateAccount(1, "john", "johnson", "jjohnson", "password", 5000)
	admin2 := acc.CreateAccount(2, "john", "petersson", "jjohnson", "password", 54000)
	admin3 := acc.CreateAccount(3, "john", "johnson", "jjohnson", "password", 500530)
	admin4 := acc.CreateAccount(4, "john", "johnson", "jjohnson", "password", 300)
	admin5 := acc.CreateAccount(5, "john", "johnson", "jjohnson", "password", 20)

	fmt.Println(admin.Info())
	accounts := make(map[int]acc.AccountHolder)
	accounts[1] = admin
	accounts[2] = admin2
	accounts[3] = admin3
	accounts[4] = admin4
	accounts[5] = admin5

	dbhandling.WriteToFile(accounts)

}
