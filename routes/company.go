package routes

import (
	"github.com/MatthewZholud/TestTaskMicroservices/company"
	"github.com/MatthewZholud/TestTaskMicroservices/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterCompanyRoutes(r *mux.Router, c company.CompanyClient) *mux.Router {

	r.HandleFunc("/company/", handlers.CreateCompany(c)).Methods(http.MethodPost)
	r.HandleFunc("/company/", handlers.UpdateCompany(c)).Methods(http.MethodPut)
	r.HandleFunc("/company/{companyId}", handlers.GetCompany(c)).Methods(http.MethodGet)
	r.HandleFunc("/company/{companyId}", handlers.FormUpdateCompany(c)).Methods(http.MethodPost)
	r.HandleFunc("/company/{companyId}", handlers.DeleteCompany(c)).Methods(http.MethodDelete)
	r.HandleFunc("/company/{companyId}/employees", handlers.GetEmployeesByCompany(c)).Methods(http.MethodGet)

	return r
}
