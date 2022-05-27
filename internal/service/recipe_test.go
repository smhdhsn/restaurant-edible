package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smhdhsn/restaurant-edible/internal/model"
	"github.com/smhdhsn/restaurant-edible/internal/repository/mock"
	"github.com/smhdhsn/restaurant-edible/util/random"
)

func TestCreateRecipe(t *testing.T) {
	// Arrange
	randTitle := random.GenerateString(5)

	fList := make([]*model.Food, 0)
	fList = append(fList, &model.Food{
		Title: randTitle,
	})

	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("BatchInsert", fList).Return(nil)

	sut := NewRecipeService(fRepoMock)

	// Act
	err := sut.CreateRecipe(fList)

	// Assert
	assert.NoError(t, err)
}
