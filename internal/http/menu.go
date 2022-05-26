package http

import (
	"net/http"

	"github.com/smhdhsn/restaurant-menu/internal/http/helper"

	mServ "github.com/smhdhsn/restaurant-menu/internal/service/menu"
)

// MenuResp is the response schema of the menu API.
type MenuResp struct {
	ID    uint32 `json:"id"`
	Title string `json:"title"`
}

// MenuHandler contains services that can be used within menu handler.
type MenuHandler struct {
	mServ *mServ.MenuService
	res   *helper.RespBody
}

// NewMenuHandler creates a new menu handler.
func NewMenuHandler(m *mServ.MenuService) *MenuHandler {
	return &MenuHandler{
		mServ: m,
		res:   &helper.RespBody{},
	}
}

// GetMenu is responsible for getting menu with available food.
func (h *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	foods, err := h.mServ.GetFoods()
	if err != nil {
		h.res.
			SetError(err).
			SetMessage("failed to get menu").
			Json(w, http.StatusBadRequest)

		return
	}

	transform := make([]MenuResp, 0)
	for _, f := range foods {
		transform = append(transform, MenuResp{
			ID:    f.ID,
			Title: f.Title,
		})
	}

	h.res.
		SetData(transform).
		Json(w, http.StatusOK)
}
