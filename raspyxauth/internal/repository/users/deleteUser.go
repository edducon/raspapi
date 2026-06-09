package users

import (
	"JWTAuth/internal/repository/constRepository"
	"context"
	"fmt"
)

func (r *UsersRepository) DeleteUser(userUUID string) error {
	ctx := context.Background()
	dataDelete, errDelete := r.Pool.Exec(ctx, fmt.Sprintf("DELETE FROM %s.%s WHERE uuid=$1",
		constRepository.AUTH_SCHEMA,
		constRepository.USERS_TABLE),
		userUUID,
	)

	if errDelete != nil {
		return errDelete
	}

	rowsAffected := dataDelete.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("nonexisting row")
	}

	return nil
}
