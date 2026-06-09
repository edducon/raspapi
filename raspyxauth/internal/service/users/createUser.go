package users

import (
	"JWTAuth/internal/models"
	"JWTAuth/internal/service/errService"
	"fmt"
	"github.com/google/uuid"
)

func (s *UsersService) CreateUser(userReqData *models.AddUserRequest) (string, error) {
	if userReqData.UUID == "" {
		userReqData.UUID = uuid.NewString()
	}

	err := s.validateUsername(userReqData.Username)
	if err != nil {
		return "", fmt.Errorf("failed to create user: failed to validate username: %w", err)
	}

	err = s.validatePassword(userReqData.Password)
	if err != nil {
		return "", fmt.Errorf("failed to create user: failed to validate password: %w", err)
	}

	passwordHash, err := s.GeneratePasswordHash(userReqData.Password)
	if err != nil {
		return "", fmt.Errorf("failed to create user: failed to generate password hash: %w", err)
	}

	user := &models.User{
		UUID:         userReqData.UUID,
		Username:     userReqData.Username,
		PasswordHash: passwordHash,
		AccessLevel:  userReqData.AccessLevel,
	}

	errCreateUser := s.repo.UsersRepository.CreateUser(user)

	if errCreateUser != nil {
		if errService.IsDuplicateError(errCreateUser) {
			return "", fmt.Errorf("failed to create user: user already exists")
		}
		return "", fmt.Errorf("failed to create user: %w", errCreateUser)
	}

	return user.UUID, nil
}
