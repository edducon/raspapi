package users

import (
	"JWTAuth/internal/models"
	"JWTAuth/internal/repository/constRepository"
	"context"
	"fmt"
)

func (r *UsersRepository) UpdateUser(userUUID string, userData *models.UpdateUserRequest) error {
	query := fmt.Sprintf("UPDATE %s.%s SET "+
		"username=$1, "+
		"access_level=$2 "+
		"WHERE uuid=$3 "+
		"RETURNING uuid",
		constRepository.AUTH_SCHEMA,
		constRepository.USERS_TABLE)

	var updatedUUID string
	ctx := context.Background()
	row := r.Pool.QueryRow(ctx, query,
		userData.Username,
		userData.AccessLevel,
		userUUID,
	)

	errUpdate := row.Scan(&updatedUUID)

	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
