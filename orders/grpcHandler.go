package main

import (
	"context"
	"log"

	pb "github.com/feliux/commons/api"
)

// type Server struct {
// 	pb.UnimplementedOrderServiceServer
// }

// func NewGRPCHandler(grpcServer *grpc.Server) *Server {
// 	handler := &Server{}
// 	pb.RegisterOrderServiceServer(grpcServer, handler)
// }

type Server struct {
	pb.OrderServiceServer
}

func (s *Server) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// get menu
	// pub new order
	// register orderDB
	log.Printf("New order received from customer %s, %v\n", in.CustomerID, in.Items)
	return &pb.CreateOrderResponse{
		ID:         "42",
		CustomerID: in.CustomerID,
		Status:     "status",
		// Items: pb.Item{
		// 	ID:       "1",
		// 	Name:     "burguer",
		// 	Quantity: 2,
		// 	PriceID:  "9",
		// },
	}, nil
}
