package service

import (
	"go-postgres-test-1/model"
	"go-postgres-test-1/repository"
)

type EmployeeService interface {
	CreateEmployee(employee model.NewEmployeeRequest) (err model.Error)
	GetAllEmployees() (employees []model.EmployeeResponse, err model.Error)
	GetAllEmployeesWithEmployer() (employees []model.EmployeeWithEmployerResponse, err model.Error)
	GetEmployee(employeeId uint) (employee model.EmployeeResponse, err model.Error)
	GetEmployeeWithEmployer(employeeId uint) (employee model.EmployeeWithEmployerResponse, err model.Error)
	UpdateEmployee(employeeId uint, employee model.UpdateEmployeeRequest) (err model.Error)
	DeleteEmployee(employeeId uint) (err model.Error)
}

type employeeService struct {
	employeeRepo repository.EmployeeRepository
	employerRepo repository.EmployerRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository, employerRepo repository.EmployerRepository) EmployeeService {
	return &employeeService{employeeRepo: employeeRepo, employerRepo: employerRepo}
}

func (service *employeeService) CreateEmployee(employee model.NewEmployeeRequest) (err model.Error) {
	employer, err := service.employerRepo.GetEmployer(employee.EmployerId)
	if err.StatusCode != 0 {
		return err
	}
	return service.employeeRepo.CreateEmployee(employee.Name, employee.Age, employee.Salary, &employer.Id)
}

func (service *employeeService) GetAllEmployees() (employees []model.EmployeeResponse, err model.Error) {
	allEmployees, err := service.employeeRepo.GetAllEmployees()
	if err.StatusCode != 0 {
		return employees, err
	}
	for _, emp := range allEmployees {
		employees = append(employees, model.GetEmployeeResponse(emp))
	}
	return employees, err
}

func (service *employeeService) GetAllEmployeesWithEmployer() (employees []model.EmployeeWithEmployerResponse, err model.Error) {
	allEmployees, err := service.employeeRepo.GetAllEmployees()
	if err.StatusCode != 0 {
		return employees, err
	}
	for _, emp := range allEmployees {
		employees = append(employees, model.GetEmployeeWithEmployerResponse(emp))
	}
	return employees, err
}

func (service *employeeService) GetEmployee(employeeId uint) (employee model.EmployeeResponse, err model.Error) {
	employeeData, err := service.employeeRepo.GetEmployee(employeeId)
	if err.StatusCode != 0 {
		return employee, err
	}
	return model.GetEmployeeResponse(employeeData), err
}

func (service *employeeService) GetEmployeeWithEmployer(employeeId uint) (employee model.EmployeeWithEmployerResponse, err model.Error) {
	employeeData, err := service.employeeRepo.GetEmployee(employeeId)
	if err.StatusCode != 0 {
		return employee, err
	}
	return model.GetEmployeeWithEmployerResponse(employeeData), err
}

func (service *employeeService) UpdateEmployee(employeeId uint, updateEmployee model.UpdateEmployeeRequest) (err model.Error) {
	employee, err := service.employeeRepo.GetEmployee(employeeId)
	if err.StatusCode != 0 {
		return err
	}

	if updateEmployee.Name != nil {
		employee.Name = *updateEmployee.Name
	}
	if updateEmployee.Age != nil {
		employee.Age = *updateEmployee.Age
	}
	if updateEmployee.Salary != nil {
		employee.Salary = *updateEmployee.Salary
	}
	if updateEmployee.EmployerId != nil {
		employer, err := service.employerRepo.GetEmployer(*updateEmployee.EmployerId)
		if err.StatusCode != 0 {
			return err
		}
		employee.EmployerId = &employer.Id
	}

	return service.employeeRepo.UpdateEmployee(employeeId, employee.Name, employee.Age, employee.Salary, employee.EmployerId, updateEmployee.UnsetEmployer)
}

func (service *employeeService) DeleteEmployee(employeeId uint) (err model.Error) {
	return service.employeeRepo.DeleteEmployee(employeeId)
}
