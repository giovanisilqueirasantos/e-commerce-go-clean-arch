package mocks

import (
	"context"

	"github.com/skeey/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockAuthValidator struct {
	mock.Mock
}

func (mav *MockAuthValidator) Validate(ctx context.Context, a *domain.Auth) (domain.IsValid, domain.Message, error) {
	args := mav.Called(ctx, a)
	return domain.IsValid(args.Bool(0)), domain.Message(args.String(1)), args.Error(2)
}
