package users

import (
	"JWTAuth/config"
	"JWTAuth/internal/repository"
)

type UsersService struct {
	cfg  *config.Config
	repo *repository.Repository
}

func NewUsersService(cfg *config.Config, repo *repository.Repository) *UsersService {
	return &UsersService{
		cfg:  cfg,
		repo: repo,
	}
}
