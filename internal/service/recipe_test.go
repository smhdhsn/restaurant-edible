package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mock"
)

func TestCreateRecipe(t *testing.T) {
	// Arrange
	fList := make([]*model.Food, 0)
	fList = append(fList, &model.Food{
		Title: randStr,
	})
	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("BatchInsert", fList).Return(nil)

	sut := NewRecipeService(fRepoMock)

	// Act
	err := sut.CreateRecipe(fList)

	// Assert
	assert.NoError(t, err)
}
