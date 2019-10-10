package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//database constants
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	connecDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dataB, err := sql.Open("postgres", connecDB)
	defer dataB.Close()
	if err != nil {
		panic(err)
	}
}
