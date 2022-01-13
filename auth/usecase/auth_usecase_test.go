package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginGetByLoginError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return("", "", errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, nil, mockAuthRepo)

	_, errToken := authUseCase.Login(context.Background(), &mockAuth)

	assert.Error(t, errToken)
}
