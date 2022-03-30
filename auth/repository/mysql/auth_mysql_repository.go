package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type authMysqlRepository struct {
	Conn *sql.DB
}

func NewAuthMysqlRepository(Conn *sql.DB) domain.AuthRepository {
	return &authMysqlRepository{Conn}
}

func (r *authMysqlRepository) GetByLogin(ctx context.Context, login string) (*domain.Auth, error) {
	query := `SELECT id, login, password FROM auth WHERE login = ?;`

	row := r.Conn.QueryRowContext(ctx, query, login)

	var res domain.Auth

	if err := row.Scan(&res.ID, &res.Login, &res.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

func (r *authMysqlRepository) StoreWithUser(ctx context.Context, a *domain.Auth, u *domain.User) error {
	storeUserQuery := `INSERT INTO users (email, first_name, last_name, phone_number, address_city, address_state, address_neighborhood, address_street, address_number, address_zipcode) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	storeAuthQuery := `INSERT INTO auth (login, password) VALUES (?, ?);`

	tx, err := r.Conn.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	storeUserStmt, storeUserStmtErr := tx.PrepareContext(ctx, storeUserQuery)

	if storeUserStmtErr != nil {
		return storeUserStmtErr
	}

	if _, err = storeUserStmt.ExecContext(ctx, u.Email, u.FirstName, u.LastName, u.PhoneNumber, u.Address.City, u.Address.State, u.Address.Neighborhood, u.Address.Street, u.Address.Number, u.Address.ZipCode); err != nil {
		tx.Rollback()
		return err
	}

	storeAuthStmt, storeAuthStmtErr := tx.PrepareContext(ctx, storeAuthQuery)

	if storeAuthStmtErr != nil {
		return storeAuthStmtErr
	}

	if _, err = storeAuthStmt.ExecContext(ctx, a.Login, a.Password); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *authMysqlRepository) Update(ctx context.Context, a *domain.Auth) error {
	query := `UPDATE auth SET login=?, password=? WHERE id=?;`

	stmt, stmtErr := r.Conn.PrepareContext(ctx, query)

	if stmtErr != nil {
		return stmtErr
	}

	exec, execErr := stmt.ExecContext(ctx, a.Login, a.Password, a.ID)

	if execErr != nil {
		return execErr
	}

	affect, affectErr := exec.RowsAffected()

	if affectErr != nil {
		return affectErr
	}

	if affect != 1 {
		return fmt.Errorf("update wrong with total rows affected: %d", affect)
	}

	return nil
}
