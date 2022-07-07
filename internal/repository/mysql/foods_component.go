package mysql

import (
	"gorm.io/gorm"
)

// unavailableFoods is the subquery responsible for getting unavailable food components by "food_id".
func unavailableFoods(db *gorm.DB) *gorm.DB {
	return db.
		Table("foods_components AS fc").
		Select("fc.food_id").
		Where(
			"fc.component_id NOT IN (?)",
			availableInventoryItems(db),
		)
}

// componentsOfFood is the subquery responsible for getting components related to given "food_id".
func componentsOfFood(db *gorm.DB, foodID uint32) *gorm.DB {
	return db.
		Table("foods_components AS fc").
		Select("fc.component_id").
		Where("fc.food_id = ?", foodID)
}
