package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mock"
)

func TestOrderFood(t *testing.T) {
	// Arrange
	f := new(model.Food)
	f.ID = randUINT
	f.Title = randStr

	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("GetAvailable").Return([]*model.Food{f}, nil)

	iRepoMock := new(mock.InventoryRepo)
	iRepoMock.On("Use", randUINT).Return(nil)

	sut := NewOrderService(fRepoMock, iRepoMock)

	// Act
	status, err := sut.OrderFood(randUINT)

	// Assert
	assert.NoError(t, err)
	assert.True(t, status)
}
