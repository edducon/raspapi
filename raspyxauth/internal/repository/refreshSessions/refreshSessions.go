package refreshSessions

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshSessionsRepository struct {
	Pool *pgxpool.Pool
}

func NewRefreshSessionsRepository(pool *pgxpool.Pool) *RefreshSessionsRepository {
	return &RefreshSessionsRepository{
		Pool: pool,
	}
}
