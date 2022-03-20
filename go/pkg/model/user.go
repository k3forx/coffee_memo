package model

import (
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
)

type UserFlag int

const (
	UserFlagEmailActivated UserFlag = 2
)

type UserFlags []UserFlag

func (uf UserFlags) Int() int {
	var num int
	for _, flag := range uf {
		num = num + int(flag)
	}
	return num
}

func NewUser(e *ent.User) User {
	return User{
		ID:        int(e.ID),
		Username:  e.Username,
		Email:     e.Email,
		Password:  e.Password,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

type User struct {
	ID             int
	Username       string
	Email          string
	Password       string
	Flags          UserFlags
	LastLoggedInAt time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (u *User) Exists() bool {
	return u.ID > 0
}
