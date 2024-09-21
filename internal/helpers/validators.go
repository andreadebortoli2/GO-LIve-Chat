package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
)

func Required(fields map[string]string) error {
	for k, f := range fields {
		if strings.TrimSpace(f) == "" {
			return fmt.Errorf("%s field cannot be blank", k)
		}
	}
	return nil
}

func LoginValidation(fields map[string]string) error {
	err := Required(fields)
	if err != nil {
		return err
	}
	if !govalidator.IsEmail(fields["email"]) {
		return errors.New("email is not correct format")
	}
	return nil
}
func NewUserValidation(fields map[string]string) error {
	err := Required(fields)
	if err != nil {
		return err
	}
	if len(strings.TrimSpace(fields["user_name"])) < 5 {
		return errors.New("user name must be long at least 5 characters")
	}
	if len(strings.TrimSpace(fields["user_name"])) > 16 {
		return errors.New("user name must be less than 16 characters")
	}
	if !govalidator.IsEmail(fields["email"]) {
		return errors.New("email is not correct format")
	}
	if len(strings.TrimSpace(fields["password"])) < 5 {
		return errors.New("password must be long at least 5 characters")
	}
	if len(strings.TrimSpace(fields["password"])) > 16 {
		return errors.New("password must be less than 16 characters")
	}
	return nil
}
