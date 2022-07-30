package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginCheckLoginExistsError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(nil, errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, nil, nil, nil, mockAuthRepo, nil)

	_, err := authUseCase.Login(context.Background(), &mockAuth)

	assert.Error(t, err)
}

func TestLoginCheckLoginExists(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(nil, nil)

	authUseCase := NewAuthUseCase(nil, nil, nil, nil, mockAuthRepo, nil)

	_, err := authUseCase.Login(context.Background(), &mockAuth)

	assert.Error(t, err)
}

func TestLoginPassIsEqualHashedPassError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockAuthService := new(mocks.MockAuthService)

	var mockAuth domain.Auth
	mockAuth.UUID = "uuid"
	mockAuth.Login = "valid login"
	mockAuth.Password = "invalid password"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(1, "uuid", "valid login", "valid password", nil)

	mockAuthService.On("PassIsEqualHashedPass", mock.Anything, mockAuth.Password, "valid password").Return(false)

	authUseCase := NewAuthUseCase(mockAuthService, nil, nil, nil, mockAuthRepo, nil)

	_, err := authUseCase.Login(context.Background(), &mockAuth)

	assert.Error(t, err)
}

func TestLoginSignTokenError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockAuthService := new(mocks.MockAuthService)
	mockTokenService := new(mocks.MockTokenService)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(1, "uuid", mockAuth.Login, mockAuth.Password, nil)

	mockAuthService.On("PassIsEqualHashedPass", mock.Anything, mockAuth.Password, mockAuth.Password).Return(true)

	var thirtyDaysInMinutes int64 = 43200

	tokenInfo := domain.TokenInfo{Info: mockAuth.Login}

	mockTokenService.On("Sign", mock.Anything, tokenInfo, thirtyDaysInMinutes).Return("", errors.New("error message"))

	authUseCase := NewAuthUseCase(mockAuthService, mockTokenService, nil, nil, mockAuthRepo, nil)

	_, err := authUseCase.Login(context.Background(), &mockAuth)

	assert.Error(t, err)
}

func TestLoginSuccess(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockAuthService := new(mocks.MockAuthService)
	mockTokenService := new(mocks.MockTokenService)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(1, "uuid", mockAuth.Login, mockAuth.Password, nil)

	mockAuthService.On("PassIsEqualHashedPass", mock.Anything, mockAuth.Password, mockAuth.Password).Return(true)

	var thirtyDaysInMinutes int64 = 43200

	tokenInfo := domain.TokenInfo{Info: mockAuth.Login}

	mockTokenService.On("Sign", mock.Anything, tokenInfo, thirtyDaysInMinutes).Return("valid token", nil)

	authUseCase := NewAuthUseCase(mockAuthService, mockTokenService, nil, nil, mockAuthRepo, nil)

	token, err := authUseCase.Login(context.Background(), &mockAuth)

	assert.Nil(t, err)
	assert.Equal(t, token, domain.Token("valid token"))
}

func TestSignUpCheckLoginExistsError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(nil, errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, nil, nil, nil, mockAuthRepo, nil)

	_, err := authUseCase.SignUp(context.Background(), &mockAuth, nil)

	assert.Error(t, err)
}

func TestSignUpLoginAlreadyExists(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(1, "uuid", "valid login", "valid password", nil)

	authUseCase := NewAuthUseCase(nil, nil, nil, nil, mockAuthRepo, nil)

	_, err := authUseCase.SignUp(context.Background(), &mockAuth, nil)

	assert.Error(t, err)
}

func TestSignUpCheckUserExistsError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"

	var mockUser domain.User
	mockUser.Email = "valid email"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(nil, nil)

	mockUserRepo.On("GetByEmail", mock.Anything, mockUser.Email).Return(nil, errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, nil, nil, nil, mockAuthRepo, mockUserRepo)

	_, err := authUseCase.SignUp(context.Background(), &mockAuth, &mockUser)

	assert.Error(t, err)
}

func TestSignUpCheckUserExists(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"

	var mockUser domain.User
	mockUser.Email = "valid email"

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(nil, nil)

	mockUserRepo.On("GetByEmail", mock.Anything, mockUser.Email).Return(1, "uuid", "user email", "user first name", "user last name", "user phone number", "user address city", "user address state", "user address neighborhood", "user address street", "user address number", "user address zipcode", nil)

	authUseCase := NewAuthUseCase(nil, nil, nil, nil, mockAuthRepo, mockUserRepo)

	_, err := authUseCase.SignUp(context.Background(), &mockAuth, &mockUser)

	assert.Error(t, err)
}

