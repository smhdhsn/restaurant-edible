package service

import (
	"testing"

	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestOrderFood(t *testing.T) {
	// Arrange
	f := new(model.Food)
	f.ID = randUINT
	f.Title = randStr

	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("GetAvailableMeals").Return([]*model.Food{f}, nil)

	iRepoMock := new(mock.InventoryRepo)
	iRepoMock.On("Use", randUINT).Return(nil)

	sut := NewOrderService(fRepoMock, iRepoMock)

	// Act
	status, err := sut.OrderFood(randUINT)

	// Assert
	assert.NoError(t, err)
	assert.True(t, status)
}
