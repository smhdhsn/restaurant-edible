package inventory

import (
	"time"

	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-menu/internal/model"

	cRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/component"
	iRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/inventory"
)

// InventoryService contains repositories that will be used within this service.
type InventoryService struct {
	iRepo iRepoContract.InventoryRepository
	cRepo cRepoContract.ComponentRepository
}

// BuyComponentsReq holds the details of buyed stocks.
type BuyComponentsReq struct {
	StockAmount uint32
	BestBefore  time.Time
	ExpiresAt   time.Time
}

// NewInventoryService creates an inventory service with it's dependencies.
func NewInventoryService(iRepo iRepoContract.InventoryRepository, cRepo cRepoContract.ComponentRepository) *InventoryService {
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

	err = s.iRepo.Buy(iList)
	if err != nil {
		return errors.Wrap(err, "failed to buy components")
	}

	return nil
}

// Recycle is responsible for cleaning up the inventory from useless items.
func (s *InventoryService) Recycle(req iRepoContract.RecycleReq) error {
	return s.iRepo.Clean(req)
}
