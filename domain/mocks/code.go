package mocks

import (
	"context"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockCodeService struct {
	mock.Mock
}

func (mcs *MockCodeService) GenerateNewCode(ctx context.Context, identifier string, length int8, number bool, symbol bool) (*domain.Code, error) {
	args := mcs.Called(ctx, identifier, length, number, symbol)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return &domain.Code{Value: args.String(0), Identifier: args.String(1)}, args.Error(2)
}

func (mcs *MockCodeService) GenerateNewCodeFake(ctx context.Context) {}

type MockCodeRepository struct {
	mock.Mock
}

func (mcr *MockAuthRepository) GetByValue(ctx context.Context, value string) (*domain.Code, error) {
	args := mcr.Called(ctx, value)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return &domain.Code{Value: args.String(0), Identifier: args.String(1)}, args.Error(2)
}

func (mcr *MockAuthRepository) DeleteByValue(ctx context.Context, value string) error {
	args := mcr.Called(ctx, value)
	return args.Error(0)
}
