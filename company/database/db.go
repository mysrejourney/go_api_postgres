package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

/*
// function name	: getEnvironmentVariableWithDefault
// arguments		: currentValueFromEnvVariable, newValueToBeSet
// return			: either value from env variable or new value which we passed
*/

func getEnvironmentVariableWithDefault(key string, defaultValue string) string {
	currentValue := os.Getenv(key)
	if currentValue == "" {
		return defaultValue
	} else {
		return currentValue
	}
}

/*
// function name	: ConstructDsn
// arguments		: none
// return			: Connection string to connect database
*/

func ConstructDsn() string {
	// host variable is to store hostname
	host := getEnvironmentVariableWithDefault("host", "localhost")

	// user variable is to store username
	user := getEnvironmentVariableWithDefault("user", "postgres")

	// password variable is to store password
	password := getEnvironmentVariableWithDefault("password", "Secret")

	// dbname variable is to store dbname
	dbname := getEnvironmentVariableWithDefault("dbname", "testdb")

	// port variable is to store port number
	port := getEnvironmentVariableWithDefault("port", "5432")

	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	// form the connection string
	connectionString := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
	fmt.Println("DSN : ", connectionString)
	return connectionString
}

/*
// function name	: InitializeDatabaseConnection
// arguments		: none
// return			: database connection object
*/

func InitializeDatabaseConnection() (*sqlx.DB, error) {
	var dbObject *sqlx.DB
	var err error
	// Below is for Postgres
	//dbObject, err = sqlx.Open("postgres", ConstructDsn())
	// Below is for mysql
	dbObject, err = sqlx.Open("mysql", "root:Secret@tcp(localhost:3306)/testdb?parseTime=true")
	return dbObject, err
}
