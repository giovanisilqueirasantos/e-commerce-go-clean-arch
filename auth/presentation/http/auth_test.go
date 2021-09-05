package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/skeey/e-commerce-go-clean-arch/domain"
	"github.com/skeey/e-commerce-go-clean-arch/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLoginWrongBody(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/login", strings.NewReader("invalidbody"))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth

	mockAuthUsecase.On("Login", mock.Anything, &mockAuth).Return("valid token", nil)
	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "", nil)

	handler := AuthHandler{
		AuthUseCase:   mockAuthUsecase,
		AuthValidator: mockAuthValidator,
	}

	handler.Login(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestLoginErrorValidatingAuth(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/login",
		strings.NewReader("{\"email\":\"invalidemail@email.com\",\"password\":\"invalid password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Email = "invalidemail@email.com"
	mockAuth.Password = "invalid password"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(false, "error message", errors.New("error message"))
	mockAuthUsecase.On("Login", mock.Anything, &mockAuth).Return("valid token", nil)

	handler := AuthHandler{
		AuthUseCase:   mockAuthUsecase,
		AuthValidator: mockAuthValidator,
	}

	handler.Login(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestLoginAuthInvalid(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/login",
		strings.NewReader("{\"email\":\"invalidemail@email.com\",\"password\":\"invalid password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Email = "invalidemail@email.com"
	mockAuth.Password = "invalid password"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(false, "error message", nil)
	mockAuthUsecase.On("Login", mock.Anything, &mockAuth).Return("valid token", nil)

	handler := AuthHandler{
		AuthUseCase:   mockAuthUsecase,
		AuthValidator: mockAuthValidator,
	}

	handler.Login(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "\"error message\"\n", rec.Body.String())
}

func TestLoginErrorGeneratingToken(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/login",
		strings.NewReader("{\"email\":\"validemail@email.com\",\"password\":\"valid password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Email = "validemail@email.com"
	mockAuth.Password = "valid password"

	mockAuthUsecase.On("Login", mock.Anything, &mockAuth).Return("", errors.New("error message"))
	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "", nil)

	handler := AuthHandler{
		AuthUseCase:   mockAuthUsecase,
		AuthValidator: mockAuthValidator,
	}

	handler.Login(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestLoginSuccess(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST,
		"/login", strings.NewReader("{\"email\":\"validemail@email.com\",\"password\":\"valid password\"}"),
	)
	assert.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Email = "validemail@email.com"
	mockAuth.Password = "valid password"

	mockAuthUsecase.On("Login", mock.Anything, &mockAuth).Return("valid token", nil)
	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "", nil)

	handler := AuthHandler{
		AuthUseCase:   mockAuthUsecase,
		AuthValidator: mockAuthValidator,
	}

	err = handler.Login(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"token\":\"valid token\"}\n", rec.Body.String())
}
