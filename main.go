package main

import (
	"fmt"
	"go-postgres-test-1/controller"
	"go-postgres-test-1/repository"
	"go-postgres-test-1/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := getDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to GORM:", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUser)
	router.GET("/users", userController.GetAllUsers)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)

	employerRepo := repository.NewEmployerRepository(db)
	employerService := service.NewEmployerService(employerRepo)
	employerController := controller.NewEmployerController(employerService)

	router.POST("/employers", employerController.CreateEmployer)
	router.GET("/employers/:id", employerController.GetEmployer)
	router.GET("/employers/:id/verbose", employerController.GetEmployerWithEmployees)
	router.GET("/employers", employerController.GetAllEmployers)
	router.GET("/employers/verbose", employerController.GetAllEmployersWithEmployees)
	router.PUT("/employers/:id", employerController.UpdateEmployer)
	router.DELETE("/employers/:id", employerController.DeleteEmployer)

	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo, employerRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	router.POST("/employees", employeeController.CreateEmployee)
	router.GET("/employees/:id", employeeController.GetEmployee)
	router.GET("/employees/:id/verbose", employeeController.GetEmployeeWithEmployer)
	router.GET("/employees", employeeController.GetAllEmployees)
	router.GET("/employees/verbose", employeeController.GetAllEmployeesWithEmployer)
	router.PUT("/employees/:id", employeeController.UpdateEmployee)
	router.DELETE("/employees/:id", employeeController.DeleteEmployee)

	router.Run(":8080")
}

func getDSN() (dsn string) {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	sslmode := os.Getenv("DATABASE_SSLMODE")

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	return dsn
}
