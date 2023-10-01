package repository

import (
	"company/model/db"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// EmployeeRepositoryInterface : Defining interface for all API calls
type EmployeeRepositoryInterface interface {
	GetEmployees(ctx *gin.Context) (db.MultipleEmployeesDBModel, error)              // Retrieve all employee details
	InsertEmployee(ctx *gin.Context, empInsertRecord db.EmployeeDBModelStruct) error // Insert employee details into table in db
	UpdateEmployee(ctx *gin.Context, empUpdateRecord db.EmployeeDBModelStruct) error // Update employee details into table in db
	DeleteEmployee(ctx *gin.Context, empDeleteRecord db.EmployeeDBModelStruct) error // Delete employee details into table in db
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
	fmt.Println("We are in repository layer")
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

/*
// function name	: InsertEmployee
// arguments		: router engine object
// return			: employees details from DB in the form of DB MODEL
*/

func (empStruct EmployeeRepositoryDBCollection) InsertEmployee(context *gin.Context, insertEmployeeRecord db.EmployeeDBModelStruct) error {
	// Query to be executed in postgres database
	// Below is for Postgres
	//insertQuery := `INSERT INTO company(id, name, age, address, salary, join_date) VALUES ($1, $2, $3, $4, $5, $6)`
	// Below is for mysql
	insertQuery := `INSERT INTO company(id, name, age, address, salary, join_date) VALUES (?, ?, ?, ?, ?, ?)`

	transaction, errBegin := empStruct.EmployeeRepositoryDBCollectionObject.Beginx()
	if errBegin != nil {
		fmt.Println("Error encountered while beginning sql transaction. Error: ", errBegin.Error())
		return errBegin
	}

	// Execute the query and get the results
	result, errorInsert := transaction.ExecContext(context, insertQuery, insertEmployeeRecord.Id, insertEmployeeRecord.Name, insertEmployeeRecord.Age, insertEmployeeRecord.Address, insertEmployeeRecord.Salary, insertEmployeeRecord.JoinDate)

	defer func() {
		if errorInsert != nil {
			fmt.Println("Rolling back changes, due to error: ", errorInsert)
			_ = transaction.Rollback()
			return
		}
		fmt.Println("Successfully inserted employee record, committing the transaction")
		_ = transaction.Commit()

	}()
	if errorInsert != nil {
		fmt.Println("Error encountered inserting employee record in repository layer. Error: ", errorInsert.Error())
		return errorInsert
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Number of rows affected: %v: ", rowsAffected)
	if rowsAffected == 0 {
		fmt.Println("Failed to insert employee record in db. no rows affected in repository layer")
		return errorInsert
	}
	return nil
}

/*
// function name	: UpdateEmployee
// arguments		: router engine object
// return			: employees details from DB in the form of DB MODEL
*/

func (empStruct EmployeeRepositoryDBCollection) UpdateEmployee(context *gin.Context, updateEmployeeRecord db.EmployeeDBModelStruct) error {
	fmt.Printf("Name: %v: ", updateEmployeeRecord.Name)
	fmt.Printf("Age: %v: ", updateEmployeeRecord.Age)
	fmt.Printf("Address: %v: ", updateEmployeeRecord.Address)
	fmt.Printf("Salary: %v: ", updateEmployeeRecord.Salary)
	fmt.Printf("Join Date: %v: ", updateEmployeeRecord.JoinDate)
	// Query to be executed in postgres database
	//Below is for Postgres
	//updateQuery := `UPDATE company SET name=$2, age=$3, address=$4, salary=$5, join_date=$6 WHERE id=$1`
	//Below is for mysql
	updateQuery := `UPDATE company SET name=?, age=?, address=?, salary=?, join_date=? WHERE id=?`

	transaction, errBegin := empStruct.EmployeeRepositoryDBCollectionObject.Beginx()
	if errBegin != nil {
		fmt.Println("Error encountered while beginning sql transaction. Error: ", errBegin.Error())
		return errBegin
	}

	// Execute the query and get the results
	//Below is for Postgres
	//result, errorUpdate := transaction.ExecContext(context, updateQuery, updateEmployeeRecord.Id, updateEmployeeRecord.Name, updateEmployeeRecord.Age, updateEmployeeRecord.Address, updateEmployeeRecord.Salary, updateEmployeeRecord.JoinDate)
	//Below is for mysql
	result, errorUpdate := transaction.ExecContext(context, updateQuery, updateEmployeeRecord.Name, updateEmployeeRecord.Age, updateEmployeeRecord.Address, updateEmployeeRecord.Salary, updateEmployeeRecord.JoinDate, updateEmployeeRecord.Id)
	defer func() {
		if errorUpdate != nil {
			fmt.Println("Rolling back changes, due to error: ", errorUpdate)
			_ = transaction.Rollback()
			return
		}
		fmt.Println("Successfully updated employee record, committing the transaction")
		_ = transaction.Commit()

	}()
	if errorUpdate != nil {
		fmt.Println("Error encountered updating employee record in repository layer. Error: ", errorUpdate.Error())
		return errorUpdate
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Number of rows affected: %v: ", rowsAffected)
	if rowsAffected == 0 {
		fmt.Println("Failed to update employee record in db. no rows affected in repository layer")
		return errorUpdate
	}
	return nil
}

/*
// function name	: DeleteEmployee
// arguments		: router engine object
// return			: employees details from DB in the form of DB MODEL
*/

func (empStruct EmployeeRepositoryDBCollection) DeleteEmployee(context *gin.Context, deleteEmployeeRecord db.EmployeeDBModelStruct) error {
	// Query to be executed in postgres database
	//Below is for Postgres
	//deleteQuery := `DELETE FROM company WHERE id = ($1)`
	//Below is for mysql
	deleteQuery := `DELETE FROM company WHERE id = ?`

	transaction, errBegin := empStruct.EmployeeRepositoryDBCollectionObject.Beginx()
	if errBegin != nil {
		fmt.Println("Error encountered while beginning sql transaction. Error: ", errBegin.Error())
		return errBegin
	}

	// Execute the query and get the results
	result, errorDelete := transaction.ExecContext(context, deleteQuery, deleteEmployeeRecord.Id)

	defer func() {
		if errorDelete != nil {
			fmt.Println("Rolling back changes, due to error: ", errorDelete)
			_ = transaction.Rollback()
			return
		}
		fmt.Println("Successfully deleted employee record, committing the transaction")
		_ = transaction.Commit()

	}()
	if errorDelete != nil {
		fmt.Println("Error encountered deleting employee record in repository layer. Error: ", errorDelete.Error())
		return errorDelete
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Number of rows affected: %v: ", rowsAffected)
	if rowsAffected == 0 {
		fmt.Println("Failed to delete employee record in db. no rows affected in repository layer")
		return errorDelete
	}
	return nil
}
