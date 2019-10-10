package acc

import(
	"github.com/artytheparty/project-0/usr"
)

//Account Structure
type Account struct {
	accNum  string
	accType rune
	accBal  float64
	accID   string
}

func createAccount(a usr.User, typ rune, initialDeposit float64) Account{
	accNum := method to figure out accountNum
	accType := typ
	accBal := iniinitialDeposit
	accId := a.id
}