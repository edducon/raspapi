package users

import (
	"JWTAuth/config"
	"JWTAuth/internal/service"
	"log/slog"
)

type UsersHandler struct {
	log     *slog.Logger
	cfg     *config.Config
	service *service.Service
}

func NewUsersHandler(log *slog.Logger, cfg *config.Config, service *service.Service) *UsersHandler {
	return &UsersHandler{
		log:     log,
		cfg:     cfg,
		service: service,
	}
}
