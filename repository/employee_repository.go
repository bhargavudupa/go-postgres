package repository

import (
	"errors"
	"fmt"
	"go-postgres-test-1/model"
	"net/http"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(name string, age int, salary float64, employerId *uint) (err model.Error)
	GetAllEmployees() (employees []model.Employee, err model.Error)
	GetEmployee(id uint) (employee model.Employee, err model.Error)
	UpdateEmployee(id uint, newName string, newAge int, newSalary float64, newEmployerId *uint, unsetEmployer *bool) (err model.Error)
	DeleteEmployee(id uint) (err model.Error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) CreateEmployee(name string, age int, salary float64, employerId *uint) (err model.Error) {
	employee := model.Employee{Name: name, Age: age, Salary: salary, EmployerId: employerId}
	result := r.db.Create(&employee)
	if result.Error != nil {
		fmt.Println("Error creating employee:", result.Error.Error())
		return model.Error{StatusCode: http.StatusInternalServerError, Message: "Error creating employee"}
	}
	return err
}

func (r *employeeRepository) GetAllEmployees() (employees []model.Employee, err model.Error) {
	result := r.db.Preload("Employer").Find(&employees)
	if result.Error != nil {
		fmt.Println("Error fetching employees:", result.Error.Error())
		return []model.Employee{}, model.Error{StatusCode: http.StatusInternalServerError, Message: "Error fetching employees"}
	}
	return employees, err
}

func (r *employeeRepository) GetEmployee(id uint) (employee model.Employee, err model.Error) {
	result := r.db.Preload("Employer").First(&employee, id)
	if result.Error != nil {
		fmt.Println("Error fetching employee:", result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Employee{}, model.Error{StatusCode: http.StatusNotFound, Message: "Employee not found"}
		} else {
			return model.Employee{}, model.Error{StatusCode: http.StatusInternalServerError, Message: "Error fetching employee"}
		}
	}
	return employee, err
}

func (r *employeeRepository) UpdateEmployee(id uint, newName string, newAge int, newSalary float64, newEmployerId *uint, unsetEmployer *bool) (err model.Error) {
	if unsetEmployer != nil && *unsetEmployer {
		newEmployerId = nil
	}

	result := r.db.Save(&model.Employee{Id: id, Name: newName, Age: newAge, Salary: newSalary, EmployerId: newEmployerId})
	if result.Error != nil {
		fmt.Println("Error updating employee:", result.Error.Error())
		return model.Error{StatusCode: http.StatusInternalServerError, Message: "Error updating employee"}
	}
	return err
}

func (r *employeeRepository) DeleteEmployee(id uint) (err model.Error) {
	_, err = r.GetEmployee(id)
	if err.StatusCode != 0 {
		return err
	}

	result := r.db.Delete(&model.Employee{}, id)
	if result.Error != nil {
		fmt.Println("Error deleting employee:", result.Error.Error())
		return model.Error{StatusCode: http.StatusInternalServerError, Message: "Error deleting employee"}
	}
	return err
}
