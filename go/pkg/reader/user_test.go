package reader_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/reader"
	db_helper "github.com/k3forx/coffee_memo/test/db"
)

func TestUser_GetByID(t *testing.T) {
	t.Parallel()

	userReader := reader.NewUserReader(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient)

	cases := map[string]struct {
		userID   int
		expected model.User
	}{
		"has_rows": {
			userID: int(user.ID),
			expected: model.User{
				ID:        int(user.ID),
				Username:  user.Username,
				Email:     user.Email,
				Flags:     model.UserFlags{0},
				Password:  user.Password,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		},
		"no_rows": {
			userID:   -100,
			expected: model.User{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := userReader.GetByID(context.Background(), c.userID)

			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("GetByID() mismatch (-want +got): %s\n", diff)
			}
		})
	}
}

func TestUser_GetByEmail(t *testing.T) {
	t.Parallel()

	userReader := reader.NewUserReader(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "getbyemail"
	})

	cases := map[string]struct {
		email    string
		expected model.User
	}{
		"has_rows": {
			email: "getbyemail",
			expected: model.User{
				ID:        int(user.ID),
				Username:  user.Username,
				Email:     user.Email,
				Password:  user.Password,
				Flags:     model.UserFlags{0},
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		},
		"no_rows": {
			email:    "",
			expected: model.User{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := userReader.GetByEmail(context.Background(), c.email)

			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("GetByEmail() mismatch (-want +got): %s\n", diff)
			}
		})
	}
}
