package usr

import "github.com/artytheparty/project-0/acc"

//User Structure
type User struct {
	id, username, password, fName, lName string
	accounts                             []acc.Account
}

//CreateUser fdsfds
func CreateUser(id string, username string, password string, fName string, lName string, accounts []acc.Account) User {
	return User{id, username, password, fName, lName, accounts}
}

//GetAccounts returns the array of user accounts.
func (a *User) GetAccounts() []acc.Account {
	return a.accounts
}
