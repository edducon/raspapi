package refreshSessions

import (
	"JWTAuth/config"
	"JWTAuth/internal/service"
	"log/slog"
)

type RefreshSessionsHandler struct {
	log     *slog.Logger
	cfg     *config.Config
	service *service.Service
}

func NewRefreshSessionsHandler(log *slog.Logger, cfg *config.Config, service *service.Service) *RefreshSessionsHandler {
	return &RefreshSessionsHandler{
		log:     log,
		cfg:     cfg,
		service: service,
	}
}
