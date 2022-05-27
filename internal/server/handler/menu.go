package handler

import (
	"net/http"

	"github.com/smhdhsn/restaurant-menu/internal/server/helper"

	serviceContract "github.com/smhdhsn/restaurant-menu/internal/service/contract"
)

// MenuResp is the response schema of the menu API.
type MenuResp struct {
	ID    uint32 `json:"id"`
	Title string `json:"title"`
}

// MenuHandler contains services that can be used within menu handler.
type MenuHandler struct {
	mServ serviceContract.MenuService
	res   *helper.RespBody
}

// NewMenuHandler creates a new menu handler.
func NewMenuHandler(ms serviceContract.MenuService) *MenuHandler {
	return &MenuHandler{
		mServ: ms,
		res:   &helper.RespBody{},
	}
}

// GetMenu is responsible for getting menu with available food.
func (h *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	iList, err := h.mServ.List()
	if err != nil {
		h.res.
			SetError(err).
			SetMessage("failed to get menu").
			Json(w, http.StatusBadRequest)

		return
	}

	transform := make([]MenuResp, 0)
	for _, i := range iList {
		transform = append(transform, MenuResp{
			ID:    i.ID,
			Title: i.Title,
		})
	}

	h.res.
		SetData(transform).
		Json(w, http.StatusOK)
}
