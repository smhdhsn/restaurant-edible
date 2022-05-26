package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository/mock"
)

func TestGetFoods(t *testing.T) {
	// Arrange
	f := model.Food{
		Title: randStr,
	}
	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("GetAvailable").Return([]*model.Food{&f}, nil)

	sut := NewMenuService(fRepoMock)

	// Act
	fList, err := sut.GetFoods()

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, fList, &f)
}
