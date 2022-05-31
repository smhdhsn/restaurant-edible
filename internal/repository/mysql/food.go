package mysql

import (
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-edible/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// FoodRepo contains repository's database connection.
type FoodRepo struct {
	model model.Food
	db    *gorm.DB
}

// NewFoodRepo creates an instance of the repository with database connection.
func NewFoodRepository(db *gorm.DB, m model.Food) repositoryContract.FoodRepository {
	return &FoodRepo{
		model: m,
		db:    db,
	}
}

// GetAvailable gets foods that their components are available.
func (r *FoodRepo) GetAvailable() (model.FoodListDTO, error) {
	result := make([]*model.Food, 0)

	tx := r.db.
		Table("foods AS f").
		Preload("Components").
		Where(
			"f.id NOT IN (?)",
			unavailableFoods(r.db),
		).
		Find(&result)

	if len(result) == 0 {
		return nil, repositoryContract.ErrEmptyResult
	}

	fListDTO := make(model.FoodListDTO, len(result))
	for i, f := range result {
		cListDTO := make(model.ComponentListDTO, len(f.Components))
		for j, c := range f.Components {
			cListDTO[j] = &model.ComponentDTO{
				ID:        c.ID,
				Title:     c.Title,
				CreatedAt: c.CreatedAt,
				UpdatedAt: c.UpdatedAt,
			}
		}

		fListDTO[i] = &model.FoodDTO{
			ID:         f.ID,
			Title:      f.Title,
			Components: cListDTO,
			CreatedAt:  f.CreatedAt,
			UpdatedAt:  f.UpdatedAt,
		}
	}

	return fListDTO, tx.Error
}

// BatchInsert is responsible for storing a chunk of data inside database.
func (r *FoodRepo) BatchInsert(fListDTO model.FoodListDTO) error {
	fList := make([]*model.Food, len(fListDTO))
	for i, f := range fListDTO {
		cList := make([]*model.Component, len(f.Components))
		for i, c := range f.Components {
			cList[i] = &model.Component{
				Title: c.Title,
			}
		}

		fList[i] = &model.Food{
			Title:      f.Title,
			Components: cList,
		}
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, f := range fList {
			for _, c := range f.Components {
				if err := tx.Where(model.Component{Title: c.Title}).FirstOrCreate(&c).Error; err != nil {
					return err
				}
			}

			if err := tx.Where(model.Food{Title: f.Title}).FirstOrCreate(&f).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
