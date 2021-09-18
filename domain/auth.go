package domain

import "context"

type Auth struct {
	Login           string `json:"login"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type AuthUseCase interface {
	Login(ctx context.Context, a *Auth) (Token, error)
	SignUp(ctx context.Context, a *Auth, u *User) error
	ForgotPassword(ctx context.Context, a *Auth) error
}

type AuthValidator interface {
	Validate(ctx context.Context, a *Auth) (IsValid, Message, error)
}
