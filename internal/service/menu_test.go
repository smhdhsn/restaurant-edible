package service

import (
	"testing"

	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetFoods(t *testing.T) {
	// Arrange
	f := model.Food{
		Title: randStr,
	}
	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("GetAvailableMeals").Return([]*model.Food{&f}, nil)

	sut := NewMenuService(fRepoMock)

	// Act
	fList, err := sut.GetFoods()

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, fList, &f)
}
