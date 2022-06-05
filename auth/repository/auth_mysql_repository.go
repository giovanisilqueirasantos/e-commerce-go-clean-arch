package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/google/uuid"
)

type authMysqlRepository struct {
	Conn *sql.DB
}

func NewAuthMysqlRepository(conn *sql.DB) domain.AuthRepository {
	return &authMysqlRepository{Conn: conn}
}

func (r *authMysqlRepository) GetByLogin(ctx context.Context, login string) (*domain.Auth, error) {
	query := `SELECT id, uuid, login, password FROM auth WHERE login = ?;`

	row := r.Conn.QueryRowContext(ctx, query, login)

	var res domain.Auth

	if err := row.Scan(&res.ID, &res.UUID, &res.Login, &res.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

func (r *authMysqlRepository) StoreWithUser(ctx context.Context, a *domain.Auth, u *domain.User) error {
	storeUserQuery := `INSERT INTO users (uuid, email, first_name, last_name, phone_number, address_city, address_state, address_neighborhood, address_street, address_number, address_zipcode) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	storeAuthQuery := `INSERT INTO auth (uuid, login, password) VALUES (?, ?);`

	tx, err := r.Conn.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	storeUserStmt, err := tx.PrepareContext(ctx, storeUserQuery)

	if err != nil {
		return err
	}

	u.UUID = uuid.NewString()
	if _, err = storeUserStmt.ExecContext(ctx, u.UUID, u.Email, u.FirstName, u.LastName, u.PhoneNumber, u.Address.City, u.Address.State, u.Address.Neighborhood, u.Address.Street, u.Address.Number, u.Address.ZipCode); err != nil {
		tx.Rollback()
		return err
	}

	storeAuthStmt, err := tx.PrepareContext(ctx, storeAuthQuery)

	if err != nil {
		return err
	}

	a.UUID = uuid.NewString()
	if _, err = storeAuthStmt.ExecContext(ctx, a.UUID, a.Login, a.Password); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *authMysqlRepository) Update(ctx context.Context, a *domain.Auth) error {
	query := `UPDATE auth SET login=?, password=? WHERE uuid=?;`

	stmt, err := r.Conn.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	exec, err := stmt.ExecContext(ctx, a.Login, a.Password, a.UUID)

	if err != nil {
		return err
	}

	affect, err := exec.RowsAffected()

	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("update wrong with total rows affected: %d", affect)
	}

	return nil
}
