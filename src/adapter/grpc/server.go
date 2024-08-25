package grpc

import (
	"fmt"
	"grpcbank/generated_proto/bank"
	"grpcbank/src/port"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	bankService port.BankServicePort
	grpcPort    int
	server      *grpc.Server
	bank.BankServiceServer
}

func NewGrpcAdapter(bankService port.BankServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		bankService: bankService,
		grpcPort:    grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port %d : %v\n", a.grpcPort, err)
	}

	log.Printf("Server listening on port %d\n", a.grpcPort)

	grpcServer := grpc.NewServer()

	a.server = grpcServer

	bank.RegisterBankServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC on port %d : %v\n", a.grpcPort, err)
	}
}

func (a *GrpcAdapter) Stop() {
	a.server.Stop()
}
