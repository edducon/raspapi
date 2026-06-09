package refreshSessions

import (
	"JWTAuth/internal/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (s *RefreshSessionsService) CreateJWT(data *models.CreateJWT) (*models.AccessRefreshTokens, error) {
	claims := jwt.MapClaims{
		"sub":          data.UserUUID,
		"access_level": data.AccessLevel,
		"exp":          time.Now().Add(time.Hour / 2).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(data.Config.JWT.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("error create jwt: %w", err)
	}

	tokens := &models.AccessRefreshTokens{
		AccessToken: token,
	}

	refreshToken, err := s.CreateRefreshSession(&models.AddRefreshSessionRequest{
		UserUUID:  data.UserUUID,
		UserAgent: data.UserAgent,
		Ip:        data.Ip,
	})

	if err != nil {
		return nil, fmt.Errorf("error create jwt: %w", err)
	}

	tokens.RefreshToken = refreshToken

	return tokens, nil
}
