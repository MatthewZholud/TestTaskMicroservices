package handlers

import (
	"context"
	"github.com/MatthewZholud/TestTaskMicroservices/employee"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//CreateCustomer CreateCustomer
func PostEmployee(e employee.EmployeeClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var empl employee.Employee

		err := parseJsonToStruct(w, r, &empl)
		if err != nil {
			return
		}

		employeeProtocol := &employee.EmployeeProto{
			Name:       &empl.Name,
			SecondName: &empl.SecondName,
			Surname:    &empl.Surname,
			PhotoUrl:   &empl.PhotoUrl,
			HireDate:   &empl.HireDate,
			Position:   &empl.Position,
			CompanyId:  &empl.CompanyID,
		}
		newEmployeeIDProto, err := e.CreateEmployee(r.Context(), employeeProtocol)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}
		Respond(w, newEmployeeIDProto)
		return
	}
}

//func PutEmployee(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	employee := Entities.Employee{}
//
//	err := parseJsonToStruct(w, r, &employee)
//	if err != nil {
//		return
//	}
//	err = DbService.Conn.PutEmployee(&employee)
//	if err != nil {
//		log.Println(err)
//		respondWithError(w, http.StatusNotFound, "Employee not found")
//		return
//	}
//	json.NewEncoder(w).Encode(employee)
//}

func GetEmployee(e employee.EmployeeClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}
		id64 := int64(id)

		incomingEmployee, err := e.GetEmployee(context.Background(), &employee.Id{Id: &id64})
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
			return
		}

		Respond(w, employee.Employee{
			ID:         *incomingEmployee.Id,
			Name:       *incomingEmployee.Name,
			SecondName: *incomingEmployee.SecondName,
			Surname:    *incomingEmployee.Surname,
			PhotoUrl:   *incomingEmployee.PhotoUrl,
			HireDate:   *incomingEmployee.HireDate,
			Position:   *incomingEmployee.Position,
			CompanyID:  *incomingEmployee.CompanyId,
		})
	}
}

//func PostEmployeeByID(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	employee := Entities.Employee{}
//	id, err := strconv.Atoi(mux.Vars(r)["id"])
//	if err != nil {
//		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//		return
//	}
//	employee.ID = int64(id)
//
//	id, err = strconv.Atoi(r.Form.Get("CompanyId"))
//	if err != nil {
//		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//		return
//	}
//	employee.CompanyID = int64(id)
//	employee.Name = r.Form.Get("name")
//	employee.SecondName = r.Form.Get("secondName")
//	employee.Surname = r.Form.Get("surname")
//	employee.HireDate = r.Form.Get("hireDate")
//	employee.Position = r.Form.Get("position")
//
//	err = DbService.Conn.PostEmployeeById(&employee)
//	if err != nil {
//		log.Println(err)
//		respondWithError(w, http.StatusMethodNotAllowed, "Invalid input")
//		return
//	}
//	json.NewEncoder(w).Encode(employee)
//}
//
//func DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	employee := Entities.Employee{}
//	id, err := strconv.Atoi(mux.Vars(r)["id"])
//	if err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid ID supplied")
//		return
//	}
//	employee.ID = int64(id)
//	err = DbService.Conn.DeleteEmployee(employee.ID)
//	if err != nil {
//		log.Println(err)
//		respondWithError(w, http.StatusNotFound, "Employee not found")
//		return
//	}
//}
