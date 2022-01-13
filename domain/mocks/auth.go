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

type MockAuthValidator struct {
	mock.Mock
}

func (mav *MockAuthValidator) Validate(ctx context.Context, a *domain.Auth) (domain.IsValid, domain.Message, error) {
	args := mav.Called(ctx, a)
	return domain.IsValid(args.Bool(0)), domain.Message(args.String(1)), args.Error(2)
}

func (mav *MockAuthValidator) ValidateLogin(ctx context.Context, login string) (domain.IsValid, domain.Message, error) {
	args := mav.Called(ctx, login)
	return domain.IsValid(args.Bool(0)), domain.Message(args.String(1)), args.Error(2)
}

type MockForgotPassResetValidator struct {
	mock.Mock
}

func (mfprv *MockForgotPassResetValidator) Validate(ctx context.Context, fpr *domain.ForgotPassReset) (domain.IsValid, domain.Message, error) {
	args := mfprv.Called(ctx, fpr)
	return domain.IsValid(args.Bool(0)), domain.Message(args.String(1)), args.Error(2)
}

type MockAuthService struct {
	mock.Mock
}

func (mas *MockAuthService) EncodePass(ctx context.Context, pass string) string {
	args := mas.Called(ctx, pass)
	return args.String(0)
}

func (mas *MockAuthService) PassIsEqualHashedPass(ctx context.Context, pass string, hashedPass string) bool {
	args := mas.Called(ctx, pass, hashedPass)
	return args.Bool(0)
}

type MockAuthRepository struct {
	mock.Mock
}

func (mar *MockAuthRepository) GetByLogin(ctx context.Context, login string) (*domain.Auth, error) {
	args := mar.Called(ctx, login)
	return &domain.Auth{Login: args.String(0), Password: args.String(1)}, args.Error(2)
}
