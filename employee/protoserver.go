package employee

import (
	"context"
	"github.com/MatthewZholud/TestTaskMicroservices/db"
)

//Server Server
type Server struct {
	Database *db.Postgres
}

//GetCustomer GetCustomer
func (s *Server) GetEmployee(ctx context.Context, in *Id) (*EmployeeProto, error) {
	var employee Employee
	rows, err := s.Database.Db.Query("SELECT * from employees WHERE employee_id = $1", in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.SecondName, &employee.Surname,
			&employee.PhotoUrl, &employee.HireDate, &employee.Position, &employee.CompanyID);
			err != nil {
			return nil, err
		}
	}

	employeeProto := ToProtoEmployee(employee)
	return &employeeProto, nil
}

//CreateCustomer CreateCustomer
func (s *Server) CreateEmployee(ctx context.Context, in *EmployeeProto) (*Id, error) {
	var empId int64
	err := s.Database.Db.QueryRow("INSERT INTO employees(name, secondName, surname, photoUrl, hireDate, position, company_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING employee_id",
		in.Name, in.SecondName, in.Surname, in.PhotoUrl, in.HireDate, in.Position, in.CompanyId).Scan(empId)
	if err != nil {
		return nil, err
	}
	return &Id{Id: empId}, nil
}

//GetAllCustomers GetAllCustomers
//func (s *Server) GetAllCustomers(ctx context.Context, in *NothingFancy) (*CustomersProto, error) {
//
//	collection := s.Database.GetCollection("customers")
//	var customers []Customer
//	collection.Find(bson.M{}).All(&customers)
//
//	customersProto := ToMultipleProtoCustomer(customers)
//	return &customersProto, nil
//}
