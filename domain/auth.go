package domain

import "context"

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ForgotPassReset struct {
	Code        string `json:"code"`
	NewPassword string `json:"newPassword"`
}

type AuthUseCase interface {
	Login(ctx context.Context, a *Auth) (Token, error)
	SignUp(ctx context.Context, a *Auth, u *User) error
	ForgotPassCode(ctx context.Context, login string) error
	ForgotPassReset(ctx context.Context, fpr *ForgotPassReset) (Token, error)
}

type AuthService interface {
	EncodePass(ctx context.Context, pass string) string
	PassIsEqualHashedPass(ctx context.Context, pass string, hashedPass string) bool
}

type AuthRepository interface {
	GetByLogin(ctx context.Context, login string) (*Auth, error)
}

type AuthValidator interface {
	Validate(ctx context.Context, a *Auth) (IsValid, Message, error)
	ValidateLogin(ctx context.Context, login string) (IsValid, Message, error)
}

type ForgotPassResetValidator interface {
	Validate(ctx context.Context, fpr *ForgotPassReset) (IsValid, Message, error)
}
