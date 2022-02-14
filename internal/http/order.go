package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/smhdhsn/food/internal/http/helper"
	"github.com/smhdhsn/food/internal/service"
)

// OrderReq is the order submittion's request schema.
type OrderReq struct {
	FoodID uint `json:"food_id"`
}

// OrderHandler contains services that can be used within order handler.
type OrderHandler struct {
	oService *service.OrderService
	res      *helper.RespBody
}

// NewOrderHandler creates a new order handler.
func NewOrderHandler(oService *service.OrderService) *OrderHandler {
	return &OrderHandler{
		oService: oService,
		res:      &helper.RespBody{},
	}
}

// SubmitOrder is responsible for submitting a food order.
func (h *OrderHandler) SubmitOrder(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(io.LimitReader(r.Body, 1024))
	defer r.Body.Close()

	var req OrderReq
	if err := decoder.Decode(&req); err != nil {
		h.res.
			SetError(err).
			SetMessage("invalid request payload or exceeded maximum payload size").
			Json(w, http.StatusBadRequest)

		return
	}

	data, err := h.oService.OrderFood(req.FoodID)
	if err != nil {
		h.res.
			SetData(data).
			SetError(err).
			SetMessage("failed to submit the order").
			Json(w, http.StatusUnprocessableEntity)

		return
	}

	h.res.
		SetData(data).
		SetMessage("your order is being prepared").
		Json(w, http.StatusAccepted)
}