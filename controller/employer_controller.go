package controller

import (
	"fmt"
	"go-postgres-test-1/model"
	"go-postgres-test-1/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployerController interface {
	CreateEmployer(ctx *gin.Context)
	GetAllEmployers(ctx *gin.Context)
	GetAllEmployersWithEmployees(ctx *gin.Context)
	GetEmployer(ctx *gin.Context)
	GetEmployerWithEmployees(ctx *gin.Context)
	UpdateEmployer(ctx *gin.Context)
	DeleteEmployer(ctx *gin.Context)
}

type employerController struct {
	service service.EmployerService
}

func NewEmployerController(service service.EmployerService) EmployerController {
	return &employerController{service: service}
}

func (controller *employerController) CreateEmployer(ctx *gin.Context) {
	var employer model.NewEmployerRequest
	err := ctx.ShouldBindJSON(&employer)
	if err != nil {
		message := "Name and Location are required and cannot be empty"
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": message,
		})
		return
	}

	serviceErr := controller.service.CreateEmployer(employer)
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Employer saved successfully",
		})
	}
}

func (controller *employerController) GetAllEmployers(ctx *gin.Context) {
	employers, serviceErr := controller.service.GetAllEmployers()
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "Employers fetched successfully",
			"employers": employers,
		})
	}
}

func (controller *employerController) GetAllEmployersWithEmployees(ctx *gin.Context) {
	employers, serviceErr := controller.service.GetAllEmployersWithEmployees()
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":    http.StatusOK,
			"message":   "Employers fetched successfully",
			"employers": employers,
		})
	}
}

func (controller *employerController) GetEmployer(ctx *gin.Context) {
	employerId := ctx.Param("id")
	id, err := strconv.ParseUint(employerId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employer-id"})
		return
	}

	employer, serviceErr := controller.service.GetEmployer(uint(id))
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"message":  "Employer fetched successfully",
			"employer": employer,
		})
	}
}

func (controller *employerController) GetEmployerWithEmployees(ctx *gin.Context) {
	employerId := ctx.Param("id")
	id, err := strconv.ParseUint(employerId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employer-id"})
		return
	}

	employer, serviceErr := controller.service.GetEmployerWithEmployees(uint(id))
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"message":  "Employers fetched successfully",
			"employer": employer,
		})
	}
}

func (controller *employerController) UpdateEmployer(ctx *gin.Context) {
	employerId := ctx.Param("id")
	id, err := strconv.ParseUint(employerId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employer-id"})
		return
	}

	var employer model.UpdateEmployerRequest
	err = ctx.ShouldBindJSON(&employer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if employer.Name == "" && employer.Location == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Employer updated successfully",
		})
		return
	}

	serviceErr := controller.service.UpdateEmployer(uint(id), employer)
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Employer updated successfully",
		})
	}
}

func (controller *employerController) DeleteEmployer(ctx *gin.Context) {
	employerId := ctx.Param("id")
	id, err := strconv.ParseUint(employerId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employer-id"})
		return
	}

	serviceErr := controller.service.DeleteEmployer(uint(id))
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Employer deleted successfully",
		})
	}

}
