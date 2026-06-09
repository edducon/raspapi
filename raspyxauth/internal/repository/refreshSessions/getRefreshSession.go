package refreshSessions

import (
	"JWTAuth/internal/models"
	"JWTAuth/internal/repository/constRepository"
	"context"
	"fmt"
)

func (r *RefreshSessionsRepository) GetRefreshSessionByToken(token string) (*models.RefreshSession, error) {
	query := fmt.Sprintf("SELECT * FROM %s.%s WHERE refresh_token=$1",
		constRepository.AUTH_SCHEMA,
		constRepository.REFRESH_SESSION_TABLE,
	)

	row := r.Pool.QueryRow(context.Background(), query, token)

	var refreshSession models.RefreshSession
	var id int
	err := row.Scan(
		&id,
		&refreshSession.UserUUID,
		&refreshSession.RefreshToken,
		&refreshSession.UserAgent,
		&refreshSession.Fingerprint,
		&refreshSession.Ip,
		&refreshSession.ExpiresIn,
		&refreshSession.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &refreshSession, nil
}
