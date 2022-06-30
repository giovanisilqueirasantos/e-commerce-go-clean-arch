package presentation

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetEmptyUUID(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/products/", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewProductHandler(echo.New(), nil)

	handler.Get(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestGetErrorFind(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/products/:uuid", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("uuid")
	c.SetParamValues("testuuid")

	mockProductUsecase := new(mocks.MockProductUsecase)

	mockProductUsecase.On("Get", mock.Anything, "testuuid").Return(nil, errors.New("error message"))

	handler := NewProductHandler(echo.New(), mockProductUsecase)

	handler.Get(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestGetSuccess(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/products/:uuid", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("uuid")
	c.SetParamValues("testuuid")

	mockProductUsecase := new(mocks.MockProductUsecase)

	mockProductUsecase.On("Get", mock.Anything, "testuuid").Return(1, "uuid", 2, "picturepath", "name", "detail", true, "color", "black", nil)

	handler := NewProductHandler(echo.New(), mockProductUsecase)

	handler.Get(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"ID\":1,\"uuid\":\"uuid\",\"rate\":2,\"pictures\":[\"picturepath\"],\"name\":\"name\",\"detail\":\"detail\",\"favorite\":true,\"attributes\":[{\"label\":\"color\",\"values\":[\"black\"]}]}\n", rec.Body.String())
}
