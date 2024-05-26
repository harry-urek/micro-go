package main

import (
	"log"
	"net/http"

	"github.com/harry-urek/common"
	pb "github.com/harry-urek/common/api"
)

type handler struct {
	// gateway
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client: client}

}
func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/coustomers/{CID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello Order is geting placed")
	coustomerID := r.PathValue("coustomerID")

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: coustomerID,
		Items:      items,
	})

}
