package auth

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignUpInput struct {
	Username string
	Email    string
	Password string
}

func (in SignUpInput) Validate() error {
	return validation.ValidateStruct(&in,
		validation.Field(&in.Username, validation.Required.Error("ユーザー名を指定してください")),
		validation.Field(&in.Email, validation.Required, is.Email),
		validation.Field(
			&in.Password,
			validation.Required.Error("パスワードは必須項目です"),
			validation.Length(8, 20).Error("パスワードは8-20文字で指定してください"),
		),
	)
}
