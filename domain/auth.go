package domain

import "context"

type Auth struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type AuthUseCase interface {
	Login(ctx context.Context, a *Auth) (Token, error)
	SignUp(ctx context.Context, a *Auth, u *User) (OK, error)
	ForgotPassword(ctx context.Context, a *Auth) (OK, error)
}

type AuthValidator interface {
	Validate(ctx context.Context, a *Auth) (IsValid, Message, error)
}
