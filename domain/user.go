package domain

import "context"

type User struct {
	ID          int64
	UUID        string      `json:"uuid"`
	Email       string      `json:"email"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	PhoneNumber string      `json:"phoneNumber"`
	Address     UserAddress `json:"address"`
}

type UserAddress struct {
	City         string `json:"city"`
	State        string `json:"state"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	ZipCode      string `json:"zipcode"`
}

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type UserValidator interface {
	Validate(ctx context.Context, u *User) (IsValid, Message)
}
