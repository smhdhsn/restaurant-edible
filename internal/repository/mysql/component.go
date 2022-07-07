package mysql

import (
	"time"

	"gorm.io/gorm"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// component represents the component table's model.
type component struct {
	ID        uint32  `gorm:"primaryKey"`
	Title     string  `gorm:"unique;not null"`
	Foods     []*food `gorm:"many2many:food_components"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ComponentRepo contains repository's database connection.
type ComponentRepo struct {
	model component
	db    *gorm.DB
}

// NewComponentRepo creates an instance of the repository with database connection.
func NewComponentRepository(db *gorm.DB) repositoryContract.ComponentRepository {
	return &ComponentRepo{
		model: component{},
		db:    db,
	}
}

// GetUnavailable is responsible for getting food components that are not avaiable.
func (r *ComponentRepo) GetUnavailable() ([]*dto.ComponentDTO, error) {
	result := make([]*component, 0)

	tx := r.db.
		Table("components AS c").
		Where(
			"c.id NOT IN (?)",
			availableInventoryItems(r.db),
		).
		Find(&result)

	cListDTO := make([]*dto.ComponentDTO, len(result))
	for i, c := range result {
		cListDTO[i] = &dto.ComponentDTO{
			ID:        c.ID,
			Title:     c.Title,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		}
	}

	return cListDTO, tx.Error
}
