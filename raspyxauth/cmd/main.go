package main

import (
	"JWTAuth/config"
	"JWTAuth/internal/app"
	"log"
)

// @title JWT authentication/authorization service
// @version 1.0.0
// @description Documentation for authentication/authorization API service
// @basePath /auth/api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
