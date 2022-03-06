package validator

import (
	"context"
	"net/mail"
	"regexp"
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

	if u.Address.City == "" {
		return false, "user's address city can not be empty"
	}

	if u.Address.Neighborhood == "" {
		return false, "user's address neighborhood can not be empty"
	}

	if u.Address.Number == "" {
		return false, "user's address number can not be empty"
	}

	if u.Address.State == "" {
		return false, "user's address state can not be empty"
	}

	if u.Address.Street == "" {
		return false, "user's address street can not be empty"
	}

	if u.Address.ZipCode == "" {
		return false, "user's address zipcode can not be empty"
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

	validPhone := regexp.MustCompile(`^\([0-9]{2}\) [0-9]{5}\-[0-9]{4}$`)

	if !validPhone.MatchString(u.PhoneNumber) {
		return false, "user's phone number must obey the format (11) 11111-1111"
	}

	return true, ""
}
