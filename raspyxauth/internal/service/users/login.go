package users

import (
	"JWTAuth/internal/models"
	"fmt"
)

func (s *UsersService) Login(userReqData *models.LoginData) (*models.User, error) {
	users, err := s.repo.UsersRepository.GetUsersByUsername(userReqData.Username)
	if err != nil {
		return nil, err
	}

	var user *models.User
	for _, u := range *users {
		if u.Username == userReqData.Username {
			user = &u
		}
	}

	if user == nil {
		return nil, fmt.Errorf("user %s not found", userReqData.Username)
	}

	if !s.CheckPassword(userReqData.Password, user.PasswordHash) {
		return nil, fmt.Errorf("invalid credentials")
	}

	user.PasswordHash = ""

	return user, nil
}
