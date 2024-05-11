package main

import (
	"log"
	"net/http"

	"github.com/feliux/commons"
	pb "github.com/feliux/commons/api"
)

type handler struct {
	grpcClient pb.OrderServiceClient
}

// NewHandler is the constructor for handling requests.
func NewHandler(c pb.OrderServiceClient) *handler {
	return &handler{grpcClient: c}
}

// registerRoutes add the routes to a HTTP server.
func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

// HandleCreateOrder takes the incoming HTTP request and sends it to a gRPC server
// for reating a new order.
func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var items []*pb.ItemsWithQuantity
	if err := commons.ReadJSON(r, &items); err != nil {
		commons.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	in := &pb.CreateOrderRequest{
		CustomerID: r.PathValue("customerID"),
		Items:      items,
	}
	log.Printf("Sending new order from customer %s, %v", in.CustomerID, in.Items)
	h.grpcClient.CreateOrder(r.Context(), in)
}
