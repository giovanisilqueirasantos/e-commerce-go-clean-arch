package presentation

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

	handler := NewAuthHandler(echo.New(), nil, nil, nil)

	handler.Login(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
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

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(false, "error message")

	handler := NewAuthHandler(echo.New(), nil, mockAuthValidator, nil)

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
	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "")

	handler := NewAuthHandler(echo.New(), mockAuthUsecase, mockAuthValidator, nil)

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
	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "")

	handler := NewAuthHandler(echo.New(), mockAuthUsecase, mockAuthValidator, nil)

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

	handler := NewAuthHandler(echo.New(), nil, nil, nil)

	handler.SignUp(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
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

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(false, "error message")

	handler := NewAuthHandler(echo.New(), nil, mockAuthValidator, nil)

	handler.SignUp(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "\"error message\"\n", rec.Body.String())
}

func TestSignUpUserInvalid(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/signup",
		strings.NewReader("{\"login\":\"valid login\",\"password\":\"valid password\",\"confirmPassword\":\"valid confirm password\",\"email\":\"invalidemail@email.com\",\"firstName\":\"invalid first name\",\"lastName\":\"invalid last name\",\"phoneNumber\":\"invalid phone number\",\"address\":{\"city\":\"invalid city\",\"state\":\"invalid state\",\"neighborhood\":\"invalid neighborhood\",\"street\":\"invalid street\",\"number\":\"invalid number\",\"zipcode\":\"invalid zipcode\"}}"),
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

	var mockUser domain.User
	mockUser.Email = "invalidemail@email.com"
	mockUser.FirstName = "invalid first name"
	mockUser.LastName = "invalid last name"
	mockUser.PhoneNumber = "invalid phone number"
	mockUser.Address = domain.UserAddress{
		City:         "invalid city",
		State:        "invalid state",
		Neighborhood: "invalid neighborhood",
		Street:       "invalid street",
		Number:       "invalid number",
		ZipCode:      "invalid zipcode",
	}

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "")
	mockUserValidator.On("Validate", mock.Anything, &mockUser).Return(false, "error message")

	handler := NewAuthHandler(echo.New(), nil, mockAuthValidator, mockUserValidator)

	handler.SignUp(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "\"error message\"\n", rec.Body.String())
}

func TestSignUpErrorOnSignUp(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/signup",
		strings.NewReader("{\"login\":\"valid login\",\"password\":\"valid password\",\"confirmPassword\":\"valid confirm password\",\"email\":\"validemail@email.com\",\"firstName\":\"valid first name\",\"lastName\":\"valid last name\",\"phoneNumber\":\"valid phone number\",\"address\":{\"city\":\"valid city\",\"state\":\"valid state\",\"neighborhood\":\"valid neighborhood\",\"street\":\"valid street\",\"number\":\"valid number\",\"zipcode\":\"valid zipcode\"}}"),
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

	var mockUser domain.User
	mockUser.Email = "validemail@email.com"
	mockUser.FirstName = "valid first name"
	mockUser.LastName = "valid last name"
	mockUser.PhoneNumber = "valid phone number"
	mockUser.Address = domain.UserAddress{
		City:         "valid city",
		State:        "valid state",
		Neighborhood: "valid neighborhood",
		Street:       "valid street",
		Number:       "valid number",
		ZipCode:      "valid zipcode",
	}

	mockAuthUsecase.On("SignUp", mock.Anything, &mockAuth, &mockUser).Return("", errors.New("error message"))
	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "")
	mockUserValidator.On("Validate", mock.Anything, &mockUser).Return(true, "")

	handler := NewAuthHandler(echo.New(), mockAuthUsecase, mockAuthValidator, mockUserValidator)

	handler.SignUp(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestSignUpSuccess(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/signup",
		strings.NewReader("{\"login\":\"valid login\",\"password\":\"valid password\",\"confirmPassword\":\"valid confirm password\",\"email\":\"validemail@email.com\",\"firstName\":\"valid first name\",\"lastName\":\"valid last name\",\"phoneNumber\":\"valid phone number\",\"address\":{\"city\":\"valid city\",\"state\":\"valid state\",\"neighborhood\":\"valid neighborhood\",\"street\":\"valid street\",\"number\":\"valid number\",\"zipcode\":\"valid zipcode\"}}"),
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

	var mockUser domain.User
	mockUser.Email = "validemail@email.com"
	mockUser.FirstName = "valid first name"
	mockUser.LastName = "valid last name"
	mockUser.PhoneNumber = "valid phone number"
	mockUser.Address = domain.UserAddress{
		City:         "valid city",
		State:        "valid state",
		Neighborhood: "valid neighborhood",
		Street:       "valid street",
		Number:       "valid number",
		ZipCode:      "valid zipcode",
	}

	mockAuthUsecase.On("SignUp", mock.Anything, &mockAuth, &mockUser).Return("valid token", nil)
	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "")
	mockUserValidator.On("Validate", mock.Anything, &mockUser).Return(true, "")

	handler := NewAuthHandler(echo.New(), mockAuthUsecase, mockAuthValidator, mockUserValidator)

	handler.SignUp(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"token\":\"valid token\"}\n", rec.Body.String())
}

