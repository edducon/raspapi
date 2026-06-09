package users

import (
	"fmt"
	"regexp"
)

func (s *UsersService) validateUsername(username string) error {
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9-_\.]{3,256}$`)
	matched := re.MatchString(username)
	if !matched {
		return fmt.Errorf("username must be between 3 and " +
			"256 characters long, start with a letter, and contain " +
			"only Latin letters, numbers, and symbols _ - and the period")
	}
	return nil
}
