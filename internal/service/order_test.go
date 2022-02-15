package service

import (
	"testing"

	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository/mock"
	"github.com/stretchr/testify/assert"
)

// Random data for testing.
var (
	randFoodID = uint(1)
)

func TestOrderFood(t *testing.T) {
	// Arrange
	f := model.Food{}
	f.ID = randFoodID
	f.Title = randTitle

	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("GetAvailableMeals").Return([]*model.Food{&f}, nil)

	iRepoMock := new(mock.InventoryRepo)
	iRepoMock.On("UseStocks", randFoodID).Return(nil)

	sut := NewOrderService(fRepoMock, iRepoMock)

	// Act
	status, err := sut.OrderFood(randFoodID)

	// Assert
	assert.NoError(t, err)
	assert.True(t, status)
}