func TestForgotPassCodeWrongBody(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/forgotpass/code", strings.NewReader("invalidbody"))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewAuthHandler(echo.New(), nil, nil, nil)

	handler.ForgotPassCode(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestForgotPassCodeErrorInvalidLogin(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/forgotpass/code",
		strings.NewReader("{\"login\":\"invalid login\""),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)

	mockAuthValidator.On("ValidateLogin", mock.Anything, "invalid login").Return(false, "error message")

	handler := NewAuthHandler(echo.New(), nil, mockAuthValidator, nil)

	handler.ForgotPassCode(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestForgotPassCodeErrorSendingCode(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/forgotpass/code",
		strings.NewReader("{\"login\":\"valid login\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)

	mockAuthUsecase.On("ForgotPassCode", mock.Anything, "valid login").Return(errors.New("error message"))
	mockAuthValidator.On("ValidateLogin", mock.Anything, "valid login").Return(true, "")

	handler := NewAuthHandler(echo.New(), mockAuthUsecase, mockAuthValidator, nil)

	handler.ForgotPassCode(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestForgotPassCodeSuccess(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/forgotpass/code",
		strings.NewReader("{\"login\":\"valid login\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthUsecase := new(mocks.MockAuthUsecase)
	mockAuthValidator := new(mocks.MockAuthValidator)

	mockAuthUsecase.On("ForgotPassCode", mock.Anything, "valid login").Return(nil)
	mockAuthValidator.On("ValidateLogin", mock.Anything, "valid login").Return(true, "")

	handler := NewAuthHandler(echo.New(), mockAuthUsecase, mockAuthValidator, nil)

	handler.ForgotPassCode(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestForgotPassResetWrongBody(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/forgotpass/reset", strings.NewReader("invalidbody"))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewAuthHandler(echo.New(), nil, nil, nil)

	handler.ForgotPassReset(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestForgotPassResetErrorEmptyCode(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/forgotpass/reset",
		strings.NewReader("{\"login\":\"valid login\",\"code\":\"\",\"newPassword\":\"valid new password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := NewAuthHandler(echo.New(), nil, nil, nil)

	handler.ForgotPassReset(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestForogtPassResetAuthInvalid(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/forgotpass/reset",
		strings.NewReader("{\"login\":\"invalid login\",\"code\":\"valid code\",\"newPassword\":\"invalid new password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Login = "invalid login"
	mockAuth.Password = "invalid new password"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(false, "error message")

	handler := NewAuthHandler(echo.New(), nil, mockAuthValidator, nil)

	handler.ForgotPassReset(c)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "\"error message\"\n", rec.Body.String())
}

func TestForgotPassResetErrorResettingPassword(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/forgotpass/reset",
		strings.NewReader("{\"login\":\"valid login\",\"code\":\"valid code\",\"newPassword\":\"valid new password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid new password"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "")

	mockAuthUsecase := new(mocks.MockAuthUsecase)

	code := domain.Code{Value: "valid code", Identifier: mockAuth.Login}

	mockAuthUsecase.On("ForgotPassReset", mock.Anything, &code, mockAuth.Password).Return("", errors.New("error message"))

	handler := NewAuthHandler(echo.New(), mockAuthUsecase, mockAuthValidator, nil)

	handler.ForgotPassReset(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.NotEqual(t, "", rec.Body.String())
}

func TestForgotPassResetSuccess(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(
		echo.POST, "/forgotpass/reset",
		strings.NewReader("{\"login\":\"valid login\",\"code\":\"valid code\",\"newPassword\":\"valid new password\"}"),
	)
	req.Header.Add("content-type", "application/json")
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthValidator := new(mocks.MockAuthValidator)
	var mockAuth domain.Auth
	mockAuth.Login = "valid login"
	mockAuth.Password = "valid new password"

	mockAuthValidator.On("Validate", mock.Anything, &mockAuth).Return(true, "")

	mockAuthUsecase := new(mocks.MockAuthUsecase)

	code := domain.Code{Value: "valid code", Identifier: mockAuth.Login}

	mockAuthUsecase.On("ForgotPassReset", mock.Anything, &code, mockAuth.Password).Return("valid token", nil)

	handler := NewAuthHandler(echo.New(), mockAuthUsecase, mockAuthValidator, nil)

	handler.ForgotPassReset(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"token\":\"valid token\"}\n", rec.Body.String())

}
