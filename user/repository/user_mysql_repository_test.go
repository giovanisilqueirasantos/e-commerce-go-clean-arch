package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetByEmailNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"id", "uuid", "email", "first_name", "last_name", "phone_number", "address_city", "address_state", "address_neighborhood", "address_street", "address_number", "address_zipcode"})

	query := regexp.QuoteMeta("SELECT id, uuid, email, first_name, last_name, phone_number, address_city, address_state, address_neighborhood, address_street, address_number, address_zipcode FROM user WHERE email = ?;")

	mock.ExpectQuery(query).WillReturnRows(rows)

	userMysqlRepository := NewUserMysqlRepository(db)

	user, err := userMysqlRepository.GetByEmail(context.Background(), "email")

	assert.NoError(t, err)
	assert.Nil(t, user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetByLoginError(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := regexp.QuoteMeta("SELECT id, uuid, email, first_name, last_name, phone_number, address_city, address_state, address_neighborhood, address_street, address_number, address_zipcode FROM user WHERE email = ?;")

	mock.ExpectQuery(query).WillReturnError(errors.New("error message"))

	userMysqlRepository := NewUserMysqlRepository(db)

	_, err = userMysqlRepository.GetByEmail(context.Background(), "email")

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

	rows := sqlmock.NewRows([]string{"id", "uuid", "email", "first_name", "last_name", "phone_number", "address_city", "address_state", "address_neighborhood", "address_street", "address_number", "address_zipcode"}).AddRow(1, "uuid", "email", "first_name", "last_name", "phone_number", "address_city", "address_state", "address_neighborhood", "address_street", "address_number", "address_zipcode")

	query := regexp.QuoteMeta("SELECT id, uuid, email, first_name, last_name, phone_number, address_city, address_state, address_neighborhood, address_street, address_number, address_zipcode FROM user WHERE email = ?;")

	mock.ExpectQuery(query).WillReturnRows(rows)

	userMysqlRepository := NewUserMysqlRepository(db)

	user, err := userMysqlRepository.GetByEmail(context.Background(), "email")

	assert.NoError(t, err)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "uuid", user.UUID)
	assert.Equal(t, "email", user.Email)
	assert.Equal(t, "first_name", user.FirstName)
	assert.Equal(t, "last_name", user.LastName)
	assert.Equal(t, "phone_number", user.PhoneNumber)
	assert.Equal(t, "address_city", user.Address.City)
	assert.Equal(t, "address_state", user.Address.State)
	assert.Equal(t, "address_neighborhood", user.Address.Neighborhood)
	assert.Equal(t, "address_street", user.Address.Street)
	assert.Equal(t, "address_number", user.Address.Number)
	assert.Equal(t, "address_zipcode", user.Address.ZipCode)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
