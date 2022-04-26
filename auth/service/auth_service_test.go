package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodePass(t *testing.T) {
	authService := NewAuthService()
	encodedPass := authService.EncodePass(context.Background(), "password")
	assert.NotEmpty(t, encodedPass)
	assert.NotEqual(t, "password", encodedPass)
}

func TestPassIsEqualHashedPass(t *testing.T) {
	authService := NewAuthService()
	encodedPass := authService.EncodePass(context.Background(), "password")
	isEncoded := authService.PassIsEqualHashedPass(context.Background(), "password", encodedPass)
	assert.True(t, isEncoded)
}
