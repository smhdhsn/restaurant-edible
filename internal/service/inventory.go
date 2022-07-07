package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"

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

// Recycle is responsible for cleaning up the inventory from useless items.
func (s *InventoryServ) Recycle(rDTO *dto.Recycle) error {
	rEntity := singleRecycleDTOToEntity(rDTO)

	return s.iRepo.Recycle(rEntity)
}

// singleRecycleDTOToEntity is responsible for transforming a recycle dto into a recycle entity struct.
func singleRecycleDTOToEntity(rDTO *dto.Recycle) *entity.Recycle {
	return &entity.Recycle{
		Finished: rDTO.Finished,
		Expired:  rDTO.Expired,
	}
}

// Use is responsible for decreasing food's ingredient stocks inside database.
func (s *InventoryServ) Use(fDTO *dto.Food) error {
	fListEntity, err := s.fRepo.GetAvailable()
	if err != nil {
		return errors.Wrap(err, "error on fetching available foods from repository")
	}

	var isAvailable bool
	for _, f := range fListEntity {
		if f.ID == fDTO.ID {
			isAvailable = true
			break
		}
	}

	if !isAvailable {
		return serviceContract.ErrLackOfComponents
	}

	fEntity := singleFoodDTOToEntity(fDTO)

	err = s.iRepo.Use(fEntity)
	if err != nil {
		return errors.Wrap(err, "error on submitting use component stocks to repository")
	}

	return nil
}

// singleFoodDTOToEntity is responsible for transforming a food dto to food entity struct.
func singleFoodDTOToEntity(fDTO *dto.Food) *entity.Food {
	cListEntity := make([]*entity.Component, len(fDTO.Components))

	for i, cDTO := range fDTO.Components {
		cListEntity[i] = &entity.Component{
			ID:        cDTO.ID,
			Title:     cDTO.Title,
			CreatedAt: cDTO.CreatedAt,
			UpdatedAt: cDTO.UpdatedAt,
		}
	}

	return &entity.Food{
		ID:         fDTO.ID,
		Title:      fDTO.Title,
		Components: cListEntity,
		CreatedAt:  fDTO.CreatedAt,
		UpdatedAt:  fDTO.UpdatedAt,
	}
}

// Buy is responsible for buying food components for the inventory, if components' stock are finished or expired.
func (s *InventoryServ) Buy(bDTO *dto.Buy) error {
	cListEntity, err := s.cRepo.GetUnavailable()
	if err != nil {
		return errors.Wrap(err, "failed to get unavailable components")
	}

	if len(cListEntity) == 0 {
		return nil
	}

	iListEntity := componentEntityAndBuyDTOToInventoryEntity(cListEntity, bDTO)

	err = s.iRepo.Buy(iListEntity)
	if err != nil {
		return errors.Wrap(err, "failed to buy components")
	}

	return nil
}

// componentEntityAndBuyDTOToInventoryEntity is responsible for squashing a list of component entity and a buy dto into inventory entity struct.
func componentEntityAndBuyDTOToInventoryEntity(cListEntity []*entity.Component, bDTO *dto.Buy) []*entity.Inventory {
	iListEntity := make([]*entity.Inventory, len(cListEntity))

	for i, cEntity := range cListEntity {
		iListEntity[i] = &entity.Inventory{
			ComponentID: cEntity.ID,
			Stock:       bDTO.Stock,
			ExpiresAt:   bDTO.ExpiresAt,
		}
	}

	return iListEntity
}
