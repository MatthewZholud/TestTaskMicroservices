package company

import (
	"context"

	"github.com/MatthewZholud/TestTaskMicroservices/db"
)

type Server struct {
	Database *db.Postgres
}

//GetOrder GetOrder
func (s *Server) GetCompany(ctx context.Context, in *Id) (*CompanyProto, error) {
	var company Company

	rows, err := s.Database.Db.Query("SELECT * from company WHERE company_id = $1", in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&company.ID, &company.Name, &company.Legalform);
			err != nil {
			return nil, err
		}
	}
	companyProto := ToProtoCompany(company)
	return &companyProto, nil
}

//CreateOrder CreateOrder
func (s *Server) CreateCompany(ctx context.Context, in *CompanyProto) (*Id, error) {
	var compId int64
	err := s.Database.Db.QueryRow("INSERT INTO company(name, legal_form) VALUES ($1, $2) RETURNING company_id", in.Name, in.Legalform).Scan(compId)
	if err != nil {
		return nil, err
	}
	return &Id{Id: compId}, nil
}

//GetAllOrders GetAllOrders
//func (s *Server) GetAllOrders(ctx context.Context, in *Noop) (*OrdersProto, error) {
//
//	var orders []Order
//
//	query := []bson.M{{"$lookup": bson.M{ // lookup the documents table here
//		"from":         "customers",
//		"localField":   "customerid",
//		"foreignField": "_id",
//		"as":           "employee",
//	}}}
//
//	if err := s.Database.GetCollection("orders").Pipe(query).All(&orders); err != nil {
//		return nil, err
//	}
//	if len(orders) == 0 {
//		return nil, errors.New("No Orders Yet")
//	}
//
//	ordersProto := ToMultipleOrderProto(orders)
//	return &ordersProto, nil
//
//}
