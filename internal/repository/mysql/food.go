package mysql

import (
	"time"

	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-edible/internal/service/dto"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// food represents the food table's model.
type food struct {
	ID         uint32       `gorm:"primaryKey"`
	Title      string       `gorm:"unique;not null"`
	Components []*component `gorm:"many2many:food_components"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// FoodRepo contains repository's database connection.
type FoodRepo struct {
	model food
	db    *gorm.DB
}

// NewFoodRepo creates an instance of the repository with database connection.
func NewFoodRepository(db *gorm.DB) repositoryContract.FoodRepository {
	return &FoodRepo{
		model: food{},
		db:    db,
	}
}

// GetAvailable gets foods that their components are available.
func (r *FoodRepo) GetAvailable() ([]*dto.FoodDTO, error) {
	result := make([]*food, 0)

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

	fListDTO := make([]*dto.FoodDTO, len(result))
	for i, f := range result {
		cListDTO := make([]*dto.ComponentDTO, len(f.Components))
		for j, c := range f.Components {
			cListDTO[j] = &dto.ComponentDTO{
				ID:        c.ID,
				Title:     c.Title,
				CreatedAt: c.CreatedAt,
				UpdatedAt: c.UpdatedAt,
			}
		}

		fListDTO[i] = &dto.FoodDTO{
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
func (r *FoodRepo) BatchInsert(fListDTO []*dto.FoodDTO) error {
	fList := make([]*food, len(fListDTO))
	for i, f := range fListDTO {
		cList := make([]*component, len(f.Components))
		for i, c := range f.Components {
			cList[i] = &component{
				Title: c.Title,
			}
		}

		fList[i] = &food{
			Title:      f.Title,
			Components: cList,
		}
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, f := range fList {
			for _, c := range f.Components {
				if err := tx.Where(component{Title: c.Title}).FirstOrCreate(&c).Error; err != nil {
					return err
				}
			}

			if err := tx.Where(food{Title: f.Title}).FirstOrCreate(&f).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
