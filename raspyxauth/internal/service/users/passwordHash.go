package users

import (
	"crypto/sha256"
	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) GeneratePasswordHash(password string) (string, error) {
	h := sha256.Sum256([]byte(password))
	hash, err := bcrypt.GenerateFromPassword(h[:], bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
