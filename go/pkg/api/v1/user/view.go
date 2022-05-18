package user

import "github.com/k3forx/coffee_memo/pkg/usecase/user"

type GetByIDView struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func newGetByIDView(out *user.GetByIDOutput) GetByIDView {
	u := out.User
	return GetByIDView{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}
