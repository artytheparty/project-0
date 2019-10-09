package dbhandling

import (
	"database/sql"
	"fmt"
)

func getAllCustomers(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM CUSTOMERS")
	for rows.Next() {
		var username string
		var pass string
		var fName string
		var lName string
		var aNum string
		var accBal float64
		rows.Scan(&username, &pass, &fName, &lName, &aNum, &accBal)
		fmt.Println(username, pass, fName, lName, aNum, accBal)
	}
}
func updateAccountBal(acc accountHolder, db *sql.DB) {

}
