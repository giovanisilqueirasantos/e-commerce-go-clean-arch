package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain/mocks"
	"github.com/labstack/echo/v4"
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

	handler := AuthHandler{}

	handler.Login(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestLoginErrorValidatingAuth(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/login",
		strings.NewReader("{\"login\":\"invalid login\",\"password\":\"invalid password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Login = "invalid login"
	mockAuth.Password = "invalid password"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(false, "error message", errors.New("error message"))

	handler := AuthHandler{
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
		strings.NewReader("{\"login\":\"invalid login\",\"password\":\"invalid password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Login = "invalid login"
	mockAuth.Password = "invalid password"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(false, "error message", nil)

	handler := AuthHandler{
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
		strings.NewReader("{\"login\":\"valid login\",\"password\":\"valid password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
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
		"/login", strings.NewReader("{\"login\":\"valid login\",\"password\":\"valid password\"}"),
	)
	assert.NoError(t, err)
	req.Header.Add("content-type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
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

func TestSignUpWrongBody(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/signup", strings.NewReader("invalidbody"))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := AuthHandler{}

	handler.SignUp(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestSignUpErrorValidatingAuth(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/signup",
		strings.NewReader("{\"login\":\"invalid login\",\"password\":\"invalid password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Login = "invalid login"
	mockAuth.Password = "invalid password"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(false, "error message", errors.New("error message"))

	handler := AuthHandler{
		AuthValidator: mockAuthValidator,
	}

	handler.SignUp(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestSignUpAuthInvalid(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/signup",
		strings.NewReader("{\"login\":\"invalid login\",\"password\":\"invalid password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Login = "invalid login"
	mockAuth.Password = "invalid password"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(false, "error message", nil)

	handler := AuthHandler{
		AuthValidator: mockAuthValidator,
	}

	handler.SignUp(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "\"error message\"\n", rec.Body.String())
}

func TestSignUpErrorValidatingUser(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/signup",
		strings.NewReader("{\"login\":\"valid login\",\"password\":\"valid password\",\"confirmPassword\":\"valid confirm password\",\"email\":\"invalidemail@email.com\",\"firstName\":\"invalid first name\",\"lastName\":\"invalid last name\",\"phoneNumber\":\"invalid phone number\",\"address\":\"invalid address\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)
	mockUserValidator := new(mocks.MockUserValidator)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"
	mockAuth.ConfirmPassword = "valid confirm password"

	var mockUser domain.User
	mockUser.Email = "invalidemail@email.com"
	mockUser.FirstName = "invalid first name"
	mockUser.LastName = "invalid last name"
	mockUser.PhoneNumber = "invalid phone number"
	mockUser.Address = "invalid address"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "", nil)
	mockUserValidator.On("Validate", mock.Anything, &mockUser).Return(false, "error message", errors.New("error message"))

	handler := AuthHandler{
		AuthValidator: mockAuthValidator,
		UserValidator: mockUserValidator,
	}

	handler.SignUp(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestSignUpUserInvalid(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/signup",
		strings.NewReader("{\"login\":\"valid login\",\"password\":\"valid password\",\"confirmPassword\":\"valid confirm password\",\"email\":\"invalidemail@email.com\",\"firstName\":\"invalid first name\",\"lastName\":\"invalid last name\",\"phoneNumber\":\"invalid phone number\",\"address\":\"invalid address\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)
	mockUserValidator := new(mocks.MockUserValidator)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"
	mockAuth.ConfirmPassword = "valid confirm password"

	var mockUser domain.User
	mockUser.Email = "invalidemail@email.com"
	mockUser.FirstName = "invalid first name"
	mockUser.LastName = "invalid last name"
	mockUser.PhoneNumber = "invalid phone number"
	mockUser.Address = "invalid address"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "", nil)
	mockUserValidator.On("Validate", mock.Anything, &mockUser).Return(false, "error message", nil)

	handler := AuthHandler{
		AuthValidator: mockAuthValidator,
		UserValidator: mockUserValidator,
	}

	handler.SignUp(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "\"error message\"\n", rec.Body.String())
}

func TestSignUpErrorOnSignUp(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/signup",
		strings.NewReader("{\"login\":\"valid login\",\"password\":\"valid password\",\"confirmPassword\":\"valid confirm password\",\"email\":\"validemail@email.com\",\"firstName\":\"valid first name\",\"lastName\":\"valid last name\",\"phoneNumber\":\"valid phone number\",\"address\":\"valid address\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)
	mockUserValidator := new(mocks.MockUserValidator)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"
	mockAuth.ConfirmPassword = "valid confirm password"

	var mockUser domain.User
	mockUser.Email = "validemail@email.com"
	mockUser.FirstName = "valid first name"
	mockUser.LastName = "valid last name"
	mockUser.PhoneNumber = "valid phone number"
	mockUser.Address = "valid address"

	mockAuthUsecase.On("SignUp", mock.Anything, &mockAuth, &mockUser).Return(errors.New("error message"))
	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "", nil)
	mockUserValidator.On("Validate", mock.Anything, &mockUser).Return(true, "", nil)

	handler := AuthHandler{
		AuthUseCase:   mockAuthUsecase,
		AuthValidator: mockAuthValidator,
		UserValidator: mockUserValidator,
	}

	handler.SignUp(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestSignUpSuccess(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/signup",
		strings.NewReader("{\"login\":\"valid login\",\"password\":\"valid password\",\"confirmPassword\":\"valid confirm password\",\"email\":\"validemail@email.com\",\"firstName\":\"valid first name\",\"lastName\":\"valid last name\",\"phoneNumber\":\"valid phone number\",\"address\":\"valid address\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)
	mockUserValidator := new(mocks.MockUserValidator)

	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid password"
	mockAuth.ConfirmPassword = "valid confirm password"

	var mockUser domain.User
	mockUser.Email = "validemail@email.com"
	mockUser.FirstName = "valid first name"
	mockUser.LastName = "valid last name"
	mockUser.PhoneNumber = "valid phone number"
	mockUser.Address = "valid address"

	mockAuthUsecase.On("SignUp", mock.Anything, &mockAuth, &mockUser).Return(nil)
	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "", nil)
	mockUserValidator.On("Validate", mock.Anything, &mockUser).Return(true, "", nil)

	handler := AuthHandler{
		AuthUseCase:   mockAuthUsecase,
		AuthValidator: mockAuthValidator,
		UserValidator: mockUserValidator,
	}

	handler.SignUp(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}
