package handler

import (
	_ "JWTAuth/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

	auth := routes.Group("/auth")

	auth.Use(h.RequestLogger())

	auth.StaticFS("/docs", http.Dir("./docs"))
	
	api := auth.Group("/api/v1")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	{
		api.POST("/login", h.RefreshSessionsHandler.Login)
		api.POST("/reg", h.RefreshSessionsHandler.Registration)
		api.POST("/refresh", h.RefreshSessionsHandler.Refresh)
		api.GET("/verify", h.RefreshSessionsHandler.Verify)

		user := api.Group("/users").Use(h.Auth())
		mod := user.Use(h.CheckAL(50))
		{
			mod.GET("/", h.UsersHandler.GetUsers)
			mod.GET("/uuid/:uuid", h.UsersHandler.GetUserByUUID)
			mod.GET("/username/:username", h.UsersHandler.GetUsersByUsername)
			mod.GET("/accesslevel/:access_level", h.UsersHandler.GetUsersByAccessLevel)
		}
		adm := user.Use(h.CheckAL(99))
		{
			adm.POST("/", h.UsersHandler.CreateUser)
			adm.PUT("/:uuid", h.UsersHandler.UpdateUser)
			adm.DELETE("/:uuid", h.UsersHandler.DeleteUser)
		}

	}

	return routes
}
