package service

import (
	"time"

	"github.com/pkg/errors"
	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository"
)

// InventoryService contains repositories that will be used within this service.
type InventoryService struct {
	iRepo repository.InventoryRepository
	cRepo repository.ComponentRepository
}

// BuyComponentsReq holds the details of buyed stocks.
type BuyComponentsReq struct {
	StockAmount uint
	BestBefore  time.Time
	ExpiresAt   time.Time
}

// NewInventoryService creates an inventory service with it's dependencies.
func NewInventoryService(iRepo repository.InventoryRepository, cRepo repository.ComponentRepository) *InventoryService {
	return &InventoryService{iRepo: iRepo, cRepo: cRepo}
}

// BuyComponents is responsible for buying food components for the inventory, if components' stock are finished or expired.
func (s *InventoryService) BuyComponents(req *BuyComponentsReq) error {
	cList, err := s.cRepo.GetUnavailable()
	if err != nil {
		return errors.Wrap(err, "failed to get unavailable components")
	}

	if len(cList) == 0 {
		return nil
	}

	iList := make([]*model.Inventory, 0)
	for _, c := range cList {
		iList = append(iList, &model.Inventory{
			ComponentID: c.ID,
			Stock:       req.StockAmount,
			BestBefore:  req.BestBefore,
			ExpiresAt:   req.ExpiresAt,
		})
	}

	err = s.iRepo.BuyStocks(iList)
	if err != nil {
		return errors.Wrap(err, "failed to buy components")
	}

	return nil
}
