package refreshSessions

import (
	"JWTAuth/internal/models"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (s *RefreshSessionsService) CreateRefreshSession(refreshSessionReqData *models.AddRefreshSessionRequest) (string, error) {
	refreshToken := uuid.NewString()
	err := s.repo.DeleteRefreshSession(refreshSessionReqData.UserUUID)
	if err != nil {
		return "", fmt.Errorf("error deleting existing refresh session: %w", err)
	}

	if len(refreshSessionReqData.UserAgent) > 256 {
		refreshSessionReqData.UserAgent = refreshSessionReqData.UserAgent[:256]
	}

	refreshSession := &models.RefreshSession{
		UserUUID:     refreshSessionReqData.UserUUID,
		RefreshToken: refreshToken,
		UserAgent:    refreshSessionReqData.UserAgent,
		Fingerprint:  "", // TODO: Figure out how to form fingerprint
		Ip:           refreshSessionReqData.Ip,
		ExpiresIn:    time.Now().Add(time.Hour * 24 * 30).Unix(),
		CreatedAt:    time.Now(),
	}

	err = s.repo.RefreshSessionsRepository.CreateRefreshSession(refreshSession)

	if err != nil {
		return "", fmt.Errorf("error create refresh session: %w", err)
	}

	return refreshToken, nil
}
