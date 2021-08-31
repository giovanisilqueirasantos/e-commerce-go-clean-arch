package domain

import "context"

type Auth struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type token string
type ok bool
type isValid bool
type message string

type AuthUseCase interface {
	Login(ctx context.Context, a *Auth) (token, error)
	SignUp(ctx context.Context, a *Auth, u *User) (ok, error)
	ForgotPassword(ctx context.Context, a *Auth) (ok, error)
}

type AuthValidator interface {
	Validate(ctx context.Context, a *Auth) (isValid, message, error)
}
