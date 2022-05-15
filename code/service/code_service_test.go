package service

import (
	"context"
	"errors"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
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

func TestValidateCodeGetByValueError(t *testing.T) {
	codeRepo := mocks.MockCodeRepository{}

	codeRepo.On("GetByValue", mock.Anything, "code value").Return(nil, errors.New("error message"))

	codeService := NewCodeService(&codeRepo)
	_, err := codeService.ValidateCode(context.Background(), &domain.Code{Value: "code value"})

	assert.Error(t, err)
}

func TestValidateCodeDeleteByValueError(t *testing.T) {
	codeRepo := mocks.MockCodeRepository{}

	codeRepo.On("GetByValue", mock.Anything, "code value").Return("code value", "code identifier", nil)
	codeRepo.On("DeleteByValue", mock.Anything, "code value").Return(errors.New("error message"))

	codeService := NewCodeService(&codeRepo)
	_, err := codeService.ValidateCode(context.Background(), &domain.Code{Identifier: "code identifier", Value: "code value"})

	assert.Error(t, err)
}

func TestValidateCodeInvalidCode(t *testing.T) {
	codeRepo := mocks.MockCodeRepository{}

	codeRepo.On("GetByValue", mock.Anything, "code wrong value").Return("code value", "code identifier", nil)

	codeService := NewCodeService(&codeRepo)
	isValid, err := codeService.ValidateCode(context.Background(), &domain.Code{Identifier: "code wrong identifier", Value: "code wrong value"})

	assert.False(t, bool(isValid))
	assert.NoError(t, err)
}

func TestValidateCode(t *testing.T) {
	codeRepo := mocks.MockCodeRepository{}

	codeRepo.On("GetByValue", mock.Anything, "code value").Return("code value", "code identifier", nil)
	codeRepo.On("DeleteByValue", mock.Anything, "code value").Return(nil)

	codeService := NewCodeService(&codeRepo)
	isValid, err := codeService.ValidateCode(context.Background(), &domain.Code{Identifier: "code identifier", Value: "code value"})

	assert.True(t, bool(isValid))
	assert.NoError(t, err)
}
