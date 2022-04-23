package auth

import "github.com/k3forx/coffee_memo/pkg/model"

type SignUpOutput struct {
	User model.User
}
type LogInOutput struct {
	User model.User
}
