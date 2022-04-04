package entity

import (
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func NewUser(m model.User) *ent.User {
	return &ent.User{
		ID:        int32(m.ID),
		Username:  m.Username,
		Email:     m.Email,
		Password:  m.Password,
		Flags:     m.Flags.Int(),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
