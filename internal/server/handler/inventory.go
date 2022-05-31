package handler

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-edible/internal/model"

	eipb "github.com/smhdhsn/restaurant-edible/internal/protos/edible/inventory"
	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
)

// InventoryHandler contains services that can be used within inventory handler.
type InventoryHandler struct {
	inventoryServ serviceContract.InventoryService
}

// NewInventoryHandler creates a new menu handler.
func NewInventoryHandler(is serviceContract.InventoryService) eipb.EdibleInventoryServiceServer {
	return &InventoryHandler{
		inventoryServ: is,
	}
}

// Recycle is responsible for recycling finished and/or expired items from inventory.
func (s *InventoryHandler) Recycle(ctx context.Context, req *eipb.InventoryRecycleRequest) (*eipb.InventoryRecycleResponse, error) {
	finished := req.GetRecycleFinished()
	expired := req.GetRecycleExpired()

	err := s.inventoryServ.Recycle(finished, expired)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := eipb.InventoryRecycleResponse{
		Status: true,
	}

	return &resp, nil
}

// Use is responsible for decreasing item's inside from inventory.
func (s *InventoryHandler) Use(ctx context.Context, req *eipb.InventoryUseRequest) (*eipb.InventoryUseResponse, error) {
	fReq := model.FoodDTO{
		ID: req.GetFoodId(),
	}

	err := s.inventoryServ.Use(&fReq)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := eipb.InventoryUseResponse{
		Status: true,
	}

	return &resp, nil
}

// Buy is responsible for increasing item's stock inside inventory.
func (s *InventoryHandler) Buy(ctx context.Context, req *eipb.InventoryBuyRequest) (*eipb.InventoryBuyResponse, error) {
	amount := req.GetAmount()
	expiresAt := time.Unix(req.GetExpiresAt(), 0)

	err := s.inventoryServ.Buy(amount, expiresAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := eipb.InventoryBuyResponse{
		Status: true,
	}

	return &resp, nil
}
