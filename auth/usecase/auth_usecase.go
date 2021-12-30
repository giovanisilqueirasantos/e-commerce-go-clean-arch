package usecase

import (
	"context"
	"fmt"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type authUseCase struct {
	authRepo domain.AuthRepository
}

func NewAuthUseCase(ar domain.AuthRepository) domain.AuthUseCase {
	return &authUseCase{
		authRepo: ar,
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

	return "token", nil
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
