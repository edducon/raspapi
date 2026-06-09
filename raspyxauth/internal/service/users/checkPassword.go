package users

import (
	"crypto/sha256"
	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) CheckPassword(password, hash string) bool {
	h := sha256.Sum256([]byte(password))
	err := bcrypt.CompareHashAndPassword([]byte(hash), h[:])
	return err == nil
}
