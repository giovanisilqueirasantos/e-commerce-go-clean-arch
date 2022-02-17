package domain

import "context"

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthUseCase interface {
	Login(ctx context.Context, a *Auth) (Token, error)
	SignUp(ctx context.Context, a *Auth, u *User) (Token, error)
	ForgotPassCode(ctx context.Context, login string) error
	ForgotPassReset(ctx context.Context, code *Code, newPass string) (Token, error)
}

type AuthService interface {
	EncodePass(ctx context.Context, pass string) string
	PassIsEqualHashedPass(ctx context.Context, pass string, hashedPass string) bool
}

type AuthRepository interface {
	GetByLogin(ctx context.Context, login string) (*Auth, error)
	StoreWithUser(ctx context.Context, a *Auth, u *User) error
	Update(ctx context.Context, a *Auth) error
}

type AuthValidator interface {
	Validate(ctx context.Context, a *Auth) (IsValid, Message)
	ValidateLogin(ctx context.Context, login string) (IsValid, Message)
}
