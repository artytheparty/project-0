package bank

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/artytheparty/project-0/usr"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func TestDeposit(t *testing.T) {
	//opening conncetion to the database
	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataB, err := sql.Open("postgres", connecDB)
	defer dataB.Close()
	if err != nil {
		panic(err)
	}
	var tempUSR usr.User = GetUsrInfo("akhv", dataB)
	fmt.Println(tempUSR)
	//if (tempUSR.username) == "akhv" {

	//}
}
