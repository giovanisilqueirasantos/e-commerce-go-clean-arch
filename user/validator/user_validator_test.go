package validator

import (
	"context"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/assert"
)

func TestValidateUserDataCanNotBeEmpty(t *testing.T) {
	isValid, message := NewUserValidator().validate(context.Background(), &domain.User{})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com"})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "first name"})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "first name", LastName: "last name"})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "first name", LastName: "last name", PhoneNumber: "phone number"})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "first name", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "first name", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "first name", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "first name", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "first name", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state", Street: "street"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "first name", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state", Street: "street", ZipCode: "zipcode"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)
}
