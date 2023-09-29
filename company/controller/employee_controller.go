package controller

import (
	"company/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// EmployeeControllerInterface : Defining interface for all API calls and get the result from db and convert it to API model

type EmployeeControllerInterface interface {
	GetEmployees(ctx *gin.Context)
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
