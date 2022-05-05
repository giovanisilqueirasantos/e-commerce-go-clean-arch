package usecase

import (
	"context"
	"fmt"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type authUseCase struct {
	authService    domain.AuthService
	tokenService   domain.TokenService
	codeService    domain.CodeService
	messageService domain.MessageService
	authRepo       domain.AuthRepository
	userRepo       domain.UserRepository
}

func NewAuthUseCase(as domain.AuthService, ts domain.TokenService, cs domain.CodeService, ms domain.MessageService, ar domain.AuthRepository, ur domain.UserRepository) domain.AuthUseCase {
	return &authUseCase{
		authService:    as,
		tokenService:   ts,
		codeService:    cs,
		messageService: ms,
		authRepo:       ar,
		userRepo:       ur,
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

func (au *authUseCase) SignUp(ctx context.Context, a *domain.Auth, u *domain.User) (domain.Token, error) {
	auth, errAuth := au.authRepo.GetByLogin(ctx, a.Login)

	if errAuth != nil {
		return "", errAuth
	}

	if auth != nil {
		return "", fmt.Errorf("auth with login %s already exists", a.Login)
	}

	user, errUser := au.userRepo.GetByEmail(ctx, u.Email)

	if errUser != nil {
		return "", errUser
	}

	if user != nil {
		return "", fmt.Errorf("user with email %s already exists", u.Email)
	}

	a.Password = au.authService.EncodePass(ctx, a.Password)

	if storeErr := au.authRepo.StoreWithUser(ctx, a, u); storeErr != nil {
		return "", storeErr
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

func (au *authUseCase) ForgotPassCode(ctx context.Context, login string) error {
	user, errUser := au.userRepo.GetByEmail(ctx, login)

	if errUser != nil {
		au.codeService.GenerateNewCodeFake(ctx)
		au.messageService.SendMessageFake(ctx)
		return errUser
	}

	if user == nil {
		au.codeService.GenerateNewCodeFake(ctx)
		au.messageService.SendMessageFake(ctx)
		return fmt.Errorf("user with login %s not found", login)
	}

	code := au.codeService.GenerateNewCode(ctx, login, 6, true, false)

	message := fmt.Sprintf("O código para recuperar sua senha é %s", code.Value)

	var messageConf domain.MessageConfig

	messageConf.Medium = "phone"
	messageConf.To = user.PhoneNumber
	messageConf.Message = message

	if errMessage := au.messageService.SendMessage(ctx, &messageConf); errMessage != nil {
		return errMessage
	}

	return nil
}

func (au *authUseCase) ForgotPassReset(ctx context.Context, code *domain.Code, newPass string) (domain.Token, error) {
	codeIsValid, errCode := au.codeService.ValidateCode(ctx, code)

	if errCode != nil {
		return "", errCode
	}

	if !codeIsValid {
		return "", fmt.Errorf("code %s with identifier %s is not valid", code.Value, code.Identifier)
	}

	auth, err := au.authRepo.GetByLogin(ctx, code.Identifier)

	if err != nil {
		return "", err
	}

	auth.Password = au.authService.EncodePass(ctx, newPass)

	if err = au.authRepo.Update(ctx, auth); err != nil {
		return "", err
	}

	var tokenInfo domain.TokenInfo

	tokenInfo.Info = code.Identifier

	var thirtyDaysInMinutes int64 = 43200

	token, tokenErr := au.tokenService.Sign(ctx, tokenInfo, thirtyDaysInMinutes)

	if tokenErr != nil {
		return "", tokenErr
	}

	return token, nil
}
