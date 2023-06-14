package validator

import "fmt"

func ValidateName(name string) error {
	if len(name) < 2 {
		return fmt.Errorf("name must be at least 2 characters")
	}

	return nil
}
