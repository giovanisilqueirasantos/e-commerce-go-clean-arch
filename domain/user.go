package domain

import "context"

type User struct {
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneName"`
	Address     string `json:"addressName"`
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
	Update(ctx context.Context, u *User) error
}
