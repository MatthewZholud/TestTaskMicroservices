package employee

func ToProtoEmployee(c Employee) EmployeeProto {
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
