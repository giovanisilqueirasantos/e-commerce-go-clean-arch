package service

import (
	"context"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	token, err := NewTokenService().Sign(context.Background(), domain.TokenInfo{Info: "token info"}, 10)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestIsValidTokenInvalid(t *testing.T) {
	ts := NewTokenService()

	token, _ := ts.Sign(context.Background(), domain.TokenInfo{Info: "token info"}, 10)

	isValid, err := ts.IsValid(context.Background(), token+"invalid string")

	assert.Error(t, err)
	assert.False(t, bool(isValid))
}

func TestIsValid(t *testing.T) {
	ts := NewTokenService()

	token, _ := ts.Sign(context.Background(), domain.TokenInfo{Info: "token info"}, 10)

	isValid, err := ts.IsValid(context.Background(), token)

	assert.NoError(t, err)
	assert.True(t, bool(isValid))
}
