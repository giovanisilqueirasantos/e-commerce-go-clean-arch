package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetByUUIDNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"id", "uuid", "name", "detail"})

	query := regexp.QuoteMeta("SELECT id, uuid, name, detail FROM product WHERE uuid = ?;")

	mock.ExpectQuery(query).WillReturnRows(rows)

	productMysqlRepository := NewProductMysqlRepository(db)

	product, err := productMysqlRepository.GetByUUID(context.Background(), "testuuid")

	assert.NoError(t, err)
	assert.Nil(t, product)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetByUUIDError(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := regexp.QuoteMeta("SELECT id, uuid, name, detail FROM product WHERE uuid = ?;")

	mock.ExpectQuery(query).WillReturnError(errors.New("error message"))

	productMysqlRepository := NewProductMysqlRepository(db)

	_, err = productMysqlRepository.GetByUUID(context.Background(), "uuid")

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetByUUID(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"id", "uuid", "name", "detail"}).AddRow(1, "uuid", "name", "detail")

	query := regexp.QuoteMeta("SELECT id, uuid, name, detail FROM product WHERE uuid = ?;")

	mock.ExpectQuery(query).WillReturnRows(rows)

	productMysqlRepository := NewProductMysqlRepository(db)

	product, err := productMysqlRepository.GetByUUID(context.Background(), "testuuid")

	assert.NoError(t, err)
	assert.Equal(t, int64(1), product.ID)
	assert.Equal(t, "uuid", product.UUID)
	assert.Equal(t, "name", product.Name)
	assert.Equal(t, "detail", product.Detail)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
