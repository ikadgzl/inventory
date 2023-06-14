package validator

import (
	"fmt"
	"regexp"
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	if ok, _ := regexp.MatchString(`[A-Z]`, password); !ok {
		return fmt.Errorf("password must include capital letters")
	}

	if ok, _ := regexp.MatchString(`[a-z]`, password); !ok {
		return fmt.Errorf("password must include small letters")
	}

	if ok, _ := regexp.MatchString(`[0-9]`, password); !ok {
		return fmt.Errorf("password must include numbers")
	}

	if ok, _ := regexp.MatchString(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`, password); !ok {
		return fmt.Errorf("password must include special characters")
	}

	return nil
}
