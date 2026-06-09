package refreshSessions

import (
	"JWTAuth/internal/repository/constRepository"
	"context"
	"fmt"
)

func (r *RefreshSessionsRepository) DeleteRefreshSession(userUUID string) error {
	ctx := context.Background()
	_, errDelete := r.Pool.Exec(ctx, fmt.Sprintf("DELETE FROM %s.%s WHERE user_uuid=$1",
		constRepository.AUTH_SCHEMA,
		constRepository.REFRESH_SESSION_TABLE),
		userUUID,
	)

	if errDelete != nil {
		return errDelete
	}

	//rowsAffected := dataDelete.RowsAffected()
	//
	//if rowsAffected == 0 {
	//	return fmt.Errorf("nonexisting row")
	//}

	return nil
}
