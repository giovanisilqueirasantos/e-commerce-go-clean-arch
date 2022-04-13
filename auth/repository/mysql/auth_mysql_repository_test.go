package mysql

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
}
