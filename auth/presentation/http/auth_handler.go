package http

import (
	"log"
	"net/http"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	AuthUseCase   domain.AuthUseCase
	AuthValidator domain.AuthValidator
	UserValidator domain.UserValidator
}

func NewAuthHandler(e *echo.Echo, auc domain.AuthUseCase, av domain.AuthValidator, uv domain.UserValidator) *authHandler {
	handler := &authHandler{
		AuthUseCase:   auc,
		AuthValidator: av,
		UserValidator: uv,
	}
	e.POST("/login", handler.Login)
	e.POST("/signup", handler.SignUp)
	e.POST("/forgotpass/code", handler.ForgotPassCode)
	e.POST("/forgotpass/reset", handler.ForgotPassReset)

	return handler
}

func (ah *authHandler) Login(c echo.Context) error {
	var auth domain.Auth

	err := c.Bind(&auth)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to interpret the submitted information")
	}

	ctx := c.Request().Context()

	isValid, message, errValid := ah.AuthValidator.Validate(ctx, &auth)

	if errValid != nil {
		log.Printf("Error validating Auth: %s", errValid.Error())
		return c.JSON(http.StatusInternalServerError, "failed to login")
	}

	if !isValid {
		return c.JSON(http.StatusBadRequest, message)
	}

	token, errToken := ah.AuthUseCase.Login(ctx, &auth)

	if errToken != nil {
		log.Printf("Error trying to generate token for Login: %s", errToken.Error())
		return c.JSON(http.StatusInternalServerError, "failed to login")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": string(token)})
}

func (ah *authHandler) SignUp(c echo.Context) error {
	var authWithUser struct {
		domain.Auth
		domain.User
	}

	err := c.Bind(&authWithUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to interpret the submitted information")
	}

	ctx := c.Request().Context()

	auth := domain.Auth{
		Login:    authWithUser.Login,
		Password: authWithUser.Password,
	}

	isValidAuth, messageAuth, errValidAuth := ah.AuthValidator.Validate(ctx, &auth)

	if errValidAuth != nil {
		log.Printf("Error validating Auth: %s", errValidAuth.Error())
		return c.JSON(http.StatusInternalServerError, "failed to sign up")
	}

	if !isValidAuth {
		return c.JSON(http.StatusBadRequest, messageAuth)
	}

	user := domain.User{
		Email:       authWithUser.Email,
		FirstName:   authWithUser.FirstName,
		LastName:    authWithUser.LastName,
		PhoneNumber: authWithUser.PhoneNumber,
		Address:     authWithUser.Address,
	}

	isValidUser, messageUser, errValidUser := ah.UserValidator.Validate(ctx, &user)

	if errValidUser != nil {
		log.Printf("Error validating User: %s", errValidUser.Error())
		return c.JSON(http.StatusInternalServerError, "failed to sign up")
	}

	if !isValidUser {
		return c.JSON(http.StatusBadRequest, messageUser)
	}

	token, tokenErr := ah.AuthUseCase.SignUp(ctx, &authWithUser.Auth, &authWithUser.User)

	if tokenErr != nil {
		log.Printf("Error trying to sign up: %s", tokenErr.Error())
		return c.JSON(http.StatusInternalServerError, "failed to sign up")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": string(token)})
}

func (ah *authHandler) ForgotPassCode(c echo.Context) error {
	var forgotPassReq struct {
		Login string `json:"login"`
	}

	err := c.Bind(&forgotPassReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to interpret the submitted information")
	}

	ctx := c.Request().Context()

	isValidLogin, messageLogin, errValidLogin := ah.AuthValidator.ValidateLogin(ctx, forgotPassReq.Login)

	if errValidLogin != nil {
		log.Printf("Error validating login: %s", errValidLogin.Error())
		return c.JSON(http.StatusInternalServerError, "failed to send forgot password code")
	}

	if !isValidLogin {
		return c.JSON(http.StatusBadRequest, messageLogin)
	}

	err = ah.AuthUseCase.ForgotPassCode(ctx, forgotPassReq.Login)

	if err != nil {
		log.Printf("Error trying to send forgot password code: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, "failed to send forgot password code")
	}

	return c.String(http.StatusOK, "")
}

func (ah *authHandler) ForgotPassReset(c echo.Context) error {
	var forgotPassResetReq struct {
		Login   string `json:"login"`
		Code    string `json:"code"`
		NewPass string `json:"newPassword"`
	}

	err := c.Bind(&forgotPassResetReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to interpret the submitted information")
	}

	if forgotPassResetReq.Code == "" {
		return c.JSON(http.StatusBadRequest, "code can not be empty")
	}

	auth := domain.Auth{Login: forgotPassResetReq.Login, Password: forgotPassResetReq.NewPass}

	ctx := c.Request().Context()

	isValid, message, errValid := ah.AuthValidator.Validate(ctx, &auth)

	if errValid != nil {
		log.Printf("Error validating Auth: %s", errValid.Error())
		return c.JSON(http.StatusInternalServerError, "failed to reset the password")
	}

	if !isValid {
		return c.JSON(http.StatusBadRequest, message)
	}

	code := domain.Code{Identifier: forgotPassResetReq.Login, Value: forgotPassResetReq.Code}

	token, errToken := ah.AuthUseCase.ForgotPassReset(ctx, &code, forgotPassResetReq.NewPass)

	if errToken != nil {
		log.Printf("Error trying to reset user's password: %s", errToken.Error())
		return c.JSON(http.StatusInternalServerError, "failed to reset the password")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": string(token)})
}
