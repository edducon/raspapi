package repository

import (
	"JWTAuth/internal/models"
	"JWTAuth/internal/repository/refreshSessions"
	"JWTAuth/internal/repository/users"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository interface {
	CreateUser(userData *models.User) error

	GetUsers() (*[]models.User, error)
	GetUserByUUID(userUUID string) (*models.User, error)
	GetUsersByUsername(username string) (*[]models.User, error)
	GetUsersByAccessLevel(userAccessLevel int) (*[]models.User, error)

	UpdateUser(userUUID string, userData *models.UpdateUserRequest) error

	DeleteUser(userUUID string) error
}

type RefreshSessionsRepository interface {
	CreateRefreshSession(refreshSessionData *models.RefreshSession) error

	GetRefreshSessionByToken(token string) (*models.RefreshSession, error)

	DeleteRefreshSession(userUUID string) error
}

type Repository struct {
	UsersRepository
	RefreshSessionsRepository
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		UsersRepository:           users.NewUsersRepository(pool),
		RefreshSessionsRepository: refreshSessions.NewRefreshSessionsRepository(pool),
	}
}
