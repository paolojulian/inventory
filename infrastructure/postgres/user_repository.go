package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	userDomain "paolojulian.dev/inventory/domain/user"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) FindByUsername(ctx context.Context, username string) (*userDomain.User, error) {
	row := repo.db.QueryRow(ctx, `
		SELECT id, username, password, role, is_active, email, first_name, last_name, mobile
		FROM users
		WHERE username = $1
	`, username)

	var user userDomain.User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Mobile,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
