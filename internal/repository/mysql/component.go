package mysql

import (
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// component represents the component table's model.
type component struct {
	ID        uint32  `gorm:"primaryKey"`
	Title     string  `gorm:"unique;not null"`
	Foods     []*food `gorm:"many2many:foods_components"`
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
func (r *ComponentRepo) GetUnavailable() ([]*entity.Component, error) {
	cListModel := make([]*component, 0)

	err := r.db.
		Table("components AS c").
		Where(
			"c.id NOT IN (?)",
			availableInventoryItems(r.db),
		).
		Find(&cListModel).
		Error
	if err != nil {
		return nil, errors.Wrap(err, "error on fetching unavailable foods from database")
	}

	cListEntity := multipleComponentModelToEntity(cListModel)

	return cListEntity, nil
}

// multipleComponentModelToEntity is responsible for transforming a list of component model to a list of component entity struct.
func multipleComponentModelToEntity(cListModel []*component) []*entity.Component {
	cListEntity := make([]*entity.Component, len(cListModel))

	for i, cModel := range cListModel {
		cListEntity[i] = &entity.Component{
			ID:        cModel.ID,
			Title:     cModel.Title,
			CreatedAt: cModel.CreatedAt,
			UpdatedAt: cModel.UpdatedAt,
		}
	}

	return cListEntity
}
