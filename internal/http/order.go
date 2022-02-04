package http

import (
	"fmt"
	"net/http"
)

// OrderHandler contains services that can be used within order handler.
type OrderHandler struct{}

// NewOrderHandler creates a new order handler.
func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

// SubmitOrder is responsible for submitting a food order.
func (h *OrderHandler) SubmitOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You're here.")
}