func TestSignUpStoreUserError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockAuthService := new(mocks.MockAuthService)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"

	var mockUser domain.User
	mockUser.Email = "user email"

	mockAuthService.On("EncodePass", mock.Anything, mockAuth.Password).Return("hashed password")

	mockUserRepo.On("GetByEmail", mock.Anything, mockUser.Email).Return(nil, nil)

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(1, "uuid", "valid login", "valid password", nil)
	mockAuthRepo.On("StoreWithUser", mock.Anything, &domain.Auth{Login: mockAuth.Login, Password: "hashed password"}, &mockUser).Return(errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, nil, nil, nil, mockAuthRepo, mockUserRepo)

	_, err := authUseCase.SignUp(context.Background(), &mockAuth, &mockUser)

	assert.Error(t, err)
}

func TestSignUpSignTokenError(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockTokenService := new(mocks.MockTokenService)
	mockAuthService := new(mocks.MockAuthService)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"

	var mockUser domain.User
	mockUser.Email = "user email"

	mockAuthService.On("EncodePass", mock.Anything, mockAuth.Password).Return("hashed password")

	mockUserRepo.On("GetByEmail", mock.Anything, mockUser.Email).Return(nil, nil)

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(1, "uuid", "valid login", "valid password", nil)
	mockAuthRepo.On("StoreWithUser", mock.Anything, &domain.Auth{Login: mockAuth.Login, Password: "hashed password"}, &mockUser).Return(nil)

	var thirtyDaysInMinutes int64 = 43200

	tokenInfo := domain.TokenInfo{Info: mockAuth.Login}

	mockTokenService.On("Sign", mock.Anything, tokenInfo, thirtyDaysInMinutes).Return("", errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, mockTokenService, nil, nil, mockAuthRepo, mockUserRepo)

	_, err := authUseCase.SignUp(context.Background(), &mockAuth, &mockUser)

	assert.Error(t, err)
}

func TestSignUpSuccess(t *testing.T) {
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockTokenService := new(mocks.MockTokenService)
	mockAuthService := new(mocks.MockAuthService)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"

	var mockUser domain.User
	mockUser.Email = "user email"

	mockAuthService.On("EncodePass", mock.Anything, mockAuth.Password).Return("hashed password")

	mockUserRepo.On("GetByEmail", mock.Anything, mockUser.Email).Return(nil, nil)

	mockAuthRepo.On("GetByLogin", mock.Anything, mockAuth.Login).Return(nil, nil)
	mockAuthRepo.On("StoreWithUser", mock.Anything, &domain.Auth{Login: mockAuth.Login, Password: "hashed password"}, &mockUser).Return(nil)

	var thirtyDaysInMinutes int64 = 43200

	tokenInfo := domain.TokenInfo{Info: mockAuth.Login}

	mockTokenService.On("Sign", mock.Anything, tokenInfo, thirtyDaysInMinutes).Return("valid token", nil)

	authUseCase := NewAuthUseCase(mockAuthService, mockTokenService, nil, nil, mockAuthRepo, mockUserRepo)

	token, err := authUseCase.SignUp(context.Background(), &mockAuth, &mockUser)

	assert.Nil(t, err)
	assert.Equal(t, token, domain.Token("valid token"))
}

func TestForgotPassCodeGetUserByLoginError(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockCodeService := new(mocks.MockCodeService)
	mockMessageService := new(mocks.MockMessageService)

	mockLogin := "valid login"

	mockUserRepo.On("GetByEmail", mock.Anything, mockLogin).Return(nil, errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, nil, mockCodeService, mockMessageService, nil, mockUserRepo)

	err := authUseCase.ForgotPassCode(context.Background(), mockLogin)

	assert.Error(t, err)
}

func TestForgotPassCodeNoUserFound(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockCodeService := new(mocks.MockCodeService)
	mockMessageService := new(mocks.MockMessageService)

	mockLogin := "valid login"

	mockUserRepo.On("GetByEmail", mock.Anything, mockLogin).Return(nil, nil)

	authUseCase := NewAuthUseCase(nil, nil, mockCodeService, mockMessageService, nil, mockUserRepo)

	err := authUseCase.ForgotPassCode(context.Background(), mockLogin)

	assert.Error(t, err)
}

