package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"

	log "github.com/smhdhsn/restaurant-edible/internal/logger"
	menuProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/menu"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
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
	fListEntity, err := s.menuServ.List()
	if err != nil {
		log.Error(err)
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := multipleFoodEntityToMenuResp(fListEntity)

	return resp, nil
}

// multipleFoodEntityToMenuResp is responsible for transforming a list of food entity to a list of food response struct.
func multipleFoodEntityToMenuResp(fListEntity []*entity.Food) *menuProto.MenuListResponse {
	fListResp := make([]*menuProto.Food, len(fListEntity))

	for i, fEntity := range fListEntity {
		cListResp := make([]*menuProto.Ingredient, len(fEntity.Components))

		for j, cEntity := range fEntity.Components {
			cListResp[j] = &menuProto.Ingredient{
				Title: cEntity.Title,
			}
		}

		fListResp[i] = &menuProto.Food{
			Id:          fEntity.ID,
			Title:       fEntity.Title,
			Ingredients: cListResp,
		}
	}

	resp := &menuProto.MenuListResponse{
		Foods: fListResp,
	}

	return resp
}
