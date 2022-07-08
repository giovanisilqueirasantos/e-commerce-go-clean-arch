package usecase

import (
	"context"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type productUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(pr domain.ProductRepository) domain.ProductUseCase {
	return &productUseCase{productRepo: pr}
}

func (pu *productUseCase) Get(ctx context.Context, uuid string) (*domain.Product, error) {
	return pu.productRepo.GetByUUID(ctx, uuid)
}
