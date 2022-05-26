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

func TestStoreError(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := regexp.QuoteMeta("INSERT INTO code (value, identifier) VALUES (?, ?);")

	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs("value", "identifier").WillReturnError(errors.New("error message"))

	codeMysqlRepository := NewCodeMysqlRepository(db)

	err = codeMysqlRepository.Store(context.Background(), &domain.Code{Value: "value", Identifier: "identifier"})

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := regexp.QuoteMeta("INSERT INTO code (value, identifier) VALUES (?, ?);")

	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs("value", "identifier").WillReturnResult(sqlmock.NewResult(1, 1))

	codeMysqlRepository := NewCodeMysqlRepository(db)

	err = codeMysqlRepository.Store(context.Background(), &domain.Code{Value: "value", Identifier: "identifier"})

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetByValueNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"value", "identifier"})

	query := regexp.QuoteMeta("SELECT value, identifier FROM code WHERE value = ?;")

	mock.ExpectQuery(query).WillReturnRows(rows)

	codeMysqlRepository := NewCodeMysqlRepository(db)

	code, err := codeMysqlRepository.GetByValue(context.Background(), "value")

	assert.NoError(t, err)
	assert.Nil(t, code)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetByValueError(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := regexp.QuoteMeta("SELECT value, identifier FROM code WHERE value = ?;")

	mock.ExpectQuery(query).WillReturnError(errors.New("error message"))

	codeMysqlRepository := NewCodeMysqlRepository(db)

	_, err = codeMysqlRepository.GetByValue(context.Background(), "value")

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetByValue(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"value", "identifier"}).AddRow("value", "identifier")

	query := regexp.QuoteMeta("SELECT value, identifier FROM code WHERE value = ?;")

	mock.ExpectQuery(query).WillReturnRows(rows)

	codeMysqlRepository := NewCodeMysqlRepository(db)

	code, err := codeMysqlRepository.GetByValue(context.Background(), "value")

	assert.NoError(t, err)
	assert.Equal(t, "value", code.Value)
	assert.Equal(t, "identifier", code.Identifier)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestDeleteError(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := regexp.QuoteMeta("DELETE FROM code WHERE value = ?;")

	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs("value").WillReturnError(errors.New("error message"))

	codeMysqlRepository := NewCodeMysqlRepository(db)

	err = codeMysqlRepository.DeleteByValue(context.Background(), "value")

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := regexp.QuoteMeta("DELETE FROM code WHERE value = ?;")

	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs("value").WillReturnResult(sqlmock.NewResult(0, 1))

	codeMysqlRepository := NewCodeMysqlRepository(db)

	err = codeMysqlRepository.DeleteByValue(context.Background(), "value")

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
