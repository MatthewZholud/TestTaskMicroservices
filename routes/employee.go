package routes

import (
	"github.com/MatthewZholud/TestTaskMicroservices/employee"
	"github.com/MatthewZholud/TestTaskMicroservices/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterEmployeeRoutes(r *mux.Router, e employee.EmployeeClient) *mux.Router {

	r.HandleFunc("/employee", handlers.CreateEmployee(e)).Methods(http.MethodPost)
	r.HandleFunc("/employee", handlers.UpdateEmployee(e)).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", handlers.GetEmployee(e)).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", handlers.FormUpdateEmployee(e)).Methods(http.MethodPost)
	r.HandleFunc("/employee/{id}", handlers.DeleteEmployee(e)).Methods(http.MethodDelete)
	return r
}
