package main

import (
	"fmt"

	"github.com/artytheparty/project-0/acc"
	"github.com/artytheparty/project-0/dbhandling"
)

func main() {
	fmt.Println("Welcome to the Banking App")
	admin := acc.CreateAccount("kjonah", "password", "Kyle", "Johna", 1, 5000)
	admin2 := acc.CreateAccount("jjohna", "password", "John", "Johna", 2, 15000)
	admin3 := acc.CreateAccount("ukra", "password", "Ursula", "Kung", 3, 5000)
	admin4 := acc.CreateAccount("shword", "password", "Shelly", "Wordinson", 4, 5000)
	admin5 := acc.CreateAccount("uthough", "password", "Umelia", "Thoughtful", 5, 5000)
	admin6 := acc.CreateAccount("kps13", "password", "Kong", "Pierce", 6, 5000)
	admin7 := acc.CreateAccount("bhide", "password", "Robert", "Hide", 7, 5000)
	admin8 := acc.CreateAccount("dlecious", "password", "David", "Lecious", 8, 5000)
	admin9 := acc.CreateAccount("hparker", "password", "Hong", "Parker", 9, 5000)
	admin10 := acc.CreateAccount("golang", "password", "Godin", "Lang", 10, 5000)

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

	dbhandling.WriteToFile(accounts)
	// accounts2 := dbhandling.ReadFile()
	// for k := range accounts2 {
	// 	accounts2[k].Info()
	// }

}
