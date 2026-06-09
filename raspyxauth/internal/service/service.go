package service

import (
	"JWTAuth/config"
	"JWTAuth/internal/models"
	"JWTAuth/internal/repository"
	"JWTAuth/internal/service/refreshSessions"
	"JWTAuth/internal/service/users"
	"github.com/golang-jwt/jwt/v5"
)

type UsersService interface {
	CreateUser(dataUserReq *models.AddUserRequest) (string, error)

	GetUsers() (*[]models.GetUserResponse, error)
	GetUserByUUID(userUUID string) (*models.GetUserResponse, error)
	GetUsersByUsername(username string) (*[]models.GetUserResponse, error)
	GetUsersByAccessLevel(userAccessLevel int) (*[]models.GetUserResponse, error)

	UpdateUser(idUser string, dataUser *models.UpdateUserRequest) error

	DeleteUser(idUser string) error

	Login(userReqData *models.LoginData) (*models.User, error)
}

type RefreshSessionsService interface {
	CreateJWT(data *models.CreateJWT) (*models.AccessRefreshTokens, error)

	VerifyJWT(tokenStr string) (*jwt.MapClaims, error)

	Refresh(refreshData *models.RefreshData) (*models.AccessRefreshTokens, error)
}

type Service struct {
	UsersService
	RefreshSessionsService
}

func NewService(cfg *config.Config, repo *repository.Repository) *Service {
	return &Service{
		UsersService:           users.NewUsersService(cfg, repo),
		RefreshSessionsService: refreshSessions.NewRefreshSessionsService(cfg, repo),
	}
}
