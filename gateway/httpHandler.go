package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/feliux/commons"
	pb "github.com/feliux/commons/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if err := validateItems(items); err != nil {
		commons.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	in := &pb.CreateOrderRequest{
		CustomerID: r.PathValue("customerID"),
		Items:      items,
	}
	log.Printf("Sending new order from customer %s, %v", in.CustomerID, in.Items)
	resp, err := h.grpcClient.CreateOrder(r.Context(), in)
	// handle rpc errors
	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			commons.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
		commons.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	commons.WriteJSON(w, http.StatusOK, resp)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return commons.ErrNoItems
	}
	for _, i := range items {
		if i.ID == "" {
			return errors.New("Item ID is required")
		}
		if i.Quantity < 0 {
			return errors.New("Item must have a valid quantity")
		}
	}
	return nil
}
