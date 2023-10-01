package controller

import (
	"company/model/api"
	"company/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EmployeeControllerInterface : Defining interface for all API calls and get the result from db and convert it to API model

type EmployeeControllerInterface interface {
	GetEmployees(ctx *gin.Context)
	InsertEmployee(ctx *gin.Context)
	UpdateEmployee(ctx *gin.Context)
	DeleteEmployee(ctx *gin.Context)
}

// EmployeeControllerStruct : Defining struct datatype for service interface

type EmployeeControllerStruct struct {
	EmployeeControllerInterfaceToService service.EmployeeServiceInterface
}

/*
// function name	: EmployeeControllerToRouter
// arguments		: service interface
// return			: interface values for EmployeeControllerInterface
*/

func EmployeeControllerToRouter(empService service.EmployeeServiceInterface) EmployeeControllerInterface {
	return EmployeeControllerStruct{EmployeeControllerInterfaceToService: empService}
}

/*
// function name	: GetEmployees
// arguments		: router engine object
// return			: Binding the API result and send the response code
*/

func (EmpControllerStruct EmployeeControllerStruct) GetEmployees(ctx *gin.Context) {
	// employeeDetailsController is a variable and store API result and send the http response code
	employeeDetailsController, err := EmpControllerStruct.EmployeeControllerInterfaceToService.GetEmployees(ctx)

	// If any failures in calling to service interface, then error is not nil.
	if err != nil {
		fmt.Println("We are in controller layer and getting internal server error")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	// If success in calling to service interface, then error is nil.
	//Then return the employee details and send the http response code
	ctx.JSON(http.StatusOK, employeeDetailsController)
}

/*
// function name	: InsertEmployee
// arguments		: router engine object
// return			: Binding the API result and send the response code
*/

func (EmpControllerStruct EmployeeControllerStruct) InsertEmployee(ctx *gin.Context) {

	fmt.Println("We are in controller layer!")

	var employeeControllerInsertRecord api.EmployeeAPIModelStruct
	bindErr := ctx.BindJSON(&employeeControllerInsertRecord)
	if bindErr != nil {
		fmt.Printf("Failed to read request body from request in controller layer. Error: %v", bindErr.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	err := EmpControllerStruct.EmployeeControllerInterfaceToService.InsertEmployee(ctx, employeeControllerInsertRecord)
	if err != nil {
		fmt.Printf("Failed to insert employee record in controller layer. Error: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	return
}

/*
// function name	: UpdateEmployee
// arguments		: router engine object
// return			: Binding the API result and send the response code
*/

func (EmpControllerStruct EmployeeControllerStruct) UpdateEmployee(ctx *gin.Context) {

	fmt.Println("We are in controller layer!")

	var employeeControllerUpdateRecord api.EmployeeAPIModelStruct
	bindErr := ctx.BindJSON(&employeeControllerUpdateRecord)
	if bindErr != nil {
		fmt.Printf("Failed to read request body from request in controller layer. Error: %v", bindErr.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	err := EmpControllerStruct.EmployeeControllerInterfaceToService.UpdateEmployee(ctx, employeeControllerUpdateRecord)
	if err != nil {
		fmt.Printf("Failed to update employee record in controller layer. Error: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	return
}

/*
// function name	: DeleteEmployee
// arguments		: router engine object
// return			: Binding the API result and send the response code
*/

func (EmpControllerStruct EmployeeControllerStruct) DeleteEmployee(ctx *gin.Context) {

	fmt.Println("We are in controller layer!")

	var employeeControllerDeleteRecord api.EmployeeAPIModelStruct
	bindErr := ctx.BindJSON(&employeeControllerDeleteRecord)
	if bindErr != nil {
		fmt.Printf("Failed to read request body from request in controller layer. Error: %v", bindErr.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	err := EmpControllerStruct.EmployeeControllerInterfaceToService.DeleteEmployee(ctx, employeeControllerDeleteRecord)
	if err != nil {
		fmt.Printf("Failed to delete employee record in controller layer. Error: %v", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	return
}
