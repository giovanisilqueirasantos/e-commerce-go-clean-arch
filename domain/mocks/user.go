package mocks

import (
	"context"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockUserValidator struct {
	mock.Mock
}

func (muv *MockUserValidator) Validate(ctx context.Context, u *domain.User) (domain.IsValid, domain.Message) {
	args := muv.Called(ctx, u)
	return domain.IsValid(args.Bool(0)), domain.Message(args.String(1))
}

type MockUserRepository struct {
	mock.Mock
}

func (mur *MockUserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	args := mur.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return &domain.User{Email: args.String(0), FirstName: args.String(1), LastName: args.String(2), PhoneNumber: args.String(3), Address: args.String(4)}, args.Error(5)
}

func (mur *MockUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := mur.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return &domain.User{Email: args.String(0), FirstName: args.String(1), LastName: args.String(2), PhoneNumber: args.String(3), Address: args.String(4)}, args.Error(5)
}

func (mur *MockUserRepository) GetByLogin(ctx context.Context, login string) (*domain.User, error) {
	args := mur.Called(ctx, login)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return &domain.User{Email: args.String(0), FirstName: args.String(1), LastName: args.String(2), PhoneNumber: args.String(3), Address: args.String(4)}, args.Error(5)
}

func (mur *MockUserRepository) Update(ctx context.Context, u *domain.User) error {
	args := mur.Called(ctx, u)
	return args.Error(0)
}
