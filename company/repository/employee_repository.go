package repository

import (
	"company/model/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// EmployeeRepositoryInterface : Defining interface for all API calls
type EmployeeRepositoryInterface interface {
	GetEmployees(ctx *gin.Context) (db.MultipleEmployeesDBModel, error) // Retrieve all employee details
}

// EmployeeRepositoryDBCollection : Defining struct datatype for db object
type EmployeeRepositoryDBCollection struct {
	EmployeeRepositoryDBCollectionObject *sqlx.DB // EmployeeRepositoryDBCollectionObject holds db object type of *sql.DB
}

/*
// function name	: EmployeeRepositoryToRouter
// arguments		: dbObject
// return			: interface values for EmployeeRepositoryInterface
*/

func EmployeeRepositoryToRouter(EmployeeRepositoryToRouterDBObject *sqlx.DB) EmployeeRepositoryInterface {
	return EmployeeRepositoryDBCollection{EmployeeRepositoryDBCollectionObject: EmployeeRepositoryToRouterDBObject}
}

/*
// function name	: GetEmployees
// arguments		: router engine object
// return			: employees details from DB in the form of DB MODEL
*/

func (empStruct EmployeeRepositoryDBCollection) GetEmployees(context *gin.Context) (db.MultipleEmployeesDBModel, error) {
	// employeeDetails is a type of MultipleEmployeesDBModel variable and store employees details from DB in the form of DB Model
	var employeeDetails db.MultipleEmployeesDBModel

	// Query to be executed in postgres database
	selectQuery := `SELECT id, name, age, address, salary, join_date from company;`

	// Execute the query and get the results
	error := empStruct.EmployeeRepositoryDBCollectionObject.SelectContext(context, &employeeDetails, selectQuery)

	// If any failures in executing the query, then error is not nil.
	if error != nil {
		fmt.Println("Error occurred while executing GetEmployees select query", error)
		return db.MultipleEmployeesDBModel{}, error // return as error
	}

	// If success in executing the query, then error is nil. Then return the employee details in the form of DB Model
	return employeeDetails, nil
}
