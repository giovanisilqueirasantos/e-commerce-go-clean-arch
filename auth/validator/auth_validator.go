package validator

import (
	"context"
	"net/mail"
	"unicode"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type authValidator struct{}

func NewAuthValidator() *authValidator {
	return &authValidator{}
}

func (av *authValidator) Validate(ctx context.Context, a *domain.Auth) (domain.IsValid, domain.Message) {
	if a.Login == "" || a.Password == "" {
		return false, "login or password can not be empty"
	}

	if _, err := mail.ParseAddress(a.Login); err == nil {
		return false, "login is not a valid email"
	}

	if len(a.Password) < 3 {
		return false, "password need to have at least 3 characters"
	}

	hasUpper := false

	for _, ch := range a.Password {
		if unicode.IsUpper(ch) {
			hasUpper = true
		}
	}

	if !hasUpper {
		return false, "password need to have a uppercase character"
	}

	hasNumber := false

	for _, ch := range a.Password {
		if unicode.IsNumber(ch) {
			hasNumber = true
		}
	}

	if !hasNumber {
		return false, "password need to have a number"
	}

	hasSymbol := false

	for _, ch := range a.Password {
		if unicode.IsSymbol(ch) {
			hasSymbol = true
		}
	}

	if !hasSymbol {
		return false, "password need to have a symbol character"
	}

	return true, ""
}
