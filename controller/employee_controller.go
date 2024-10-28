package controller

import (
	"fmt"
	"go-postgres-test-1/model"
	"go-postgres-test-1/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController interface {
	CreateEmployee(ctx *gin.Context)
	GetAllEmployees(ctx *gin.Context)
	GetAllEmployeesWithEmployer(ctx *gin.Context)
	GetEmployee(ctx *gin.Context)
	GetEmployeeWithEmployer(ctx *gin.Context)
	UpdateEmployee(ctx *gin.Context)
	DeleteEmployee(ctx *gin.Context)
}

type employeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) EmployeeController {
	return &employeeController{service: service}
}

func (controller *employeeController) CreateEmployee(ctx *gin.Context) {
	var employee model.NewEmployeeRequest
	err := ctx.ShouldBindJSON(&employee)
	if err != nil {
		message := "Name, Age and Salary are required and cannot be empty"
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": message,
		})
		return
	}

	serviceErr := controller.service.CreateEmployee(employee)
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Employee saved successfully",
		})
	}
}

func (controller *employeeController) GetAllEmployees(ctx *gin.Context) {
	employees, serviceErr := controller.service.GetAllEmployees()
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "Employees fetched successfully",
			"employees": employees,
		})
	}
}

func (controller *employeeController) GetAllEmployeesWithEmployer(ctx *gin.Context) {
	employees, serviceErr := controller.service.GetAllEmployeesWithEmployer()
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "Employees fetched successfully",
			"employees": employees,
		})
	}
}

func (controller *employeeController) GetEmployee(ctx *gin.Context) {
	employeeId := ctx.Param("id")
	id, err := strconv.ParseUint(employeeId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee-id"})
		return
	}

	employee, serviceErr := controller.service.GetEmployee(uint(id))
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"message":  "Employee fetched successfully",
			"employee": employee,
		})
	}
}

func (controller *employeeController) GetEmployeeWithEmployer(ctx *gin.Context) {
	employeeId := ctx.Param("id")
	id, err := strconv.ParseUint(employeeId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee-id"})
		return
	}

	employee, serviceErr := controller.service.GetEmployeeWithEmployer(uint(id))
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"message":  "Employees fetched successfully",
			"employee": employee,
		})
	}
}

func (controller *employeeController) UpdateEmployee(ctx *gin.Context) {
	employeeId := ctx.Param("id")
	id, err := strconv.ParseUint(employeeId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee-id"})
		return
	}

	var employee model.UpdateEmployeeRequest
	err = ctx.ShouldBindJSON(&employee)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("Log Update Employee Controller", employee)

	if employee.Name == nil && employee.Age == nil && employee.Salary == nil && employee.EmployerId == nil && employee.UnsetEmployer == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Employee updated successfully",
		})
		return
	}

	serviceErr := controller.service.UpdateEmployee(uint(id), employee)
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Employee updated successfully",
		})
	}
}

func (controller *employeeController) DeleteEmployee(ctx *gin.Context) {
	employeeId := ctx.Param("id")
	id, err := strconv.ParseUint(employeeId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee-id"})
		return
	}

	serviceErr := controller.service.DeleteEmployee(uint(id))
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Employee deleted successfully",
		})
	}

}
