package presentation

import (
	"log"
	"net/http"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	ProductUseCase domain.ProductUseCase
}

func NewProductHandler(e *echo.Echo, puc domain.ProductUseCase) *productHandler {
	handler := &productHandler{
		ProductUseCase: puc,
	}
	e.GET("/produtos/:uuid", handler.Get)

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
