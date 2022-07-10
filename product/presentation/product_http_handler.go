package presentation

import (
	"log"
	"net/http"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	ProductUseCase domain.ProductUseCase
	TokenService   domain.TokenService
}

func NewProductHandler(e *echo.Echo, puc domain.ProductUseCase, ts domain.TokenService) *productHandler {
	handler := &productHandler{
		ProductUseCase: puc,
		TokenService:   ts,
	}

	auth := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, "request not authorized")
			}
			if isValid, err := handler.TokenService.IsValid(c.Request().Context(), domain.Token(authHeader)); err != nil {
				return c.JSON(http.StatusInternalServerError, "failed to authorize request")
			} else if !isValid {
				return c.JSON(http.StatusUnauthorized, "request not authorized")
			}
			return next(c)
		}
	}

	e.GET("/products/:uuid", handler.Get, auth)

	return handler
}

func (ph *productHandler) Get(c echo.Context) error {
	uuid := c.Param("uuid")

	if uuid == "" {
		return c.JSON(http.StatusBadRequest, "uuid param is not valid")
	}

	product, err := ph.ProductUseCase.Get(c.Request().Context(), uuid)

	if err != nil {
		log.Printf("Error trying to get a product: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, "failed to get the product")
	}

	return c.JSON(http.StatusOK, product)
}
