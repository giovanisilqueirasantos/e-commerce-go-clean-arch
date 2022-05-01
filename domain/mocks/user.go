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

func (mur *MockUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := mur.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return &domain.User{ID: int64(args.Int(0)), UUID: args.String(1), Email: args.String(2), FirstName: args.String(3), LastName: args.String(4), PhoneNumber: args.String(5), Address: domain.UserAddress{City: args.String(6), State: args.String(7), Neighborhood: args.String(8), Street: args.String(9), Number: args.String(10), ZipCode: args.String(11)}}, args.Error(12)
}
