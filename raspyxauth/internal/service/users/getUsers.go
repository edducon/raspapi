package users

import (
	"JWTAuth/internal/models"
)

func (s *UsersService) GetUsers() (*[]models.GetUserResponse, error) {
	users, err := s.repo.UsersRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	var userResp []models.GetUserResponse

	for _, user := range *users {
		resp := models.GetUserResponse{
			UUID:        user.UUID,
			Username:    user.Username,
			AccessLevel: user.AccessLevel,
		}

		userResp = append(userResp, resp)
	}

	return &userResp, nil
}

func (s *UsersService) GetUserByUUID(userUUID string) (*models.GetUserResponse, error) {
	user, err := s.repo.UsersRepository.GetUserByUUID(userUUID)
	if err != nil {
		return nil, err
	}

	userResp := &models.GetUserResponse{
		UUID:        user.UUID,
		Username:    user.Username,
		AccessLevel: user.AccessLevel,
	}

	return userResp, nil
}

func (s *UsersService) GetUsersByUsername(username string) (*[]models.GetUserResponse, error) {
	users, err := s.repo.UsersRepository.GetUsersByUsername(username)
	if err != nil {
		return nil, err
	}

	var userResp []models.GetUserResponse

	for _, user := range *users {
		resp := models.GetUserResponse{
			UUID:        user.UUID,
			Username:    user.Username,
			AccessLevel: user.AccessLevel,
		}

		userResp = append(userResp, resp)
	}

	return &userResp, nil
}

func (s *UsersService) GetUsersByAccessLevel(userAccessLevel int) (*[]models.GetUserResponse, error) {
	users, err := s.repo.UsersRepository.GetUsersByAccessLevel(userAccessLevel)
	if err != nil {
		return nil, err
	}

	var userResp []models.GetUserResponse

	for _, user := range *users {
		resp := models.GetUserResponse{
			UUID:        user.UUID,
			Username:    user.Username,
			AccessLevel: user.AccessLevel,
		}

		userResp = append(userResp, resp)
	}

	return &userResp, nil
}
