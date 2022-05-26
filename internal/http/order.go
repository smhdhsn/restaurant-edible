package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/smhdhsn/restaurant-menu/internal/http/helper"

	oServContract "github.com/smhdhsn/restaurant-menu/internal/service/contract/order"
)

// OrderReq is the order submittion's request schema.
type OrderReq struct {
	FoodID uint32 `json:"food_id"`
}

// OrderHandler contains services that can be used within order handler.
type OrderHandler struct {
	oServ oServContract.OrderService
	res   *helper.RespBody
}

// NewOrderHandler creates a new order handler.
func NewOrderHandler(o oServContract.OrderService) *OrderHandler {
	return &OrderHandler{
		oServ: o,
		res:   &helper.RespBody{},
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

	data, err := h.oServ.OrderFood(req.FoodID)
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
