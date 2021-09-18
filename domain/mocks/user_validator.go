package mocks

import (
	"context"

	"github.com/skeey/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockUserValidator struct {
	mock.Mock
}

func (muv *MockUserValidator) Validate(ctx context.Context, u *domain.User) (domain.IsValid, domain.Message, error) {
	args := muv.Called(ctx, u)
	return domain.IsValid(args.Bool(0)), domain.Message(args.String(1)), args.Error(2)
}
