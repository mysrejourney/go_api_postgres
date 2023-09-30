package mapper

import (
	"company/model/api"
	"company/model/db"
)

/*
// function name	: FromDMEmployeeModelToAPIModel
// arguments		: Collections of employee details
// return			: Return employees details in API model
*/

func FromDMEmployeeModelToAPIModel(EmployeeDetailsMapper db.MultipleEmployeesDBModel) api.GetEmployeeResponseAPIModel {
	// employeeAPIModel is a type of GetEmployeeResponseAPIModel variable and store employees details from DB in the form of API Model
	var employeeAPIModel api.GetEmployeeResponseAPIModel

	// Looping through all employees details one by one
	for _, employee := range EmployeeDetailsMapper {
		// Convert each employee details in the form of API model and append them one be one
		employeeAPIModel.MultipleEmployeesAPIModel = append(employeeAPIModel.MultipleEmployeesAPIModel, ToEmployeeAPIModel(employee))
	}

	// Return the final API model which contains all employees details
	return employeeAPIModel
}

/*
// function name	: ToEmployeeAPIModel
// arguments		: only one employee details in the form of DB model
// return			: Return employee details in API model
*/

func ToEmployeeAPIModel(EmployeeDetailsMapper db.EmployeeDBModelStruct) api.EmployeeAPIModelStruct {
	return api.EmployeeAPIModelStruct{
		APIId:       EmployeeDetailsMapper.Id,
		APIName:     EmployeeDetailsMapper.Name,
		APIAge:      EmployeeDetailsMapper.Age,
		APIAddress:  EmployeeDetailsMapper.Address,
		APISalary:   EmployeeDetailsMapper.Salary,
		APIJoinDate: EmployeeDetailsMapper.JoinDate.Format("2006-01-02"),
	}
}
