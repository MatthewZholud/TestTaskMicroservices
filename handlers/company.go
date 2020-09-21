package handlers

import (
	"context"
	"github.com/MatthewZholud/TestTaskMicroservices/company"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func CreateCompany(c company.CompanyClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var comp company.Company

		err := parseJsonToStruct(w, r, &comp)
		if err != nil {
			return
		}

		companyProtocol := &company.CompanyProto{
			Name:      comp.Name,
			Legalform: comp.Legalform,
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


func GetCompany(c company.CompanyClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		id64 := int64(id)

		incomingCompany, err := c.GetCompany(context.Background(), &company.Id{Id: id64})

		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		Respond(w, company.Company{
			ID:        incomingCompany.Id,
			Name:      incomingCompany.Name,
			Legalform: incomingCompany.Legalform,
		})
	}
}