func TestForgotPassCodeSendMessageError(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockCodeService := new(mocks.MockCodeService)
	mockMessageService := new(mocks.MockMessageService)

	mockLogin := "valid login"

	mockUserRepo.On("GetByEmail", mock.Anything, mockLogin).Return(1, "uuid", "user email", "user first name", "user last name", "user phone number", "user address city", "user address state", "user address neighborhood", "user address street", "user address number", "user address zipcode", nil)

	mockCodeService.On("GenerateNewCode", mock.Anything, mockLogin, int8(6), true, false).Return("generated code", mockLogin, nil)

	var messageConf domain.MessageConfig

	messageConf.Medium = "phone"
	messageConf.To = "user phone number"
	messageConf.Message = "O código para recuperar sua senha é generated code"

	mockMessageService.On("SendMessage", mock.Anything, &messageConf).Return(errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, nil, mockCodeService, mockMessageService, nil, mockUserRepo)

	err := authUseCase.ForgotPassCode(context.Background(), mockLogin)

	assert.Error(t, err)
}

func TestForgotPassCodeSuccess(t *testing.T) {
	mockUserRepo := new(mocks.MockUserRepository)
	mockCodeService := new(mocks.MockCodeService)
	mockMessageService := new(mocks.MockMessageService)

	mockLogin := "valid login"

	mockUserRepo.On("GetByEmail", mock.Anything, mockLogin).Return(1, "uuid", "user email", "user first name", "user last name", "user phone number", "user address city", "user address state", "user address neighborhood", "user address street", "user address number", "user address zipcode", nil)

	mockCodeService.On("GenerateNewCode", mock.Anything, mockLogin, int8(6), true, false).Return("generated code", mockLogin, nil)

	var messageConf domain.MessageConfig

	messageConf.Medium = "phone"
	messageConf.To = "user phone number"
	messageConf.Message = "O código para recuperar sua senha é generated code"

	mockMessageService.On("SendMessage", mock.Anything, &messageConf).Return(nil)

	authUseCase := NewAuthUseCase(nil, nil, mockCodeService, mockMessageService, nil, mockUserRepo)

	err := authUseCase.ForgotPassCode(context.Background(), mockLogin)

	assert.Nil(t, err)
}

func TestForgotPassResetValidateCodeError(t *testing.T) {
	mockCodeService := new(mocks.MockCodeService)

	var mockCode domain.Code

	mockCode.Identifier = "identifier"
	mockCode.Value = "Value"

	mockCodeService.On("ValidateCode", mock.Anything, &mockCode).Return(false, errors.New("error message"))

	authUseCase := NewAuthUseCase(nil, nil, mockCodeService, nil, nil, nil)

	_, err := authUseCase.ForgotPassReset(context.Background(), &mockCode, "new pass")

	assert.Error(t, err)
}

func TestForgotPassResetCodeInvalid(t *testing.T) {
	mockCodeService := new(mocks.MockCodeService)

	var mockCode domain.Code

	mockCode.Identifier = "identifier"
	mockCode.Value = "Value"

	mockCodeService.On("ValidateCode", mock.Anything, &mockCode).Return(false, nil)

	authUseCase := NewAuthUseCase(nil, nil, mockCodeService, nil, nil, nil)

	_, err := authUseCase.ForgotPassReset(context.Background(), &mockCode, "new pass")

	assert.Error(t, err)
}

func TestForgotPassResetGetAuthByLoginError(t *testing.T) {
	mockCodeService := new(mocks.MockCodeService)
	mockAuthService := new(mocks.MockAuthService)
	mockAuthRepo := new(mocks.MockAuthRepository)

	var mockCode domain.Code

	mockNewPass := "new pass"
	mockEncodedNewPass := "encoded new pass"

	mockCode.Identifier = "identifier"
	mockCode.Value = "Value"

	mockCodeService.On("ValidateCode", mock.Anything, &mockCode).Return(true, nil)

	mockAuthService.On("EncodePass", mock.Anything, mockNewPass).Return(mockEncodedNewPass)

	mockAuthRepo.On("GetByLogin", mock.Anything, mockCode.Identifier).Return(nil, errors.New("error message"))

	authUseCase := NewAuthUseCase(mockAuthService, nil, mockCodeService, nil, mockAuthRepo, nil)

	_, err := authUseCase.ForgotPassReset(context.Background(), &mockCode, mockNewPass)

	assert.Error(t, err)
}

