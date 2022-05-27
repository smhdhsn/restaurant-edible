package service

import (
	"time"

	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-edible/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
)

// InventoryServ contains repositories that will be used within this service.
type InventoryServ struct {
	iRepo repositoryContract.InventoryRepository
	cRepo repositoryContract.ComponentRepository
	fRepo repositoryContract.FoodRepository
}

// NewInventoryService creates an inventory service with it's dependencies.
func NewInventoryService(
	ir repositoryContract.InventoryRepository,
	cr repositoryContract.ComponentRepository,
	fr repositoryContract.FoodRepository,
) serviceContract.InventoryService {
	return &InventoryServ{
		iRepo: ir,
		cRepo: cr,
		fRepo: fr,
	}
}

// Buy is responsible for buying food components for the inventory, if components' stock are finished or expired.
func (s *InventoryServ) Buy(stock uint32, expiresAt time.Time) error {
	cList, err := s.cRepo.GetUnavailable()
	if err != nil {
		return errors.Wrap(err, "failed to get unavailable components")
	}

	if len(cList) == 0 {
		return nil
	}

	iListDTO := make([]*model.InventoryDTO, len(cList))
	for i, c := range cList {
		iListDTO[i] = &model.InventoryDTO{
			ComponentID: c.ID,
			Stock:       stock,
			ExpiresAt:   expiresAt,
		}
	}

	err = s.iRepo.Buy(iListDTO)
	if err != nil {
		return errors.Wrap(err, "failed to buy components")
	}

	return nil
}

// GetFood is responsible for fetching available meals from database.
func (s *InventoryServ) Use(fDTO *model.FoodDTO) error {
	foods, err := s.fRepo.GetAvailable()
	if err != nil {
		return errors.Wrap(err, "failed to get available foods")
	}

	var isAvailable bool
	for _, f := range foods {
		if f.ID == fDTO.ID {
			isAvailable = true
			break
		}
	}

	if !isAvailable {
		return errors.New("requested order cannot be fulfilled because of the lack of components")
	}

	err = s.iRepo.Use(fDTO)
	if err != nil {
		return errors.Wrap(err, "failed to use components")
	}

	return nil
}

// Recycle is responsible for cleaning up the inventory from useless items.
func (s *InventoryServ) Recycle(finished, expired bool) error {
	return s.iRepo.Recycle(finished, expired)
}
