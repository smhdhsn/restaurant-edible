package handler

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-edible/internal/service/dto"

	log "github.com/smhdhsn/restaurant-edible/internal/logger"
	inventoryProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/inventory"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
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
	rDTO := singleRecycleReqToDTO(req)

	err := s.inventoryServ.Recycle(rDTO)
	if err != nil {
		log.Error(err)
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := new(inventoryProto.InventoryRecycleResponse)

	return resp, nil
}

// singleRecycleReqToDTO is responsible for transforming a recycle proto request to recycle dto struct.
func singleRecycleReqToDTO(req *inventoryProto.InventoryRecycleRequest) *dto.Recycle {
	return &dto.Recycle{
		Finished: req.GetRecycleFinished(),
		Expired:  req.GetRecycleExpired(),
	}
}

// Use is responsible for decreasing item's inside from inventory.
func (s *InventoryHandler) Use(ctx context.Context, req *inventoryProto.InventoryUseRequest) (*inventoryProto.InventoryUseResponse, error) {
	fDTO := singleUseReqToFoodDTO(req)

	err := s.inventoryServ.Use(fDTO)
	if err != nil {
		if errors.Is(err, serviceContract.ErrLackOfComponents) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		log.Error(err)
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := new(inventoryProto.InventoryUseResponse)

	return resp, nil
}

// singleUseReqToFoodDTO is responsible for transforming a use proto request to food dto struct.
func singleUseReqToFoodDTO(req *inventoryProto.InventoryUseRequest) *dto.Food {
	return &dto.Food{
		ID: req.GetFoodId(),
	}
}

// Buy is responsible for increasing item's stock inside inventory.
func (s *InventoryHandler) Buy(ctx context.Context, req *inventoryProto.InventoryBuyRequest) (*inventoryProto.InventoryBuyResponse, error) {
	bDTO := singleBuyReqToDTO(req)

	err := s.inventoryServ.Buy(bDTO)
	if err != nil {
		log.Error(err)
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := new(inventoryProto.InventoryBuyResponse)

	return resp, nil
}

// singleBuyReqToDTO is responsible for transforming a buy proto request to buy dto struct.
func singleBuyReqToDTO(req *inventoryProto.InventoryBuyRequest) *dto.Buy {
	return &dto.Buy{
		Stock:     req.GetAmount(),
		ExpiresAt: time.Unix(req.GetExpiresAt(), 0),
	}
}
