package repository

import (
	"errors"
	"fmt"
	"go-postgres-test-1/model"
	"net/http"

	"gorm.io/gorm"
)

type EmployerRepository interface {
	CreateEmployer(name, location string) (err model.Error)
	GetAllEmployers() (employers []model.Employer, err model.Error)
	GetEmployer(id uint) (employer model.Employer, err model.Error)
	UpdateEmployer(id uint, newName, newLocation string) (err model.Error)
	DeleteEmployer(id uint) (err model.Error)
}

type employerRepository struct {
	db *gorm.DB
}

func NewEmployerRepository(db *gorm.DB) EmployerRepository {
	return &employerRepository{db: db}
}

func (r *employerRepository) CreateEmployer(name, location string) (err model.Error) {
	employer := model.Employer{Name: name, Location: location}
	result := r.db.Create(&employer)
	if result.Error != nil {
		fmt.Println("Error creating employer:", result.Error.Error())
		return model.Error{StatusCode: http.StatusInternalServerError, Message: "Error creating employer"}
	}
	return err
}

func (r *employerRepository) GetAllEmployers() (employers []model.Employer, err model.Error) {
	result := r.db.Preload("Employees").Find(&employers)
	if result.Error != nil {
		fmt.Println("Error fetching employers:", result.Error.Error())
		return []model.Employer{}, model.Error{StatusCode: http.StatusInternalServerError, Message: "Error fetching employers"}
	}
	return employers, err
}

func (r *employerRepository) GetEmployer(id uint) (employer model.Employer, err model.Error) {
	result := r.db.Preload("Employees").First(&employer, id)
	if result.Error != nil {
		fmt.Println("Error fetching employer:", result.Error.Error())
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Employer{}, model.Error{StatusCode: http.StatusNotFound, Message: "Employer not found"}
		} else {
			return model.Employer{}, model.Error{StatusCode: http.StatusInternalServerError, Message: "Error fetching employer"}
		}
	}
	return employer, err
}

func (r *employerRepository) UpdateEmployer(id uint, newName, newLocation string) (err model.Error) {
	employer, err := r.GetEmployer(id)
	if err.StatusCode != 0 {
		return err
	}

	if newName != "" {
		employer.Name = newName
	}
	if newLocation != "" {
		employer.Location = newLocation
	}

	result := r.db.Save(&employer)
	if result.Error != nil {
		fmt.Println("Error updating employer:", result.Error.Error())
		return model.Error{StatusCode: http.StatusInternalServerError, Message: "Error updating employer"}
	}
	return err
}

func (r *employerRepository) DeleteEmployer(id uint) (err model.Error) {
	_, err = r.GetEmployer(id)
	if err.StatusCode != 0 {
		return err
	}

	result := r.db.Delete(&model.Employer{}, id)
	if result.Error != nil {
		fmt.Println("Error deleting employer:", result.Error.Error())
		return model.Error{StatusCode: http.StatusInternalServerError, Message: "Error deleting employer"}
	}
	return err
}
