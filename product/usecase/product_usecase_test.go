package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetError(t *testing.T) {
	mockProductRepo := new(mocks.MockProductRepository)

	mockProductRepo.On("GetByUUID", mock.Anything, "uuid").Return(nil, errors.New("error message"))

	productUseCase := NewProductUseCase(mockProductRepo)

	_, err := productUseCase.Get(context.Background(), "uuid")

	assert.Error(t, err)
}

func TestGetNotExists(t *testing.T) {
	mockProductRepo := new(mocks.MockProductRepository)

	mockProductRepo.On("GetByUUID", mock.Anything, "uuid").Return(nil, nil)

	productUseCase := NewProductUseCase(mockProductRepo)

	product, err := productUseCase.Get(context.Background(), "uuid")

	assert.Nil(t, product)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	mockProductRepo := new(mocks.MockProductRepository)

	mockProductRepo.On("GetByUUID", mock.Anything, "uuid").Return(1, "uuid", 2, "picturepath", "name", "detail", true, "color", "black", nil)

	productUseCase := NewProductUseCase(mockProductRepo)

	product, err := productUseCase.Get(context.Background(), "uuid")

	assert.Nil(t, err)
	assert.Equal(t, int64(1), product.ID)
	assert.Equal(t, "uuid", product.UUID)
	assert.Equal(t, float32(2), product.Rate)
	assert.Equal(t, "picturepath", product.Pictures[0])
	assert.Equal(t, "name", product.Name)
	assert.Equal(t, "detail", product.Detail)
	assert.Equal(t, true, product.Favorite)
	assert.Equal(t, "color", product.Attributes[0].Label)
	assert.Equal(t, "black", product.Attributes[0].Values[0])
}
