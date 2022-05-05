package mocks

import (
	"context"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockCodeService struct {
	mock.Mock
}

func (mcs *MockCodeService) GenerateNewCode(ctx context.Context, identifier string, length int8, number bool, symbol bool) *domain.Code {
	args := mcs.Called(ctx, identifier, length, number, symbol)
	return &domain.Code{Value: args.String(0), Identifier: args.String(1)}
}

func (mcs *MockCodeService) GenerateNewCodeFake(ctx context.Context) {}

func (mcs *MockCodeService) ValidateCode(ctx context.Context, c *domain.Code) (domain.IsValid, error) {
	args := mcs.Called(ctx, c)
	return domain.IsValid(args.Bool(0)), args.Error(1)
}

type MockCodeRepository struct {
	mock.Mock
}

func (mcr *MockCodeRepository) GetByValue(ctx context.Context, value string) (*domain.Code, error) {
	args := mcr.Called(ctx, value)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return &domain.Code{Value: args.String(0), Identifier: args.String(1)}, args.Error(2)
}

func (mcr *MockCodeRepository) DeleteByValue(ctx context.Context, value string) error {
	args := mcr.Called(ctx, value)
	return args.Error(0)
}
