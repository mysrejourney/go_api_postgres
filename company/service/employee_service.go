package service

import (
	"company/mapper"
	"company/model/api"
	"company/model/db"
	"company/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// EmployeeServiceInterface : Defining interface for all API calls and get the result from db and convert it to API model

type EmployeeServiceInterface interface {
	GetEmployees(ctx *gin.Context) (api.GetEmployeeResponseAPIModel, error)
	InsertEmployee(ctx *gin.Context, empServiceInsert api.EmployeeAPIModelStruct) error // Insert employee details into table in db
	UpdateEmployee(ctx *gin.Context, empServiceDelete api.EmployeeAPIModelStruct) error // Update employee details into table in db
	DeleteEmployee(ctx *gin.Context, empServiceDelete api.EmployeeAPIModelStruct) error // Delete employee details into table in db
}

// EmployeeServiceStruct : Defining struct datatype for repository interface
type EmployeeServiceStruct struct {
	EmployeeServiceInterfaceToRepository repository.EmployeeRepositoryInterface
}

/*
// function name	: EmployeeServiceToRouter
// arguments		: repository interface
// return			: interface values for EmployeeServiceInterface
*/

func EmployeeServiceToRouter(employeeRepo repository.EmployeeRepositoryInterface) EmployeeServiceInterface {
	return EmployeeServiceStruct{EmployeeServiceInterfaceToRepository: employeeRepo}
}

/*
// function name	: GetEmployees
// arguments		: router engine object
// return			: employees details from db and convert it to API model
*/

func (empServiceStruct EmployeeServiceStruct) GetEmployees(context *gin.Context) (api.GetEmployeeResponseAPIModel, error) {
	fmt.Println("We are in service layer")
	// employeeDetailsService is a variable and store employees details from DB and convert them in the form of API Model
	employeeDetailsService, err := empServiceStruct.EmployeeServiceInterfaceToRepository.GetEmployees(context)

	// If any failures in calling to repository interface, then error is not nil.
	if err != nil {
		fmt.Println("Error occurred while executing getEmployees in the service call", err)
		return api.GetEmployeeResponseAPIModel{}, err
	}

	// If success in calling to repository interface, then error is nil.
	//Then return the employee details in the form of API Model
	return mapper.FromDMEmployeeModelToAPIModel(employeeDetailsService), nil
}

/*
// function name	: InsertEmployee
// arguments		: router engine object
// return			: employees details from db and convert it to API model
*/

func (empServiceStruct EmployeeServiceStruct) InsertEmployee(ctx *gin.Context, empServiceInsertRecord api.EmployeeAPIModelStruct) error {
	fmt.Println("We are in service layer")
	formattedJoinDate, _ := time.Parse("2006-01-02", empServiceInsertRecord.APIJoinDate)

	employeeServiceInsertEmployee := db.EmployeeDBModelStruct{Id: empServiceInsertRecord.APIId, Name: empServiceInsertRecord.APIName,
		Age: empServiceInsertRecord.APIAge, Address: empServiceInsertRecord.APIAddress,
		Salary: empServiceInsertRecord.APISalary, JoinDate: formattedJoinDate}

	errInsertService := empServiceStruct.EmployeeServiceInterfaceToRepository.InsertEmployee(ctx, employeeServiceInsertEmployee)

	if errInsertService == nil {
		fmt.Println("Employee record is inserted in service layer!")
		return nil
	} else {
		fmt.Printf("Error occurred while inserting employee record in service layer. Error: %v\n", errInsertService.Error())
		return errInsertService
	}

}

/*
// function name	: UpdateEmployee
// arguments		: router engine object
// return			: employees details from db and convert it to API model
*/

func (empServiceStruct EmployeeServiceStruct) UpdateEmployee(ctx *gin.Context, empServiceUpdateRecord api.EmployeeAPIModelStruct) error {
	fmt.Println("We are in service layer")
	formattedJoinDate, _ := time.Parse("2006-01-02", empServiceUpdateRecord.APIJoinDate)
	employeeServiceUpdateEmployee := db.EmployeeDBModelStruct{Id: empServiceUpdateRecord.APIId, Name: empServiceUpdateRecord.APIName,
		Age: empServiceUpdateRecord.APIAge, Address: empServiceUpdateRecord.APIAddress,
		Salary: empServiceUpdateRecord.APISalary, JoinDate: formattedJoinDate}

	errUpdateService := empServiceStruct.EmployeeServiceInterfaceToRepository.UpdateEmployee(ctx, employeeServiceUpdateEmployee)

	if errUpdateService == nil {
		fmt.Println("Employee record is updated in service layer!")
		return nil
	} else {
		fmt.Printf("Error occurred while updating employee record in service layer. Error: %v\n", errUpdateService.Error())
		return errUpdateService
	}

}

/*
// function name	: DeleteEmployee
// arguments		: router engine object
// return			: employees details from db and convert it to API model
*/

func (empServiceStruct EmployeeServiceStruct) DeleteEmployee(ctx *gin.Context, empServiceDeleteRecord api.EmployeeAPIModelStruct) error {
	fmt.Println("We are in service layer")
	employeeServiceDeleteEmployee := db.EmployeeDBModelStruct{Id: empServiceDeleteRecord.APIId}

	errDeleteService := empServiceStruct.EmployeeServiceInterfaceToRepository.DeleteEmployee(ctx, employeeServiceDeleteEmployee)

	if errDeleteService == nil {
		fmt.Println("Employee record is deleted in service layer!")
		return nil
	} else {
		fmt.Printf("Error occurred while deleting employee record in service layer. Error: %v\n", errDeleteService.Error())
		return errDeleteService
	}

}
