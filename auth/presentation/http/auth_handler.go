package http

import (
	"log"
	"net/http"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthUseCase   domain.AuthUseCase
	AuthValidator domain.AuthValidator
	UserValidator domain.UserValidator
}

func InitAuthHandler(e *echo.Echo, auc domain.AuthUseCase, av domain.AuthValidator, uv domain.UserValidator) {
	handler := &AuthHandler{
		AuthUseCase:   auc,
		AuthValidator: av,
		UserValidator: uv,
	}

	e.POST("/login", handler.Login)
	e.POST("/signup", handler.SignUp)
	e.POST("/forgotpass/code", handler.ForgotPasswordCode)
}

func (ah *AuthHandler) Login(c echo.Context) error {
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

func (ah *AuthHandler) SignUp(c echo.Context) error {
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
		Login:           authWithUser.Login,
		Password:        authWithUser.Password,
		ConfirmPassword: authWithUser.ConfirmPassword,
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

	err = ah.AuthUseCase.SignUp(ctx, &authWithUser.Auth, &authWithUser.User)

	if err != nil {
		log.Printf("Error trying to sign up: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, "failed to sign up")
	}

	return c.String(http.StatusOK, "")
}

func (ah *AuthHandler) ForgotPasswordCode(c echo.Context) error {
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

	err = ah.AuthUseCase.ForgotPasswordCode(ctx, forgotPassReq.Login)

	if err != nil {
		log.Printf("Error trying to send forgot password code: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, "failed to send forgot password code")
	}

	return c.String(http.StatusOK, "")
}
