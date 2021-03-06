package model

import (
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
)

type UserFlag int

const (
	UserFlagEmailActivated UserFlag = 2
)

var userFlagsList = UserFlags{
	UserFlagEmailActivated,
}

func (uf UserFlag) Num() int {
	return int(uf)
}

type UserFlags []UserFlag

func (uf UserFlags) Int() int {
	var num int
	for _, flag := range uf {
		num = num + int(flag)
	}
	return num
}

func NewUser(e *ent.User) User {
	flags := make([]UserFlag, len(userFlagsList))
	for i, f := range userFlagsList {
		if e.Flags&f.Num() == f.Num() {
			flags[i] = f
		}
	}
	return User{
		ID:        int(e.ID),
		Username:  e.Username,
		Email:     e.Email,
		Flags:     flags,
		Password:  e.Password,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func NewSignUpUser(username, email string) User {
	return User{
		Username: username,
		Email:    email,
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

func (u *User) HasCoffeeBean(cb UserCoffeeBean) bool {
	return u.ID == cb.User.ID
}
