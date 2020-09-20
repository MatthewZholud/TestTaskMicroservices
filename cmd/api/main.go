package main

import (
	"log"
	"net/http"

	"github.com/MatthewZholud/TestTaskMicroservices/routes"
	"github.com/gorilla/mux"

	"github.com/MatthewZholud/TestTaskMicroservices/company"
	"github.com/MatthewZholud/TestTaskMicroservices/employee"
	"google.golang.org/grpc"
)

const (
	apiGatewayPort           = ":3000"
	employeeMicroServiceAddr = "localhost:4443"
	companyMicroServiceAddr  = "localhost:3443"
)

func main() {

	emplConn, err := grpc.Dial(employeeMicroServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to employee service: %v", err)
	}
	defer emplConn.Close()

	compConn, err := grpc.Dial(companyMicroServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to company service : %v", err)
	}
	defer compConn.Close()

	e := employee.NewEmployeeClient(emplConn)
	c := company.NewCompanyClient(compConn)

	r := mux.NewRouter()
	routes.RegisterEmployeeRoutes(r, e)
	routes.RegisterCompanyRoutes(r, c)
	log.Fatal(http.ListenAndServe(apiGatewayPort, r))

}
