package mysql

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetByLoginNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"id", "login", "password"})

	query := "SELECT id, login, password FROM auth WHERE login = \\?;"

	mock.ExpectQuery(query).WillReturnRows(rows)

	authMysqlRepository := NewAuthMysqlRepository(db)

	auth, err := authMysqlRepository.GetByLogin(context.Background(), "login")

	assert.NoError(t, err)
	assert.Nil(t, auth)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetByLoginError(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := "SELECT id, login, password FROM auth WHERE login = \\?;"

	mock.ExpectQuery(query).WillReturnError(errors.New("error message"))

	authMysqlRepository := NewAuthMysqlRepository(db)

	_, err = authMysqlRepository.GetByLogin(context.Background(), "login")

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetByLogin(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"id", "login", "password"}).AddRow(1, "login", "password")

	query := "SELECT id, login, password FROM auth WHERE login = \\?;"

	mock.ExpectQuery(query).WillReturnRows(rows)

	authMysqlRepository := NewAuthMysqlRepository(db)

	auth, err := authMysqlRepository.GetByLogin(context.Background(), "login")

	assert.NoError(t, err)
	assert.Equal(t, int64(1), auth.ID)
	assert.Equal(t, "login", auth.Login)
	assert.Equal(t, "password", auth.Password)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestStoreWithUserStoreUserError(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := regexp.QuoteMeta("INSERT INTO users (email, first_name, last_name, phone_number, address_city, address_state, address_neighborhood, address_street, address_number, address_zipcode) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")

	mock.ExpectBegin()
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs("", "", "", "", "", "", "", "", "", "").WillReturnError(errors.New("error message"))
	mock.ExpectRollback()

	authMysqlRepository := NewAuthMysqlRepository(db)

	err = authMysqlRepository.StoreWithUser(context.Background(), &domain.Auth{}, &domain.User{})

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestStoreWithUserStoreAuthError(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	storeUserQuery := regexp.QuoteMeta("INSERT INTO users (email, first_name, last_name, phone_number, address_city, address_state, address_neighborhood, address_street, address_number, address_zipcode) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	storeAuthQuery := regexp.QuoteMeta("INSERT INTO auth (login, password) VALUES (?, ?);")

	mock.ExpectBegin()
	mock.ExpectPrepare(storeUserQuery)
	mock.ExpectExec(storeUserQuery).WithArgs("", "", "", "", "", "", "", "", "", "").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare(storeAuthQuery)
	mock.ExpectExec(storeAuthQuery).WithArgs("", "").WillReturnError(errors.New("error message"))
	mock.ExpectRollback()

	authMysqlRepository := NewAuthMysqlRepository(db)

	err = authMysqlRepository.StoreWithUser(context.Background(), &domain.Auth{}, &domain.User{})

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestStoreWithUser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	storeUserQuery := regexp.QuoteMeta("INSERT INTO users (email, first_name, last_name, phone_number, address_city, address_state, address_neighborhood, address_street, address_number, address_zipcode) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	storeAuthQuery := regexp.QuoteMeta("INSERT INTO auth (login, password) VALUES (?, ?);")

	mock.ExpectBegin()
	mock.ExpectPrepare(storeUserQuery)
	mock.ExpectExec(storeUserQuery).WithArgs("", "", "", "", "", "", "", "", "", "").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare(storeAuthQuery)
	mock.ExpectExec(storeAuthQuery).WithArgs("", "").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	authMysqlRepository := NewAuthMysqlRepository(db)

	err = authMysqlRepository.StoreWithUser(context.Background(), &domain.Auth{}, &domain.User{})

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
