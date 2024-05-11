package main

import (
	"context"
	"log"
	"net"

	"github.com/feliux/commons"
	pb "github.com/feliux/commons/api"
	"google.golang.org/grpc"
)

var (
	grpcOrdersAddr string = commons.EnvString("GRPC_ORDERS_ADDR", ":8081")
	tls            bool   = false
)

func main() {
	// Transport layer
	l, err := net.Listen("tcp", grpcOrdersAddr)
	if err != nil {
		log.Fatalf("Failed to listen on addr: %v\n", err)
	}
	defer l.Close()
	opts := commons.NewGRPCServerOps(tls)
	grpcServer := grpc.NewServer(opts...)
	defer grpcServer.Stop()
	// Data layer
	store := NewStore()
	// Business layer
	svc := NewService(store)
	svc.CreateOrder(context.Background())
	// NewGRPCHandler(grpcServer)
	pb.RegisterOrderServiceServer(grpcServer, &Server{})
	log.Printf("Listening at %s\n", grpcOrdersAddr)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("Failed to start gRPC server: %v\n", err)
	}
}
