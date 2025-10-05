package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	userDomain "paolojulian.dev/inventory/domain/user"
	warehouseDomain "paolojulian.dev/inventory/domain/warehouse"
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
		if err == pgx.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Save(ctx context.Context, userToSave *userDomain.User) (*userDomain.User, error) {
	row := r.db.QueryRow(ctx, `
		INSERT INTO users (id, username, password, role, is_active, email, first_name, last_name, mobile)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, username, password, role, is_active, email, first_name, last_name, mobile
	`, userToSave.ID, userToSave.Username, userToSave.Password, userToSave.Role, userToSave.IsActive, userToSave.Email, userToSave.FirstName, userToSave.LastName, userToSave.Mobile)

	var created userDomain.User

	if err := row.Scan(
		&created.ID,
		&created.Username,
		&created.Password,
		&created.Role,
		&created.IsActive,
		&created.Email,
		&created.FirstName,
		&created.LastName,
		&created.Mobile,
	); err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *UserRepository) GetSummary(ctx context.Context, userID string) (*userDomain.UserSummary, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, first_name, last_name
		FROM users
		WHERE id = $1
	`, userID)

	var summary userDomain.UserSummary
	if err := row.Scan(
		&summary.ID,
		&summary.FirstName,
		&summary.LastName,
	); err != nil {
		return nil, err
	}

	return &summary, nil
}

// GetWarehouseSummary temporarily uses user data as warehouse data
// TODO: Create proper warehouse repository when warehouse management is implemented
func (r *UserRepository) GetWarehouseSummary(ctx context.Context, warehouseID string) (*warehouseDomain.WarehouseSummary, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, first_name
		FROM users
		WHERE id = $1
	`, warehouseID)

	var summary warehouseDomain.WarehouseSummary
	if err := row.Scan(
		&summary.ID,
		&summary.Name,
	); err != nil {
		return nil, err
	}

	return &summary, nil
}
