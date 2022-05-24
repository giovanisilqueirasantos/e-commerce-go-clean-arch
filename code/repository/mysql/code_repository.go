package mysql

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

	stmt, stmtErr := r.Conn.PrepareContext(ctx, query)

	if stmtErr != nil {
		return stmtErr
	}

	exec, execErr := stmt.ExecContext(ctx, c.Value, c.Identifier)

	if execErr != nil {
		return execErr
	}

	affect, affectErr := exec.RowsAffected()

	if affectErr != nil {
		return affectErr
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

	stmt, stmtErr := r.Conn.PrepareContext(ctx, query)

	if stmtErr != nil {
		return stmtErr
	}

	exec, execErr := stmt.ExecContext(ctx, value)

	if execErr != nil {
		return execErr
	}

	affect, affectErr := exec.RowsAffected()

	if affectErr != nil {
		return affectErr
	}

	if affect != 1 {
		return fmt.Errorf("error trying to remove code with total rows affected: %d", affect)
	}

	return nil
}
