package mocks

import (
	"context"

	"github.com/skeey/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockAuthUsecase struct {
	mock.Mock
}

func (m *MockAuthUsecase) Login(ctx context.Context, a *domain.Auth) (domain.Token, error) {
	args := m.Called(ctx, a)
	return domain.Token(args.String(0)), args.Error(1)
}

func (m *MockAuthUsecase) SignUp(ctx context.Context, a *domain.Auth, u *domain.User) (domain.OK, error) {
	args := m.Called(ctx, a, u)
	return domain.OK(args.Bool(0)), args.Error(1)
}

func (m *MockAuthUsecase) ForgotPassword(ctx context.Context, a *domain.Auth) (domain.OK, error) {
	args := m.Called(ctx, a)
	return domain.OK(args.Bool(0)), args.Error(1)
}
