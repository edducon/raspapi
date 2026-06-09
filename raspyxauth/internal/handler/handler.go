package handler

import (
	"JWTAuth/config"
	"JWTAuth/internal/handler/refreshSessions"
	"JWTAuth/internal/handler/users"
	"JWTAuth/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type UsersHandler interface {
	CreateUser(ctx *gin.Context)

	GetUsers(ctx *gin.Context)
	GetUserByUUID(ctx *gin.Context)
	GetUsersByUsername(ctx *gin.Context)
	GetUsersByAccessLevel(ctx *gin.Context)

	UpdateUser(ctx *gin.Context)

	DeleteUser(ctx *gin.Context)
}

type RefreshSessionsHandler interface {
	Login(ctx *gin.Context)

	Registration(ctx *gin.Context)

	Verify(ctx *gin.Context)

	Refresh(ctx *gin.Context)
}

type Handler struct {
	log      *slog.Logger
	cfg      *config.Config
	services *service.Service
	UsersHandler
	RefreshSessionsHandler
}

func NewHandler(log *slog.Logger, cfg *config.Config, services *service.Service) *Handler {
	return &Handler{
		log:                    log,
		cfg:                    cfg,
		services:               services,
		UsersHandler:           users.NewUsersHandler(log, cfg, services),
		RefreshSessionsHandler: refreshSessions.NewRefreshSessionsHandler(log, cfg, services),
	}
}
