package company

import "github.com/MatthewZholud/TestTaskMicroservices/employee"

func ToMultipleProtoEmployee(c []employee.Employee) Employees {

	employeesProto := Employees{}
	for _, m := range c {
		x := ToProtoEmployee(m)
		employeesProto.Employees = append(employeesProto.Employees, &x)
	}
	return employeesProto
}

func ToProtoEmployee(c employee.Employee) EmployeeProto {
	employeeProto := EmployeeProto{
		Id:         c.ID,
		Name:       c.Name,
		SecondName: c.SecondName,
		Surname:    c.Surname,
		PhotoUrl:   c.PhotoUrl,
		HireDate:   c.HireDate,
		Position:   c.Position,
		CompanyId:  c.CompanyID,
	}
	return employeeProto
}


func ToProtoCompany(c Company) CompanyProto {
	companyProto := CompanyProto{
		Id:        c.ID,
		Name:      c.Name,
		Legalform: c.Legalform,
	}
	return companyProto
}

