package controller

import (
	"fmt"
	"go-postgres-test-1/model"
	"go-postgres-test-1/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{service: service}
}

func (controller *userController) CreateUser(ctx *gin.Context) {
	var user model.NewUserRequest
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		message := "Username, Email and Password are required and cannot be empty"
		fmt.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": message,
		})
		return
	}

	serviceErr := controller.service.CreateUser(user)
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User saved successfully",
		})
	}
}

func (controller *userController) GetAllUsers(ctx *gin.Context) {
	users, serviceErr := controller.service.GetAllUsers()
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Users fetched successfully",
			"users":   users,
		})
	}
}

func (controller *userController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user-id"})
		return
	}

	user, serviceErr := controller.service.GetUser(uint(id))
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User fetched successfully",
			"user":    user,
		})
	}
}

func (controller *userController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user-id"})
		return
	}

	var user model.UpdateUserRequest
	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if user.Email == "" && user.Password == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User updated successfully",
		})
		return
	}

	serviceErr := controller.service.UpdateUser(uint(id), user)
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User updated successfully",
		})
	}
}

func (controller *userController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user-id"})
		return
	}

	serviceErr := controller.service.DeleteUser(uint(id))
	if serviceErr.StatusCode != 0 {
		ctx.JSON(serviceErr.StatusCode, gin.H{
			"status":  serviceErr.StatusCode,
			"message": serviceErr.Message,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User deleted successfully",
		})
	}
}
