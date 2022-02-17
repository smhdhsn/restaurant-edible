package service

import (
	"testing"

	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateRecipe(t *testing.T) {
	// Arrange
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
