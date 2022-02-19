package validator

import (
	"context"
	"net/mail"
	"strings"
	"unicode"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type userValidator struct{}

func NewUserValidator() *userValidator {
	return &userValidator{}
}

func (uv *userValidator) validate(ctx context.Context, u *domain.User) (domain.IsValid, domain.Message) {
	if u.Email == "" {
		return false, "user's email can not be empty"
	}

	if u.FirstName == "" {
		return false, "user's first name can not be empty"
	}

	if u.LastName == "" {
		return false, "user's last name can not be empty"
	}

	if u.PhoneNumber == "" {
		return false, "user's phone number can not be empty"
	}

	if u.Address == "" {
		return false, "user's address can not be empty"
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return false, "user's email is not a valid email"
	}

	firstNameIsAllLetter := true

	for _, ch := range u.FirstName {
		if !unicode.IsLetter(ch) {
			firstNameIsAllLetter = false
		}
	}

	if !firstNameIsAllLetter {
		return false, "user's first name must contain only letters"
	}

	lastNameIsAllLetter := true
	lastNameWords := strings.Fields(u.LastName)

	for _, word := range lastNameWords {
		for _, ch := range word {
			if !unicode.IsLetter(ch) {
				lastNameIsAllLetter = false
			}
		}
	}

	if !lastNameIsAllLetter {
		return false, "user's last name must contain only letters and spaces"
	}

	return true, ""
}
