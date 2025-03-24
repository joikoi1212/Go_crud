package validations

import (
	"fmt"
)

func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

func ValidatePassword(password string, confirmPassword string) error {
	if password != confirmPassword {
		return fmt.Errorf("password is required")
	}
	return nil
}
