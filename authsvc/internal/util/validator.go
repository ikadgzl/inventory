package util

import (
	"github.com/ikadgzl/inventory/common/proto/auth"
	"github.com/ikadgzl/inventory/common/validator"
)

func ValidateLoginReqest(req *auth.LoginRequest) error {
	err := validator.ValidateEmail(req.Email)
	if err != nil {
		return err
	}

	err = validator.ValidatePassword(req.Password)
	if err != nil {
		return err
	}

	return nil
}

func ValidateRegisterRequest(req *auth.RegisterRequest) error {
	err := validator.ValidateName(req.Name)
	if err != nil {
		return err
	}

	err = validator.ValidateEmail(req.Email)
	if err != nil {
		return err
	}

	err = validator.ValidatePassword(req.Password)
	if err != nil {
		return err
	}

	return nil
}
