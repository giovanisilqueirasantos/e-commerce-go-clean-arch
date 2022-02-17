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

func (m *MockAuthUsecase) SignUp(ctx context.Context, a *domain.Auth, u *domain.User) (domain.Token, error) {
	args := m.Called(ctx, a, u)
	return domain.Token(args.String(0)), args.Error(1)
}

func (m *MockAuthUsecase) ForgotPassCode(ctx context.Context, login string) error {
	args := m.Called(ctx, login)
	return args.Error(0)
}

func (m *MockAuthUsecase) ForgotPassReset(ctx context.Context, code *domain.Code, newPass string) (domain.Token, error) {
	args := m.Called(ctx, code, newPass)
	return domain.Token(args.String(0)), args.Error(1)
}

type MockAuthValidator struct {
	mock.Mock
}

func (mav *MockAuthValidator) Validate(ctx context.Context, a *domain.Auth) (domain.IsValid, domain.Message) {
	args := mav.Called(ctx, a)
	return domain.IsValid(args.Bool(0)), domain.Message(args.String(1))
}

func (mav *MockAuthValidator) ValidateLogin(ctx context.Context, login string) (domain.IsValid, domain.Message) {
	args := mav.Called(ctx, login)
	return domain.IsValid(args.Bool(0)), domain.Message(args.String(1))
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
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return &domain.Auth{Login: args.String(0), Password: args.String(1)}, args.Error(2)
}

func (mar *MockAuthRepository) StoreWithUser(ctx context.Context, a *domain.Auth, u *domain.User) error {
	args := mar.Called(ctx, a, u)
	return args.Error(0)
}

func (mar *MockAuthRepository) Update(ctx context.Context, a *domain.Auth) error {
	args := mar.Called(ctx, a)
	return args.Error(0)
}
