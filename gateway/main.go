package main

import (
	"log"
	"net/http"

	"github.com/feliux/commons"
	pb "github.com/feliux/commons/api"
	"google.golang.org/grpc"
)

var (
	httpAddr       string = commons.EnvString("HTTP_ADDR", ":8080")
	grpcOrdersAddr string = commons.EnvString("GRPC_ORDERS_ADDR", ":8081")
	tls            bool   = false
)

// Only transport layer for this service
func main() {
	// GRPC client
	opts := commons.NewGRPCClientOps(tls)
	conn, err := grpc.Dial(grpcOrdersAddr, opts...)
	defer conn.Close()
	if err != nil {
		log.Fatalf("gRPC connection failed: %v", err)
	}
	grpcClient := pb.NewOrderServiceClient(conn)
	//HTTP server
	mux := http.NewServeMux()
	handler := NewHandler(grpcClient)
	handler.registerRoutes(mux)
	log.Printf("Starting HTTP server on %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start HTTP server", err)
	}
}
