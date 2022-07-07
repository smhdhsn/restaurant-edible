package handler

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	inventoryProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/inventory"
	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// InventoryHandler contains services that can be used within inventory handler.
type InventoryHandler struct {
	inventoryServ serviceContract.InventoryService
}

// NewInventoryHandler creates a new menu handler.
func NewInventoryHandler(is serviceContract.InventoryService) inventoryProto.EdibleInventoryServiceServer {
	return &InventoryHandler{
		inventoryServ: is,
	}
}

// Recycle is responsible for recycling finished and/or expired items from inventory.
func (s *InventoryHandler) Recycle(ctx context.Context, req *inventoryProto.InventoryRecycleRequest) (*inventoryProto.InventoryRecycleResponse, error) {
	finished := req.GetRecycleFinished()
	expired := req.GetRecycleExpired()

	err := s.inventoryServ.Recycle(finished, expired)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := inventoryProto.InventoryRecycleResponse{
		Status: true,
	}

	return &resp, nil
}

// Use is responsible for decreasing item's inside from inventory.
func (s *InventoryHandler) Use(ctx context.Context, req *inventoryProto.InventoryUseRequest) (*inventoryProto.InventoryUseResponse, error) {
	fReq := dto.FoodDTO{
		ID: req.GetFoodId(),
	}

	err := s.inventoryServ.Use(&fReq)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := inventoryProto.InventoryUseResponse{
		Status: true,
	}

	return &resp, nil
}

// Buy is responsible for increasing item's stock inside inventory.
func (s *InventoryHandler) Buy(ctx context.Context, req *inventoryProto.InventoryBuyRequest) (*inventoryProto.InventoryBuyResponse, error) {
	amount := req.GetAmount()
	expiresAt := time.Unix(req.GetExpiresAt(), 0)

	err := s.inventoryServ.Buy(amount, expiresAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := inventoryProto.InventoryBuyResponse{
		Status: true,
	}

	return &resp, nil
}
