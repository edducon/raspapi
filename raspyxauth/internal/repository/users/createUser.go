package users

import (
	"JWTAuth/internal/models"
	"JWTAuth/internal/repository/constRepository"
	"context"
	"fmt"
)

func (r *UsersRepository) CreateUser(userData *models.User) error {
	ctx := context.Background()
	_, errInsert := r.Pool.Exec(ctx, fmt.Sprintf("INSERT INTO %s.%s ("+
		"uuid, "+
		"username, "+
		"password_hash, "+
		"access_level) "+
		"VALUES ($1, $2, $3, $4)",
		constRepository.AUTH_SCHEMA,
		constRepository.USERS_TABLE),
		userData.UUID,
		userData.Username,
		userData.PasswordHash,
		userData.AccessLevel,
	)

	if errInsert != nil {
		return errInsert
	}

	return nil
}
