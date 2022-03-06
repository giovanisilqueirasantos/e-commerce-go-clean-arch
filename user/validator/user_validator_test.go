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

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname"})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name"})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name", PhoneNumber: "phone number"})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)

	isValid, message = NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state", Street: "street"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)
}

func TestValidateEmailInvalid(t *testing.T) {
	isValid, message := NewUserValidator().validate(context.Background(), &domain.User{Email: "email@emailcom", FirstName: "firstname", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state", Street: "street", ZipCode: "zipcode"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)
}

func TestValidateFirstNameContainOnlyLetters(t *testing.T) {
	isValid, message := NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname123", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state", Street: "street", ZipCode: "zipcode"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)
}

func TestValidateLastNameContainOnlyLetters(t *testing.T) {
	isValid, message := NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name 123", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state", Street: "street", ZipCode: "zipcode"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)
}

func TestValidatePhoneNumberInvalid(t *testing.T) {
	isValid, message := NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name", PhoneNumber: "phone number", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state", Street: "street", ZipCode: "zipcode"}})

	assert.False(t, bool(isValid))
	assert.NotEmpty(t, message)
}

func TestValidateSuccess(t *testing.T) {
	isValid, _ := NewUserValidator().validate(context.Background(), &domain.User{Email: "email@email.com", FirstName: "firstname", LastName: "last name", PhoneNumber: "(11) 12345-1234", Address: domain.UserAddress{City: "city", Neighborhood: "neighborhood", Number: "number", State: "state", Street: "street", ZipCode: "zipcode"}})

	assert.True(t, bool(isValid))
}
