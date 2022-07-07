package mysql

import (
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// food represents the food table's model.
type food struct {
	ID         uint32       `gorm:"primaryKey"`
	Title      string       `gorm:"unique;not null"`
	Components []*component `gorm:"many2many:foods_components"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// FoodRepo contains repository's database connection.
type FoodRepo struct {
	model food
	db    *gorm.DB
}

// NewFoodRepository creates an instance of the repository with database connection.
func NewFoodRepository(db *gorm.DB) repositoryContract.FoodRepository {
	return &FoodRepo{
		model: food{},
		db:    db,
	}
}

// GetAvailable gets foods that their components are available.
func (r *FoodRepo) GetAvailable() ([]*entity.Food, error) {
	fListModel := make([]*food, 0)

	err := r.db.
		Table("foods AS f").
		Preload("Components").
		Where(
			"f.id NOT IN (?)",
			unavailableFoods(r.db),
		).
		Find(&fListModel).
		Error
	if err != nil {
		return nil, errors.Wrap(err, "error on fetching available foods from database")
	}

	fListEntity := multipleFoodModelToEntity(fListModel)

	return fListEntity, nil
}

// multipleFoodModelToEntity is responsible for transforming a list of food model to a list of food entity struct.
func multipleFoodModelToEntity(fListModel []*food) []*entity.Food {
	fListEntity := make([]*entity.Food, len(fListModel))

	for i, fModel := range fListModel {
		cListEntity := make([]*entity.Component, len(fModel.Components))

		for j, cModel := range fModel.Components {
			cListEntity[j] = &entity.Component{
				ID:        cModel.ID,
				Title:     cModel.Title,
				CreatedAt: cModel.CreatedAt,
				UpdatedAt: cModel.UpdatedAt,
			}
		}

		fListEntity[i] = &entity.Food{
			ID:         fModel.ID,
			Title:      fModel.Title,
			Components: cListEntity,
			CreatedAt:  fModel.CreatedAt,
			UpdatedAt:  fModel.UpdatedAt,
		}
	}

	return fListEntity
}

// BatchInsert is responsible for storing a chunk of data inside database.
func (r *FoodRepo) BatchInsert(fListEntity []*entity.Food) error {
	fListModel := multipleFoodEntityToModel(fListEntity)

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, f := range fListModel {
			for _, c := range f.Components {
				if err := tx.Where(component{Title: c.Title}).FirstOrCreate(&c).Error; err != nil {
					return errors.Wrap(err, "error on calling first or create on component model")
				}
			}

			if err := tx.Where(food{Title: f.Title}).FirstOrCreate(&f).Error; err != nil {
				return errors.Wrap(err, "error on calling first or create on food model")
			}
		}

		return nil
	})
}

// multipleFoodEntityToModel is responsible for transforming a list of food entity to a list of food model struct.
func multipleFoodEntityToModel(fListEntity []*entity.Food) []*food {
	fListModel := make([]*food, len(fListEntity))

	for i, fEntity := range fListEntity {
		cListModel := make([]*component, len(fEntity.Components))

		for i, c := range fEntity.Components {
			cListModel[i] = &component{
				Title: c.Title,
			}
		}

		fListModel[i] = &food{
			Title:      fEntity.Title,
			Components: cListModel,
		}
	}

	return fListModel
}
