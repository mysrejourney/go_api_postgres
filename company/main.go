package main

import (
	"company/database"
	"company/router"
	"fmt"
	"net/http"
)

// Script execution starts here
func main() {
	// Making database connection
	dbConnection, _ := database.InitializeDatabaseConnection()

	// Initiating router and passing database connection object as an argument
	ginRouterEngine := router.Init(dbConnection)

	// Running the router engine at port 8080
	routerStartingError := http.ListenAndServe(":8080", ginRouterEngine)
	if routerStartingError != nil {
		fmt.Println("Unable to start server, error: ", routerStartingError)
	}
	fmt.Println("Server started at port 8080")
}
