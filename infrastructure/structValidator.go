package infrastructure

import (
	"abduselam-arabianmejlis/domain"

	"github.com/go-playground/validator/v10"
)

func ValidateUser(user *domain.User) error{
	var validate = validator.New()
	err := validate.Struct(user)

	return err

}