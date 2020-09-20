package handlers

import (
	"context"
	"github.com/MatthewZholud/TestTaskMicroservices/company"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllOrders  GetAllOrders
//func GetAllOrders(o order.OrderClient) http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		ordersProto, err := o.GetAllOrders(r.Context(), &order.Noop{})
//		if err != nil {
//			boom.BadRequest(w, err)
//			return
//		}
//		var os []order.Order
//		for _, orderIteree := range ordersProto.Orders {
//			os = append(os, order.Order{
//				ID:         bson.ObjectIdHex(*orderIteree.Id),
//				Name:       *orderIteree.Name,
//				Address:    *orderIteree.Address,
//				CustomerID: convert.StringPtrToMongoID(orderIteree.CustomerId),
//				Customer: func(cpros *order.CustomersProto) []customer.Customer {
//					customers := []customer.Customer{}
//					for _, c := range cpros.Customers {
//						customers = append(customers, customer.Customer{
//							ID:      convert.StringPtrToMongoID(c.Id),
//							Name:    *c.Name,
//							Address: *c.Address,
//							Email:   *c.Email,
//						})
//					}
//					return customers
//				}(orderIteree.Customers),
//			})
//		}
//		utils.Respond(w, os)
//	}
//}
//
////GetOrder GetOrder
//func GetOrder(o order.OrderClient) http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//		oID := mux.Vars(r)["id"]
//		o, err := o.GetOrder(context.Background(), &order.Id{Id: &oID})
//		if err != nil {
//			boom.BadRequest(w, err)
//			return
//		}
//
//		utils.Respond(w, order.Order{
//			ID:         bson.ObjectIdHex(*o.Id),
//			Name:       *o.Name,
//			Address:    *o.Address,
//			CustomerID: convert.StringPtrToMongoID(o.CustomerId),
//			Customer: func(cpros *order.CustomersProto) []customer.Customer {
//				customers := []customer.Customer{}
//				for _, c := range cpros.Customers {
//					customers = append(customers, customer.Customer{
//						ID:      convert.StringPtrToMongoID(c.Id),
//						Name:    *c.Name,
//						Address: *c.Address,
//						Email:   *c.Email,
//					})
//				}
//				return customers
//			}(o.Customers),
//		})
//
//	}
//}
//
////CreateOrder CreateOrder
//func CreateOrder(o order.OrderClient) http.HandlerFunc {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		var ord order.Order
//		err := utils.DecodeRequest(r, &ord)
//		if err != nil {
//			boom.BadRequest(w, err)
//			return
//		}
//
//		orderID := string(bson.NewObjectId())
//		orderProtocol := &order.OrderProto{
//			Id:         &orderID,
//			Name:       &ord.Name,
//			Address:    &ord.Address,
//			CustomerId: convert.MongoIDToStringPtr(ord.CustomerID),
//		}
//		newOrderIDProtocol, err := o.CreateOrder(r.Context(), orderProtocol)
//		if err != nil {
//			boom.BadRequest(w, err)
//			return
//		}
//		utils.Respond(w, struct {
//			CustomerID string
//		}{
//			*newOrderIDProtocol.Id,
//		})
//	}
//}

func PostCompany(c company.CompanyClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var comp company.Company

		err := parseJsonToStruct(w, r, &comp)
		if err != nil {
			return
		}

		companyProtocol := &company.CompanyProto{
			Name:      &comp.Name,
			Legalform: &comp.Legalform,
		}
		newCompanyIDProto, err := c.CreateCompany(r.Context(), companyProtocol)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}
		Respond(w, newCompanyIDProto)
		return
	}
}

//func PutCompany(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	company := Entities.Company{}
//
//	err := parseJsonToStruct(w, r, &company)
//	if err != nil {
//		return
//	}
//	err = DbService.Conn.PutCompany(&company)
//	if err != nil {
//		log.Println(err)
//		respondWithError(w, http.StatusNotFound, "Company not found")
//		return
//	}
//	json.NewEncoder(w).Encode(company)
//}

func GetCompany(c company.CompanyClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}
		id64 := int64(id)

		incomingCompany, err := c.GetCompany(context.Background(), &company.Id{Id: &id64})
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		Respond(w, company.Company{
			ID:        *incomingCompany.Id,
			Name:      *incomingCompany.Name,
			Legalform: *incomingCompany.Legalform,
		})
	}
}

//func PostCompanyByID(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	company := Entities.Company{}
//	id, err := strconv.Atoi(mux.Vars(r)["companyId"])
//	if err != nil || id < 1 {
//		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//		return
//	}
//	company.ID = int64(id)
//
//	company.Name = r.Form.Get("name")
//	company.LegalForm = r.Form.Get("status")
//
//	err = DbService.Conn.PostCompanyById(&company)
//	if err != nil {
//		log.Println(err)
//		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//		return
//	}
//	json.NewEncoder(w).Encode(company)
//}
//
//func DeleteCompanyByID(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	company := Entities.Company{}
//	if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	id, err := strconv.Atoi(mux.Vars(r)["companyId"])
//	if err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
//		return
//	}
//	company.ID = int64(id)
//	err = DbService.Conn.DeleteCompany(company.ID)
//	if err != nil {
//		log.Println(err)
//		respondWithError(w, http.StatusNotFound, "Company not found")
//		return
//	}
//}
//
//func GetEmployeeByCompanyID(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	company := Entities.Company{}
//	id, err := strconv.Atoi(mux.Vars(r)["companyId"])
//	if err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
//		return
//	}
//	company.ID = int64(id)
//	employees, err := DbService.Conn.GetEmployeesByCompanyId(company.ID)
//	if err != nil {
//		log.Println(err)
//		respondWithError(w, http.StatusNotFound, "Company not found")
//		return
//	}
//	json.NewEncoder(w).Encode(employees)
//	w.WriteHeader(http.StatusOK)
//}
