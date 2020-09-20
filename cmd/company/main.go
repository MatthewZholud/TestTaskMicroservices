package main

import (
	"log"
	"net"

	"github.com/MatthewZholud/TestTaskMicroservices/company"
	"github.com/MatthewZholud/TestTaskMicroservices/db"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	companyServiceAddr = ":4443"
)

func main() {
	postgres, err := db.NewPostgresDB()
	if err != nil {
		panic(err.Error())
	}

	s := grpc.NewServer()
	company.RegisterCompanyServer(s, &company.Server{Database: postgres})
	lis, err := net.Listen("tcp", companyServiceAddr)
	if err != nil {
		panic(err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
