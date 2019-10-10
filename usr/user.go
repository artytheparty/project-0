package usr

import (
	_ "fmt"

	"github.com/artytheparty/project-0/acc"
)

//User Structure
type User struct {
	id, username, password, fName, lName string
	accounts                             []acc.Account
}
