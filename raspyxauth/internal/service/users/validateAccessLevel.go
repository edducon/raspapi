package users

import "JWTAuth/internal/models"

func (s *UsersService) ValidateAccessLevel(user *models.User) bool {
	return user.AccessLevel >= 0 && user.AccessLevel < 100
}
