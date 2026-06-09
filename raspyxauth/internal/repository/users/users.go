package users

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	Pool *pgxpool.Pool
}

func NewUsersRepository(pool *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{
		Pool: pool,
	}
}
