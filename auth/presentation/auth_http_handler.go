package presentation

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

	if err := c.Bind(&auth); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to interpret the submitted information")
	}

	ctx := c.Request().Context()

	isValid, message := ah.AuthValidator.Validate(ctx, &auth)

	if !isValid {
		return c.JSON(http.StatusBadRequest, message)
	}

	token, err := ah.AuthUseCase.Login(ctx, &auth)

	if err != nil {
		log.Printf("Error trying to generate token for Login: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, "failed to login")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": string(token)})
}

func (ah *authHandler) SignUp(c echo.Context) error {
	var authWithUser struct {
		domain.Auth
		domain.User
	}

	if err := c.Bind(&authWithUser); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to interpret the submitted information")
	}

	ctx := c.Request().Context()

	auth := domain.Auth{
		Login:    authWithUser.Login,
		Password: authWithUser.Password,
	}

	isValid, message := ah.AuthValidator.Validate(ctx, &auth)

	if !isValid {
		return c.JSON(http.StatusBadRequest, message)
	}

	user := domain.User{
		Email:       authWithUser.Email,
		FirstName:   authWithUser.FirstName,
		LastName:    authWithUser.LastName,
		PhoneNumber: authWithUser.PhoneNumber,
		Address:     authWithUser.Address,
	}

	isValid, message = ah.UserValidator.Validate(ctx, &user)

	if !isValid {
		return c.JSON(http.StatusBadRequest, message)
	}

	token, err := ah.AuthUseCase.SignUp(ctx, &authWithUser.Auth, &authWithUser.User)

	if err != nil {
		log.Printf("Error trying to sign up: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, "failed to sign up")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": string(token)})
}

func (ah *authHandler) ForgotPassCode(c echo.Context) error {
	var forgotPassReq struct {
		Login string `json:"login"`
	}

	if err := c.Bind(&forgotPassReq); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to interpret the submitted information")
	}

	ctx := c.Request().Context()

	isValid, message := ah.AuthValidator.ValidateLogin(ctx, forgotPassReq.Login)

	if !isValid {
		return c.JSON(http.StatusBadRequest, message)
	}

	if err := ah.AuthUseCase.ForgotPassCode(ctx, forgotPassReq.Login); err != nil {
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

	if err := c.Bind(&forgotPassResetReq); err != nil {
		return c.JSON(http.StatusBadRequest, "failed to interpret the submitted information")
	}

	if forgotPassResetReq.Code == "" {
		return c.JSON(http.StatusBadRequest, "code can not be empty")
	}

	auth := domain.Auth{Login: forgotPassResetReq.Login, Password: forgotPassResetReq.NewPass}

	ctx := c.Request().Context()

	isValid, message := ah.AuthValidator.Validate(ctx, &auth)

	if !isValid {
		return c.JSON(http.StatusBadRequest, message)
	}

	code := domain.Code{Identifier: forgotPassResetReq.Login, Value: forgotPassResetReq.Code}

	token, err := ah.AuthUseCase.ForgotPassReset(ctx, &code, forgotPassResetReq.NewPass)

	if err != nil {
		log.Printf("Error trying to reset user's password: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, "failed to reset the password")
	}

	return c.JSON(http.StatusOK, map[string]string{"token": string(token)})
}
