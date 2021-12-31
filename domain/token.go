package domain

import (
	"context"
)

type Token string

type TokenInfo struct {
	Info string
}

type TokenService interface {
	Sign(ctx context.Context, info TokenInfo, expirationInMinutes int64) (Token, error)
	IsValid(ctx context.Context, token Token) (IsValid, error)
}
