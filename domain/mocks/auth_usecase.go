package mocks

import (
	"context"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockAuthUsecase struct {
	mock.Mock
}

func (m *MockAuthUsecase) Login(ctx context.Context, a *domain.Auth) (domain.Token, error) {
	args := m.Called(ctx, a)
	return domain.Token(args.String(0)), args.Error(1)
}

func (m *MockAuthUsecase) SignUp(ctx context.Context, a *domain.Auth, u *domain.User) error {
	args := m.Called(ctx, a, u)
	return args.Error(0)
}

func (m *MockAuthUsecase) ForgotPassCode(ctx context.Context, login string) error {
	args := m.Called(ctx, login)
	return args.Error(0)
}

func (m *MockAuthUsecase) ForgotPassReset(ctx context.Context, fpr *domain.ForgotPassReset) (domain.Token, error) {
	args := m.Called(ctx, fpr)
	return domain.Token(args.String(0)), args.Error(1)
}
