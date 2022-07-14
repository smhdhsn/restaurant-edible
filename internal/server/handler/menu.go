package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/smhdhsn/restaurant-edible/internal/logger"
	menuProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/menu"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// MenuHandler contains services that can be used within menu handler.
type MenuHandler struct {
	menuServ serviceContract.MenuService
}

// NewMenuHandler creates a new menu handler.
func NewMenuHandler(ms serviceContract.MenuService) menuProto.EdibleMenuServiceServer {
	return &MenuHandler{
		menuServ: ms,
	}
}

// List is responsible for getting menu.
func (s *MenuHandler) List(ctx context.Context, req *menuProto.MenuListRequest) (*menuProto.MenuListResponse, error) {
	fListDTO, err := s.menuServ.List()
	if err != nil {
		log.Error(err)
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := multipleFoodDTOToMenuResp(fListDTO)

	return resp, nil
}

// multipleFoodDTOToMenuResp is responsible for transforming a list of food dto to a list of food response struct.
func multipleFoodDTOToMenuResp(fListDTO []*dto.Food) *menuProto.MenuListResponse {
	fListResp := make([]*menuProto.Food, len(fListDTO))

	for i, fDTO := range fListDTO {
		cListResp := make([]*menuProto.Ingredient, len(fDTO.Components))

		for j, cDTO := range fDTO.Components {
			cListResp[j] = &menuProto.Ingredient{
				Title: cDTO.Title,
			}
		}

		fListResp[i] = &menuProto.Food{
			Id:          fDTO.ID,
			Title:       fDTO.Title,
			Ingredients: cListResp,
		}
	}

	return &menuProto.MenuListResponse{
		Foods: fListResp,
	}
}
