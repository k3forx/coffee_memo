package user

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r SignUpRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required, validation.Length(8, 20), is.Email),
	)
}
