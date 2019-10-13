package emp

import "database/sql"

//Employee structure holds the values needed to edit accounts
type Employee struct {
	empID       string
	empUsername string
	empPass     string
	empFName    string
	empLName    string
}

//CreateNewEmployee creates a new employee struct and returns it
func CreateNewEmployee(id string, username string, pass string, fName string, lName string) Employee {
	return Employee{id, username, pass, fName, lName}
}

//GetEmployeeID returns the employee ID
func (a *Employee) GetEmployeeID() string {
	return a.empID
}

//GetEmployeeUsername returns the employee Username
func (a *Employee) GetEmployeeUsername() string {
	return a.empUsername
}

//GetEmployeePass returns the employee Password
func (a *Employee) GetEmployeePass() string {
	return a.empPass
}

//GetEmployeeFName returns the employee FName
func (a *Employee) GetEmployeeFName() string {
	return a.empFName
}

//GetEmployeeLName returns the employee LName
func (a *Employee) GetEmployeeLName() string {
	return a.empLName
}

//GetEmployeeInfo returns the Employee struct from DB
func GetEmployeeInfo(usrname string, db *sql.DB) Employee {
	row := db.QueryRow("SELECT * FROM employees WHERE empid=$1", usrname)
	var id string
	var username string
	var pass string
	var fName string
	var lName string
	row.Scan(&id, &username, &pass, &fName, &lName)
	return CreateNewEmployee(id, username, pass, fName, lName)

}
