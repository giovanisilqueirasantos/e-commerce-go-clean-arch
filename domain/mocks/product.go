package mocks

import (
	"context"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type MockProductUsecase struct {
	mock.Mock
}

func (mpu *MockProductUsecase) Get(ctx context.Context, uuid string) (*domain.Product, error) {
	args := mpu.Called(ctx, uuid)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return &domain.Product{ID: int64(args.Int(0)), UUID: args.String(1), Rate: float32(args.Int(2)), Pictures: []string{args.String(3)}, Name: args.String(4), Detail: args.String(5), Favorite: args.Bool(6), Attributes: []domain.Attribute{domain.Attribute{Label: args.String(7), Values: []string{args.String(8)}}}}, args.Error(9)
}
