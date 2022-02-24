package domain

import "context"

type User struct {
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
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByLogin(ctx context.Context, login string) (*User, error)
	Update(ctx context.Context, u *User) error
}

type UserValidator interface {
	Validate(ctx context.Context, u *User) (IsValid, Message)
}
