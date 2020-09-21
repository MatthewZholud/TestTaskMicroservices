package routes

import (
	"github.com/MatthewZholud/TestTaskMicroservices/employee"
	"github.com/MatthewZholud/TestTaskMicroservices/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

//RegisterCustomerRoutes RegisterCustomerRoutes
func RegisterEmployeeRoutes(r *mux.Router, e employee.EmployeeClient) *mux.Router {

	r.HandleFunc("/employee", handlers.CreateEmployee(e)).Methods(http.MethodPost)
	//r.HandleFunc("/employee", handlers.PutEmployee(e)).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", handlers.GetEmployee(e)).Methods(http.MethodGet)
	//r.HandleFunc("/employee/{id}", handlers.PostEmployeeByID(e)).Methods(http.MethodPost)
	//r.HandleFunc("/employee/{id}", handlers.DeleteEmployeeByID(e)).Methods(http.MethodDelete)
	return r
}
