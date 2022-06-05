package repository

import (
	"context"
	"database/sql"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
)

type userMysqlRepository struct {
	Conn *sql.DB
}

func NewUserMysqlRepository(conn *sql.DB) domain.UserRepository {
	return &userMysqlRepository{Conn: conn}
}

func (r *userMysqlRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT id, uuid, email, first_name, last_name, phone_number, address_city, address_state, address_neighborhood, address_street, address_number, address_zipcode FROM user WHERE email = ?;`

	row := r.Conn.QueryRowContext(ctx, query, email)

	var res domain.User

	if err := row.Scan(&res.ID, &res.UUID, &res.Email, &res.FirstName, &res.LastName, &res.PhoneNumber, &res.Address.City, &res.Address.State, &res.Address.Neighborhood, &res.Address.Street, &res.Address.Number, &res.Address.ZipCode); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}
