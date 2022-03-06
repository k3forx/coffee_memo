package model

import (
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
)

func NewUser(e *ent.User) User {
	return User{
		ID: int(e.ID),
	}
}

type User struct {
	ID             int
	Username       string
	Email          string
	Password       string
	LastLoggedInAt time.Time
}

func (u *User) Exists() bool {
	return u.ID > 0
}