package users

import (
	"JWTAuth/internal/service/errService"
	"fmt"
	"strings"
)

func (s *UsersService) DeleteUser(userUUID string) error {
	err := s.repo.UsersRepository.DeleteUser(userUUID)

	if err != nil {
		if errService.IsForeignKeyError(err) {
			return fmt.Errorf("cannot delete user: it is being referenced by other records: %w", err)
		}
		if strings.Contains(err.Error(), "nonexisting row") {
			return fmt.Errorf("user not found: %w", err)
		}
		return err
	}

	return nil
}
