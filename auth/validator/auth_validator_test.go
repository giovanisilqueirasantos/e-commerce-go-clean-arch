package validator

import (
	"context"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyLoginOrPassword(t *testing.T) {
	isLoginValid, isLoginValidMessage := NewAuthValidator().Validate(context.Background(), &domain.Auth{Login: "", Password: "valid pass"})

	assert.False(t, bool(isLoginValid))
	assert.NotEmpty(t, isLoginValidMessage)

	isPassValid, isPassValidMessage := NewAuthValidator().Validate(context.Background(), &domain.Auth{Login: "valid login", Password: ""})

	assert.False(t, bool(isPassValid))
	assert.NotEmpty(t, isPassValidMessage)
}

func TestValidateEmailInvalid(t *testing.T) {
	isLoginValid, isLoginValidMessage := NewAuthValidator().Validate(context.Background(), &domain.Auth{Login: "invalid login", Password: "valid pass"})

	assert.False(t, bool(isLoginValid))
	assert.NotEmpty(t, isLoginValidMessage)
}

func TestValidatePasswordWith2Char(t *testing.T) {
	isPassValid, isPassValidMessage := NewAuthValidator().Validate(context.Background(), &domain.Auth{Login: "login@email.com", Password: "pa"})

	assert.False(t, bool(isPassValid))
	assert.NotEmpty(t, isPassValidMessage)
}

func TestValidatePasswordWithNoUpper(t *testing.T) {
	isPassValid, isPassValidMessage := NewAuthValidator().Validate(context.Background(), &domain.Auth{Login: "login@email.com", Password: "pass"})

	assert.False(t, bool(isPassValid))
	assert.NotEmpty(t, isPassValidMessage)
}

func TestValidatePasswordWithNoNumber(t *testing.T) {
	isPassValid, isPassValidMessage := NewAuthValidator().Validate(context.Background(), &domain.Auth{Login: "login@email.com", Password: "pasS"})

	assert.False(t, bool(isPassValid))
	assert.NotEmpty(t, isPassValidMessage)
}

func TestValidatePasswordWithNoSymbol(t *testing.T) {
	isPassValid, isPassValidMessage := NewAuthValidator().Validate(context.Background(), &domain.Auth{Login: "login@email.com", Password: "pasS1"})

	assert.False(t, bool(isPassValid))
	assert.NotEmpty(t, isPassValidMessage)
}

func TestValidateAuthValid(t *testing.T) {
	isAuthValid, _ := NewAuthValidator().Validate(context.Background(), &domain.Auth{Login: "login@email.com", Password: "pasS1$"})

	assert.True(t, bool(isAuthValid))
}
