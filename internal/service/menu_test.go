package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smhdhsn/restaurant-edible/internal/model"
	"github.com/smhdhsn/restaurant-edible/internal/repository/mock"
	"github.com/smhdhsn/restaurant-edible/util/random"
)

func TestGetFoods(t *testing.T) {
	// Arrange
	randTitle := random.GenerateString(5)

	f := model.Food{
		Title: randTitle,
	}
	fRepoMock := new(mock.FoodRepo)
	fRepoMock.On("GetAvailable").Return([]*model.Food{&f}, nil)

	sut := NewMenuService(fRepoMock)

	// Act
	fList, err := sut.List()

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, fList, &f)
}
