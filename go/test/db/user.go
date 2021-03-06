package db_helper

import (
	"context"
	"testing"
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	entUser "github.com/k3forx/coffee_memo/pkg/ent/user"
)

func NewUser() *ent.User {
	return &ent.User{
		Username:  "test-user",
		Email:     "test-email",
		Password:  "test-password",
		Flags:     0,
		CreatedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC),
	}
}

func InsertUser(client *ent.Client, user *ent.User) (*ent.User, error) {
	u, err := client.User.Create().
		SetUsername(user.Username).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetFlags(user.Flags).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		Save(context.Background())
	return u, err
}

func DeleteUser(ctx context.Context, client *ent.Client, user *ent.User) {
	client.User.Delete().Where(entUser.IDEQ(user.ID))
}

func InsertAndDeleteUsers(tb testing.TB, client *ent.Client, setters ...func(u *ent.User)) *ent.User {
	tb.Helper()

	user := NewUser()
	for _, set := range setters {
		set(user)
	}

	u, err := InsertUser(client, user)
	if err != nil {
		tb.Fatalf("InsertAndDeleteUsers failed: %+v\n", err)
	}
	tb.Cleanup(func() {
		DeleteUser(context.Background(), client, u)
	})

	return u
}
