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

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(nil, errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, nil, mockAuthRepo)

	_, errToken := authUseCase.Login(context.Background(), &mockAuth)

	assert.Error(t, errToken)
}

func TestLoginAuthNil(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(nil, nil)

	authUseCase := NewAuthUseCase(nil, nil, mockAuthRepo)

	_, errToken := authUseCase.Login(context.Background(), &mockAuth)

	assert.Error(t, errToken)
}

func TestLoginPassIsEqualHashedPassError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockAuthService := new(mocks.MockAuthService)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "invalid password"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return("valid login", "valid password", nil)

	mockAuthService.On("PassIsEqualHashedPass", mock.Anything, mockAuth.Password, "valid password").Return(false)

	authUseCase := NewAuthUseCase(mockAuthService, nil, mockAuthRepo)

	_, errToken := authUseCase.Login(context.Background(), &mockAuth)

	assert.Error(t, errToken)
}

func TestLoginSignTokenError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockAuthService := new(mocks.MockAuthService)
	mockTokenService := new(mocks.MockTokenService)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(mockAuth.Login, mockAuth.Password, nil)

	mockAuthService.On("PassIsEqualHashedPass", mock.Anything, mockAuth.Password, mockAuth.Password).Return(true)

	var thirtyDaysInMinutes int64 = 43200

	tokenInfo := domain.TokenInfo{Info: mockAuth.Login}

	mockTokenService.On("Sign", mock.Anything, tokenInfo, thirtyDaysInMinutes).Return("", errors.New("error message"))

	authUseCase := NewAuthUseCase(mockAuthService, mockTokenService, mockAuthRepo)

	_, errToken := authUseCase.Login(context.Background(), &mockAuth)

	assert.Error(t, errToken)
}

func TestLoginSuccess(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockAuthService := new(mocks.MockAuthService)
	mockTokenService := new(mocks.MockTokenService)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(mockAuth.Login, mockAuth.Password, nil)

	mockAuthService.On("PassIsEqualHashedPass", mock.Anything, mockAuth.Password, mockAuth.Password).Return(true)

	var thirtyDaysInMinutes int64 = 43200

	tokenInfo := domain.TokenInfo{Info: mockAuth.Login}

	mockTokenService.On("Sign", mock.Anything, tokenInfo, thirtyDaysInMinutes).Return("valid token", nil)

	authUseCase := NewAuthUseCase(mockAuthService, mockTokenService, mockAuthRepo)

	token, errToken := authUseCase.Login(context.Background(), &mockAuth)

	assert.Nil(t, errToken)
	assert.Equal(t, token, domain.Token("valid token"))
}