func TestForgotPassResetUpdateAuthError(t *testing.T) {
	mockCodeService := new(mocks.MockCodeService)
	mockAuthService := new(mocks.MockAuthService)
	mockAuthRepo := new(mocks.MockAuthRepository)

	var mockCode domain.Code

	mockNewPass := "new pass"
	mockEncodedNewPass := "encoded new pass"

	mockCode.Identifier = "identifier"
	mockCode.Value = "Value"

	mockCodeService.On("ValidateCode", mock.Anything, &mockCode).Return(true, nil)

	mockAuthService.On("EncodePass", mock.Anything, mockNewPass).Return(mockEncodedNewPass)

	var auth domain.Auth

	auth.ID = 1
	auth.UUID = "uuid"
	auth.Login = mockCode.Identifier
	auth.Password = mockEncodedNewPass

	mockAuthRepo.On("GetByLogin", mock.Anything, auth.Login).Return(1, "uuid", auth.Login, "valid password", nil)
	mockAuthRepo.On("Update", mock.Anything, &auth).Return(errors.New("error message"))

	authUseCase := NewAuthUseCase(mockAuthService, nil, mockCodeService, nil, mockAuthRepo, nil)

	_, err := authUseCase.ForgotPassReset(context.Background(), &mockCode, mockNewPass)

	assert.Error(t, err)
}

func TestForgotPassResetSignTokenError(t *testing.T) {
	mockCodeService := new(mocks.MockCodeService)
	mockAuthService := new(mocks.MockAuthService)
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockTokenService := new(mocks.MockTokenService)

	var mockCode domain.Code

	mockNewPass := "new pass"
	mockEncodedNewPass := "encoded new pass"

	mockCode.Identifier = "identifier"
	mockCode.Value = "Value"

	mockCodeService.On("ValidateCode", mock.Anything, &mockCode).Return(true, nil)

	mockAuthService.On("EncodePass", mock.Anything, mockNewPass).Return(mockEncodedNewPass)

	var auth domain.Auth

	auth.ID = 1
	auth.UUID = "uuid"
	auth.Login = mockCode.Identifier
	auth.Password = mockEncodedNewPass

	mockAuthRepo.On("GetByLogin", mock.Anything, auth.Login).Return(1, "uuid", auth.Login, "valid password", nil)
	mockAuthRepo.On("Update", mock.Anything, &auth).Return(nil)

	var thirtyDaysInMinutes int64 = 43200

	tokenInfo := domain.TokenInfo{Info: mockCode.Identifier}

	mockTokenService.On("Sign", mock.Anything, tokenInfo, thirtyDaysInMinutes).Return("", errors.New("error message"))

	authUseCase := NewAuthUseCase(mockAuthService, mockTokenService, mockCodeService, nil, mockAuthRepo, nil)

	_, err := authUseCase.ForgotPassReset(context.Background(), &mockCode, mockNewPass)

	assert.Error(t, err)
}

func TestForgotPassResetSuccess(t *testing.T) {
	mockCodeService := new(mocks.MockCodeService)
	mockAuthService := new(mocks.MockAuthService)
	mockAuthRepo := new(mocks.MockAuthRepository)
	mockTokenService := new(mocks.MockTokenService)

	var mockCode domain.Code

	mockNewPass := "new pass"
	mockEncodedNewPass := "encoded new pass"

	mockCode.Identifier = "identifier"
	mockCode.Value = "Value"

	mockCodeService.On("ValidateCode", mock.Anything, &mockCode).Return(true, nil)

	mockAuthService.On("EncodePass", mock.Anything, mockNewPass).Return(mockEncodedNewPass)

	var auth domain.Auth

	auth.ID = 1
	auth.UUID = "uuid"
	auth.Login = mockCode.Identifier
	auth.Password = mockEncodedNewPass

	mockAuthRepo.On("GetByLogin", mock.Anything, auth.Login).Return(1, "uuid", auth.Login, "valid password", nil)
	mockAuthRepo.On("Update", mock.Anything, &auth).Return(nil)

	var thirtyDaysInMinutes int64 = 43200

	tokenInfo := domain.TokenInfo{Info: mockCode.Identifier}

	mockTokenService.On("Sign", mock.Anything, tokenInfo, thirtyDaysInMinutes).Return("valid token", nil)

	authUseCase := NewAuthUseCase(mockAuthService, mockTokenService, mockCodeService, nil, mockAuthRepo, nil)

	token, err := authUseCase.ForgotPassReset(context.Background(), &mockCode, mockNewPass)

	assert.Nil(t, err)
	assert.Equal(t, token, domain.Token("valid token"))
}
