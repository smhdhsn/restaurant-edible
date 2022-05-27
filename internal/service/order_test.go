package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smhdhsn/restaurant-edible/internal/model"
	"github.com/smhdhsn/restaurant-edible/internal/repository/mock"
	"github.com/smhdhsn/restaurant-edible/util/random"
)

func TestOrderFood(t *testing.T) {
	// Arrange
	randID := random.GenerateUint32(1, 100)
	randTitle := random.GenerateString(5)

	f := model.Food{
		ID:    randID,
		Title: randTitle,
	}

	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("GetAvailable").Return([]*model.Food{&f}, nil)

	iRepoMock := new(mock.InventoryRepo)
	iRepoMock.On("Use", randID).Return(nil)

	sut := NewOrderService(fRepoMock, iRepoMock)

	// Act
	status, err := sut.OrderFood(randID)

	// Assert
	assert.NoError(t, err)
	assert.True(t, status)
}