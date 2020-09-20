package main

import (
	"log"
	"net"

	"github.com/MatthewZholud/TestTaskMicroservices/db"
	"github.com/MatthewZholud/TestTaskMicroservices/employee"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	employeeServicePort = ":3443"
)

func main() {
	postgres, err := db.NewPostgresDB()
	if err != nil {
		panic(err.Error())
	}
	s := grpc.NewServer()
	employee.RegisterEmployeeServer(s, &employee.Server{Database: postgres})
	lis, err := net.Listen("tcp", employeeServicePort)
	if err != nil {
		panic(err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
