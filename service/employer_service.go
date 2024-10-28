package service

import (
	"go-postgres-test-1/model"
	"go-postgres-test-1/repository"
)

type EmployerService interface {
	CreateEmployer(employer model.NewEmployerRequest) (err model.Error)
	GetAllEmployers() (employers []model.EmployerResponse, err model.Error)
	GetAllEmployersWithEmployees() (employers []model.EmployerWithEmployeesResponse, err model.Error)
	GetEmployer(employerId uint) (employer model.EmployerResponse, err model.Error)
	GetEmployerWithEmployees(employerId uint) (employer model.EmployerWithEmployeesResponse, err model.Error)
	UpdateEmployer(employerId uint, employer model.UpdateEmployerRequest) (err model.Error)
	DeleteEmployer(employerId uint) (err model.Error)
}

type employerService struct {
	repo repository.EmployerRepository
}

func NewEmployerService(repo repository.EmployerRepository) EmployerService {
	return &employerService{repo: repo}
}

func (service *employerService) CreateEmployer(employer model.NewEmployerRequest) (err model.Error) {
	return service.repo.CreateEmployer(employer.Name, employer.Location)
}

func (service *employerService) GetAllEmployers() (employers []model.EmployerResponse, err model.Error) {
	allEmployers, err := service.repo.GetAllEmployers()
	if err.StatusCode != 0 {
		return employers, err
	}
	for _, emp := range allEmployers {
		employers = append(employers, model.GetEmployerResponse(emp))
	}
	return employers, err
}

func (service *employerService) GetAllEmployersWithEmployees() (employers []model.EmployerWithEmployeesResponse, err model.Error) {
	allEmployers, err := service.repo.GetAllEmployers()
	if err.StatusCode != 0 {
		return employers, err
	}
	for _, emp := range allEmployers {
		employers = append(employers, model.GetEmployerWithEmployeesResponse(emp))
	}
	return employers, err
}

func (service *employerService) GetEmployer(employerId uint) (employer model.EmployerResponse, err model.Error) {
	employerData, err := service.repo.GetEmployer(employerId)
	if err.StatusCode != 0 {
		return employer, err
	}
	return model.GetEmployerResponse(employerData), err
}

func (service *employerService) GetEmployerWithEmployees(employerId uint) (employer model.EmployerWithEmployeesResponse, err model.Error) {
	employerData, err := service.repo.GetEmployer(employerId)
	if err.StatusCode != 0 {
		return employer, err
	}
	return model.GetEmployerWithEmployeesResponse(employerData), err
}

func (service *employerService) UpdateEmployer(employerId uint, employer model.UpdateEmployerRequest) (err model.Error) {
	return service.repo.UpdateEmployer(employerId, employer.Name, employer.Location)
}

func (service *employerService) DeleteEmployer(employerId uint) (err model.Error) {
	return service.repo.DeleteEmployer(employerId)
}
