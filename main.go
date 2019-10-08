package main

import (
	"fmt"

	"github.com/artytheparty/project-0/acc"
	"github.com/artytheparty/project-0/dbhandling"
)

func main() {
	fmt.Println("Welcome to the Banking App")
	admin := acc.CreateAccount("arty", "password", "arty", "sucess", 1, 5000)
	admin2 := acc.CreateAccount("jjohna", "password", "john", "johna", 1, 5000)
	admin3 := acc.CreateAccount("usernam", "password", "john", "johna", 1, 5000)
	admin4 := acc.CreateAccount("shword", "password", "john", "johna", 1, 5000)
	admin5 := acc.CreateAccount("uthough", "password", "john", "johna", 1, 5000)
	admin6 := acc.CreateAccount("arty", "password", "arty", "sucess", 1, 5000)
	admin7 := acc.CreateAccount("jjohna", "password", "john", "johna", 1, 5000)
	admin8 := acc.CreateAccount("usernam", "password", "john", "johna", 1, 5000)
	admin9 := acc.CreateAccount("shword", "password", "john", "johna", 1, 5000)
	admin10 := acc.CreateAccount("uthough", "password", "john", "johna", 1, 5000)

	//fmt.Println(admin.Info())
	accounts := make(map[int]acc.AccountHolder)
	accounts[1] = admin
	accounts[2] = admin2
	accounts[3] = admin3
	accounts[4] = admin4
	accounts[5] = admin5
	accounts[6] = admin6
	accounts[7] = admin7
	accounts[8] = admin8
	accounts[9] = admin9
	accounts[10] = admin10

	//dbhandling.WriteToFile(accounts)
	accounts2 := dbhandling.ReadFile()
	for k := range accounts2 {
		accounts2[k].Info()
	}

}
