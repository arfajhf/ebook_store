package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	database *pgxpool.Pool
}

func NewPostgresRepository(
	database *pgxpool.Pool,
) Repository {
	return &postgresRepository{
		database: database,
	}
}

func (repository *postgresRepository) Create(
	ctx context.Context,
	user User,
) (User, error) {
	const query = `
		INSERT INTO users (
			name,
			email,
			password_hash,
			role
		)
		VALUES ($1, $2, $3, 'customer')
		RETURNING
			id,
			name,
			email,
			role,
			is_active,
			created_at
	`

	var createdUser User

	err := repository.database.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.Email,
		&createdUser.Role,
		&createdUser.IsActive,
		&createdUser.CreatedAt,
	)
	if err != nil {
		var postgresError *pgconn.PgError

		if errors.As(err, &postgresError) &&
			postgresError.Code == "23505" {
			return User{}, ErrEmailAlreadyUsed
		}

		return User{}, fmt.Errorf("create user: %w", err)
	}

	return createdUser, nil
}
