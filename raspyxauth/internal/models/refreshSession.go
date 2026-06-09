package models

import (
	"JWTAuth/config"
	"time"
)

type RefreshSession struct {
	UserUUID     string
	RefreshToken string
	UserAgent    string
	Fingerprint  string
	Ip           string
	ExpiresIn    int64
	CreatedAt    time.Time
}

type AddRefreshSessionRequest struct {
	UserUUID  string
	UserAgent string
	Ip        string
}

type AccessRefreshTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken = string

type CreateJWT struct {
	UserUUID    string
	UserAgent   string
	Ip          string
	AccessLevel int
	Config      *config.Config
}

type LoginData struct {
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
}

type RegData struct {
	Username string `json:"username" example:"username"`
	Password string `json:"password" example:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" example:"4b6c34bd-01f0-4fbe-be65-217984b3e33d"`
}

type RefreshData struct {
	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"user_agent"`
	Ip           string `json:"ip"`
}
