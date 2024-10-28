package model

type Employee struct {
	Id         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"size:100;not null"`
	Age        int     `gorm:"check:age > 0"`
	Salary     float64 `gorm:"type:decimal(10,2);not null"`
	EmployerId *uint
	Employer   *Employer `gorm:"foreignKey:EmployerId"`
}

type EmployeeResponse struct {
	Id         uint    `json:"id"`
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Salary     float64 `json:"salary"`
	EmployerId *uint   `json:"employerId"`
}

type EmployeeWithEmployerResponse struct {
	Id       uint              `json:"id"`
	Name     string            `json:"name"`
	Age      int               `json:"age"`
	Salary   float64           `json:"salary"`
	Employer *EmployerResponse `json:"employer"`
}

type EmployeeResponseForEmployer struct {
	Id     uint    `json:"id"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}

type NewEmployeeRequest struct {
	Name       string  `json:"name" binding:"required"`
	Age        int     `json:"age" binding:"required"`
	Salary     float64 `json:"salary" binding:"required"`
	EmployerId uint    `json:"employerId"`
}

type UpdateEmployeeRequest struct {
	Name          *string  `json:"name"`
	Age           *int     `json:"age"`
	Salary        *float64 `json:"salary"`
	EmployerId    *uint    `json:"employerId"`
	UnsetEmployer *bool    `json:"unsetEmployer"`
}

func GetEmployeeResponse(employee Employee) EmployeeResponse {
	data := EmployeeResponse{
		Id:     employee.Id,
		Name:   employee.Name,
		Age:    employee.Age,
		Salary: employee.Salary,
	}
	if employee.EmployerId != nil {
		data.EmployerId = employee.EmployerId
	}
	return data
}

func GetEmployeeResponseForEmployer(employee Employee) EmployeeResponseForEmployer {
	return EmployeeResponseForEmployer{
		Id:     employee.Id,
		Name:   employee.Name,
		Age:    employee.Age,
		Salary: employee.Salary,
	}
}

func GetEmployeeWithEmployerResponse(employee Employee) EmployeeWithEmployerResponse {
	data := EmployeeWithEmployerResponse{
		Id:     employee.Id,
		Name:   employee.Name,
		Age:    employee.Age,
		Salary: employee.Salary,
	}
	if employee.EmployerId != nil {
		data.Employer = &EmployerResponse{
			Id:       employee.Employer.Id,
			Name:     employee.Employer.Name,
			Location: employee.Employer.Location,
		}
	}
	return data
}
