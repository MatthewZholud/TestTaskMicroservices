package routes

import (
	"github.com/MatthewZholud/TestTaskMicroservices/company"
	"github.com/MatthewZholud/TestTaskMicroservices/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterCompanyRoutes(r *mux.Router, c company.CompanyClient) *mux.Router {

	r.HandleFunc("/company/", handlers.CreateCompany(c)).Methods(http.MethodPost)
	//r.HandleFunc("/company/", handlers.PutCompany(c)).Methods(http.MethodPut)
	r.HandleFunc("/company/{companyId}", handlers.GetCompany(c)).Methods(http.MethodGet)
	//r.HandleFunc("/company/{companyId}", handlers.PostCompanyByID(c)).Methods(http.MethodPost)
	//r.HandleFunc("/company/{companyId}", handlers.DeleteCompanyByID(c)).Methods(http.MethodDelete)
	//r.HandleFunc("/company/{companyId}/employees", handlers.GetEmployeeByCompanyID(c)).Methods(http.MethodGet)

	return r
}
