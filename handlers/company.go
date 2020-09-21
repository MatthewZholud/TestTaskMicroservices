package handlers

import (
	"context"
	"github.com/MatthewZholud/TestTaskMicroservices/company"
	"github.com/MatthewZholud/TestTaskMicroservices/employee"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateCompany(c company.CompanyClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var comp company.Company

		err := parseJsonToStruct(w, r, &comp)
		if err != nil {
			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
			return
		}

		companyProtocol := &company.CompanyProto{
			Name:      comp.Name,
			Legalform: comp.Legalform,
		}
		newCompanyIDProto, err := c.CreateCompany(r.Context(), companyProtocol)
		if err != nil {
			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
			return
		}
		Respond(w, newCompanyIDProto)
	}
}

func GetCompany(c company.CompanyClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		id64 := int64(id)

		incomingCompany, err := c.GetCompany(context.Background(), &company.Id{Id: id64})
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Employee not found")
			return
		}

		Respond(w, company.Company{
			ID:        incomingCompany.Id,
			Name:      incomingCompany.Name,
			Legalform: incomingCompany.Legalform,
		})
	}
}

func DeleteCompany(c company.CompanyClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		id64 := int64(id)

		_, err = c.DeleteCompany(context.Background(), &company.Id{Id: id64})

		if err != nil {
			respondWithError(w, http.StatusNotFound, "Company not found")
			return
		}

	}
}

func UpdateCompany(c company.CompanyClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var comp company.Company

		err := parseJsonToStruct(w, r, &comp)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		companyProtocol := &company.CompanyProto{
			Id:        comp.ID,
			Name:      comp.Name,
			Legalform: comp.Legalform,
		}
		_, err = c.UpdateCompany(r.Context(), companyProtocol)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Employee not found")
			return
		}
	}
}

func FormUpdateCompany(c company.CompanyClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var comp company.Company

		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
		if err != nil {
			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
			return
		}

		id64 := int64(id)

		err = parseJsonToStruct(w, r, &comp)
		if err != nil {
			return
		}

		companyProtocol := &company.CompanyProto{
			Id:        id64,
			Name:      comp.Name,
			Legalform: comp.Legalform,
		}
		_, err = c.FormUpdateCompany(r.Context(), companyProtocol)
		if err != nil {
			respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
			return
		}
	}
}

func GetEmployeesByCompany(c company.CompanyClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if IsNumericAndPositive(mux.Vars(r)["companyId"]) != true {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(mux.Vars(r)["companyId"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		id64 := int64(id)

		employeesProto, err := c.GetEmployeesByCompany(context.Background(), &company.Id{Id: id64})
		if err != nil {
			respondWithError(w, http.StatusNotFound, "ICompany not found")
			return
		}
		var employees []employee.Employee
		for _, empl := range employeesProto.Employees {
			employees = append(employees, employee.Employee{
				ID:         empl.Id,
				Name:       empl.Name,
				SecondName: empl.SecondName,
				Surname:    empl.Surname,
				PhotoUrl:   empl.PhotoUrl,
				HireDate:   empl.HireDate,
				Position:   empl.Position,
				CompanyID:  empl.CompanyId,
			})
		}

		Respond(w, employees)
	}
}
