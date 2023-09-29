package api

import (
	"database/sql"
)

// EmployeeAPIModelStruct : Defining struct for one employee API model

type EmployeeAPIModelStruct struct {
	APIId       int             `json:"id"`
	APIName     string          `json:"name"`
	APIAge      sql.NullInt64   `json:"age"`
	APIAddress  sql.NullString  `json:"address"`
	APISalary   sql.NullFloat64 `json:"salary"`
	APIJoinDate string          `json:"join_date"`
}

// MultipleEmployeesAPIModel : Defining struct multiple employees for API model
type MultipleEmployeesAPIModel []EmployeeAPIModelStruct

// GetEmployeeResponseAPIModel: Defining struct multiple employees to display in the form of json.
/*
Example:
{
  "employees": [
    {
      "id": 1,
      "name": "Paul",
      "age": {
        "Int64": 32,
        "Valid": true
      },
      "address": {
        "String": "California ",
        "Valid": true
      },
      "salary": {
        "Float64": 20000,
        "Valid": true
      },
      "join_date": "2001-07-13"
    },
 }
*/

type GetEmployeeResponseAPIModel struct {
	MultipleEmployeesAPIModel []EmployeeAPIModelStruct `json:"employees"`
}
