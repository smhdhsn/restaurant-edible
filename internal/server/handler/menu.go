package handler

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	menuProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/menu"
	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
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
	fListDTO, err := s.menuServ.List()
	if err != nil {
		if errors.Is(err, repositoryContract.ErrEmptyResult) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	fList := make([]*menuProto.Food, len(fListDTO))
	for i, f := range fListDTO {
		cList := make([]*menuProto.Ingredient, len(f.Components))
		for j, c := range f.Components {
			cList[j] = &menuProto.Ingredient{
				Title: c.Title,
			}
		}

		fList[i] = &menuProto.Food{
			Id:          f.ID,
			Title:       f.Title,
			Ingredients: cList,
		}
	}

	resp := &menuProto.MenuListResponse{
		Foods: fList,
	}

	return resp, nil
}
