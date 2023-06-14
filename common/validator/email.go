package validator

import (
	"fmt"
	"regexp"
)

func ValidateEmail(email string) error {
	r, err := regexp.Compile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	if err != nil {
		return err
	}

	if !r.MatchString(email) {
		return fmt.Errorf("invalid email")
	}

	return nil
}
