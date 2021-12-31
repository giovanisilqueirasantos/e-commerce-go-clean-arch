package usecase

import (
	"context"
	"fmt"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type authUseCase struct {
	authService  domain.AuthService
	tokenService domain.TokenService
	authRepo     domain.AuthRepository
}

func NewAuthUseCase(as domain.AuthService, ts domain.TokenService, ar domain.AuthRepository) domain.AuthUseCase {
	return &authUseCase{
		authService:  as,
		tokenService: ts,
		authRepo:     ar,
	}
}

func (au *authUseCase) Login(ctx context.Context, a *domain.Auth) (domain.Token, error) {
	auth, errAuth := au.authRepo.GetByLogin(ctx, a.Login)

	if errAuth != nil {
		return "", errAuth
	}

	if auth == nil {
		return "", fmt.Errorf("auth with login %s not found", a.Login)
	}

	if !au.authService.PassIsEqualHashedPass(ctx, a.Password, auth.Password) {
		return "", fmt.Errorf("wrong password for login %s", a.Login)
	}

	var tokenInfo domain.TokenInfo

	tokenInfo.Info = a.Login

	var thirtyDaysInMinutes int64 = 43200

	token, tokenErr := au.tokenService.Sign(ctx, tokenInfo, thirtyDaysInMinutes)

	if tokenErr != nil {
		return "", tokenErr
	}

	return token, nil
}

func (au *authUseCase) SignUp(ctx context.Context, a *domain.Auth, u *domain.User) error {
	return nil
}

func (au *authUseCase) ForgotPassCode(ctx context.Context, login string) error {
	return nil
}

func (au *authUseCase) ForgotPassReset(ctx context.Context, fpr *domain.ForgotPassReset) (domain.Token, error) {
	return "token", nil
}
