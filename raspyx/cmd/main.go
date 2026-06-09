package main

import (
	"log"
	"raspyx2/config"
	"raspyx2/internal/app"
)

// @title           Raspyx
// @version         2.0.0
// @description     API for schedules

// @BasePath  /raspyx/api

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
