package users

import (
	"JWTAuth/internal/models"
	"JWTAuth/internal/service/errService"
	"fmt"
)

func (s *UsersService) UpdateUser(userUUID string, userData *models.UpdateUserRequest) error {
	errUpdateUser := s.repo.UsersRepository.UpdateUser(userUUID, userData)

	if errUpdateUser != nil {
		if errService.IsDuplicateError(errUpdateUser) {
			return fmt.Errorf("failed to update user: target user already exists")
		}
		return fmt.Errorf("failed to update user: %w", errUpdateUser)
	}

	return nil
}
