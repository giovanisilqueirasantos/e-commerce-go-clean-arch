package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type codeMysqlRepository struct {
	Conn *sql.DB
}

func NewCodeMysqlRepository(conn *sql.DB) domain.CodeRepository {
	return &codeMysqlRepository{Conn: conn}
}

func (r *codeMysqlRepository) Store(ctx context.Context, c *domain.Code) error {
	query := `INSERT INTO code (value, identifier) VALUES (?, ?);`

	stmt, err := r.Conn.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	exec, err := stmt.ExecContext(ctx, c.Value, c.Identifier)

	if err != nil {
		return err
	}

	affect, err := exec.RowsAffected()

	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("error trying to store code with total rows affected: %d", affect)
	}

	return nil
}

func (r *codeMysqlRepository) GetByValue(ctx context.Context, value string) (*domain.Code, error) {
	query := `SELECT value, identifier FROM code WHERE value = ?;`

	row := r.Conn.QueryRowContext(ctx, query, value)

	var res domain.Code

	if err := row.Scan(&res.Value, &res.Identifier); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

func (r *codeMysqlRepository) DeleteByValue(ctx context.Context, value string) error {
	query := `DELETE FROM code WHERE value = ?;`

	stmt, err := r.Conn.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	exec, err := stmt.ExecContext(ctx, value)

	if err != nil {
		return err
	}

	affect, err := exec.RowsAffected()

	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("error trying to remove code with total rows affected: %d", affect)
	}

	return nil
}
