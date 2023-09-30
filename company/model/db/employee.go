package db

import (
	"database/sql"
	"gopkg.in/guregu/null.v3"
	_ "gopkg.in/guregu/null.v3"
	"time"
)

// EmployeeDBModelStruct : Defining struct for one employee DB model

type NullInt64 sql.NullInt64

//type EmployeeDBModelStruct struct {
//	Id       int             `db:"id"`
//	Name     string          `db:"name"`
//	Age      sql.NullInt64   `db:"age"`
//	Address  sql.NullString  `db:"address"`
//	Salary   sql.NullFloat64 `db:"salary"`
//	JoinDate time.Time       `db:"join_date"`
//}

//type EmployeeDBModelStruct struct {
//	Id       int       `db:"id"`
//	Name     string    `db:"name"`
//	Age      int       `db:"age"`
//	Address  string    `db:"address"`
//	Salary   string    `db:"salary"`
//	JoinDate time.Time `db:"join_date"`
//}

type EmployeeDBModelStruct struct {
	Id       int         `db:"id"`
	Name     string      `db:"name"`
	Age      null.Int    `db:"age"`
	Address  null.String `db:"address"`
	Salary   null.Float  `db:"salary"`
	JoinDate time.Time   `db:"join_date"`
}

// MultipleEmployeesDBModel : Defining struct multiple employees for API model

type MultipleEmployeesDBModel []EmployeeDBModelStruct
