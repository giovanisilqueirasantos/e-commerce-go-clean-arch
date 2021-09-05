package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/skeey/e-commerce-go-clean-arch/domain"
)

type AuthHandler struct {
	AuthUseCase   domain.AuthUseCase
	AuthValidator domain.AuthValidator
}

func InitAuthHandler(e *echo.Echo, auc domain.AuthUseCase, av domain.AuthValidator) {
	handler := &AuthHandler{
		AuthUseCase:   auc,
		AuthValidator: av,
	}

	e.POST("/login", handler.Login)
	// e.POST("/signup", handler.SignUp)
	// e.POST("/forgotpass", handler.ForgotPassword)
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

// func (ah *AuthHandler) SignUp(c echo.Context) error {

// }

// func (ah *AuthHandler) ForgotPassword(c echo.Context) error {

// }
