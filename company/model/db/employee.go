package db

import (
	"database/sql"
	"time"
)

// EmployeeDBModelStruct : Defining struct for one employee DB model

type EmployeeDBModelStruct struct {
	APIId       int             `db:"id"`
	APIName     string          `db:"name"`
	APIAge      sql.NullInt64   `db:"age"`
	APIAddress  sql.NullString  `db:"address"`
	APISalary   sql.NullFloat64 `db:"salary"`
	APIJoinDate time.Time       `db:"join_date"`
}

// MultipleEmployeesDBModel : Defining struct multiple employees for API model

type MultipleEmployeesDBModel []EmployeeDBModelStruct
