package main

import "net/http"
import "log"

type handler struct {
	// gateway

}

func NewHandler() *handler {
	return &handler{}

}
func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/coustomers/{CID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello Order is geting placed")

}
