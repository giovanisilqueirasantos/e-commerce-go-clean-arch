package domain

import "context"

type Code struct {
	Value      string
	Identifier string
}

type CodeService interface {
	GenerateNewCode(ctx context.Context, identifier string, length int8, number bool, symbol bool) (*Code, error)
	GenerateNewCodeFake(ctx context.Context)
	ValidateCode(ctx context.Context, c *Code) (IsValid, error)
}

type CodeRepository interface {
	GetByValue(ctx context.Context, value string) (*Code, error)
	DeleteByValue(ctx context.Context, value string) error
}
