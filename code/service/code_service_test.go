package service

import (
	"context"
	"errors"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCodeServiceError(t *testing.T) {
	codeRepo := mocks.MockCodeRepository{}

	codeRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Code")).Return(errors.New("error message"))

	codeService := NewCodeService(&codeRepo)
	_, err := codeService.GenerateNewCode(context.Background(), "identifier", 8, false, false)

	assert.Error(t, err)
}

func TestNewCodeService(t *testing.T) {
	codeRepo := mocks.MockCodeRepository{}

	codeRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Code")).Return(nil)

	codeService := NewCodeService(&codeRepo)
	code, err := codeService.GenerateNewCode(context.Background(), "identifier", 8, false, false)

	assert.Nil(t, err)
	assert.Equal(t, "identifier", code.Identifier)
	assert.Len(t, code.Value, 8)
}
