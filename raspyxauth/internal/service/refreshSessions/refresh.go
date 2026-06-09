package refreshSessions

import (
	"JWTAuth/internal/models"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

func (s *RefreshSessionsService) Refresh(refreshData *models.RefreshData) (*models.AccessRefreshTokens, error) {
	refreshSession, err := s.repo.RefreshSessionsRepository.GetRefreshSessionByToken(refreshData.RefreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("refresh token not found")
		}
		return nil, fmt.Errorf("error get refresh session: %w", err)
	}

	if refreshSession.ExpiresIn < time.Now().Unix() {
		_ = s.repo.DeleteRefreshSession(refreshSession.UserUUID)
		return nil, fmt.Errorf("refresh token is expired")
	}

	user, err := s.repo.UsersRepository.GetUserByUUID(refreshSession.UserUUID)
	if err != nil {
		return nil, fmt.Errorf("error getting user by uuid: %w", err)
	}

	tokens, err := s.CreateJWT(&models.CreateJWT{
		UserUUID:    refreshSession.UserUUID,
		UserAgent:   refreshData.UserAgent,
		Ip:          refreshData.Ip,
		AccessLevel: user.AccessLevel,
		Config:      s.cfg,
	})
	if err != nil {
		return nil, fmt.Errorf("error create jwt: %w", err)
	}

	return tokens, nil
}
