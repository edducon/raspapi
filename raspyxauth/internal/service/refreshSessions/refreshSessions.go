package refreshSessions

import (
	"JWTAuth/config"
	"JWTAuth/internal/repository"
)

type RefreshSessionsService struct {
	cfg  *config.Config
	repo *repository.Repository
}

func NewRefreshSessionsService(cfg *config.Config, repo *repository.Repository) *RefreshSessionsService {
	return &RefreshSessionsService{
		cfg:  cfg,
		repo: repo,
	}
}
