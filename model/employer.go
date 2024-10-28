package model

type Employer struct {
	Id        uint        `gorm:"primaryKey"`
	Name      string      `gorm:"size:100;not null"`
	Location  string      `gorm:"size:100;not null"`
	Employees []*Employee `gorm:"foreignKey:EmployerId"`
}

type EmployerResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type EmployerWithEmployeesResponse struct {
	Id        uint                          `json:"id"`
	Name      string                        `json:"name"`
	Location  string                        `json:"location"`
	Employees []EmployeeResponseForEmployer `json:"employees"`
}

type NewEmployerRequest struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

type UpdateEmployerRequest struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

func GetEmployerResponse(employer Employer) EmployerResponse {
	return EmployerResponse{
		Id:       employer.Id,
		Name:     employer.Name,
		Location: employer.Location,
	}
}

func GetEmployerWithEmployeesResponse(employer Employer) EmployerWithEmployeesResponse {
	employees := []EmployeeResponseForEmployer{}
	for _, employee := range employer.Employees {
		employees = append(employees, GetEmployeeResponseForEmployer(*employee))
	}
	return EmployerWithEmployeesResponse{
		Id:        employer.Id,
		Name:      employer.Name,
		Location:  employer.Location,
		Employees: employees,
	}
}
