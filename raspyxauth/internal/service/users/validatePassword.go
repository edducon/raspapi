package users

import (
	"fmt"
	"unicode"
)

func (s *UsersService) validatePassword(password string) error {
	//re := regexpMustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?!.*\s).{8,256}$`)
	err := fmt.Errorf("password must consist of at least 8 " +
		"characters and have at least one lowercase letter, one " +
		"uppercase letter and one digit, and have no spaces")

	if len(password) < 8 || len(password) > 256 {
		return err
	}

	var (
		hasDigit = false
		hasLower = false
		hasUpper = false
		hasSpace = false
	)

	for _, char := range password {
		switch {
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsSpace(char):
			hasSpace = true
		}
	}

	if !hasDigit || !hasLower || !hasUpper || hasSpace {
		return err
	}

	return nil
}
