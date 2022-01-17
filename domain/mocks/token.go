package mocks

import (
	"context"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockTokenService struct {
	mock.Mock
}

func (mts *MockTokenService) Sign(ctx context.Context, info domain.TokenInfo, expirationInMinutes int64) (domain.Token, error) {
	args := mts.Called(ctx, info, expirationInMinutes)
	return domain.Token(args.String(0)), args.Error(1)
}

func (mts *MockTokenService) IsValid(ctx context.Context, token domain.Token) (domain.IsValid, error) {
	args := mts.Called(ctx, token)
	return domain.IsValid(args.Bool(0)), args.Error(1)
}
