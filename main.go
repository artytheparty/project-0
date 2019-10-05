package main

import (
	"fmt"

	"github.com/artytheparty/project-0/acc"
)

func main() {
	fmt.Println("Welcome to the Banking App")
	admin := acc.CreateAccount(5, "john", "johnson", "jjohnson", "password", 5000)
	fmt.Println(admin.Info())
}
