package repository

import (
	"context"
	"database/sql"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type productMysqlRepository struct {
	Conn *sql.DB
}

func NewProductMysqlRepository(conn *sql.DB) domain.ProductRepository {
	return &productMysqlRepository{Conn: conn}
}

func (pmr *productMysqlRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Product, error) {
	query := `SELECT id, uuid, name, detail FROM product WHERE uuid = ?;`

	row := pmr.Conn.QueryRowContext(ctx, query, uuid)

	var res domain.Product

	if err := row.Scan(&res.ID, &res.UUID, &res.Name, &res.Detail); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}
