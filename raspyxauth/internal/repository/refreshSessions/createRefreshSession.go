package refreshSessions

import (
	"JWTAuth/internal/models"
	"JWTAuth/internal/repository/constRepository"
	"context"
	"fmt"
)

func (r *RefreshSessionsRepository) CreateRefreshSession(refreshSessionData *models.RefreshSession) error {
	ctx := context.Background()
	_, errInsert := r.Pool.Exec(ctx, fmt.Sprintf("INSERT INTO %s.%s ("+
		"user_uuid, "+
		"refresh_token, "+
		"user_agent, "+
		"fingerprint, "+
		"ip, "+
		"expires_in) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		constRepository.AUTH_SCHEMA,
		constRepository.REFRESH_SESSION_TABLE),
		refreshSessionData.UserUUID,
		refreshSessionData.RefreshToken,
		refreshSessionData.UserAgent,
		refreshSessionData.Fingerprint,
		refreshSessionData.Ip,
		refreshSessionData.ExpiresIn,
	)

	if errInsert != nil {
		return errInsert
	}

	return nil
}
