package model_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func TestUserFlags_Int(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		flags    model.UserFlags
		expected int
	}{
		"success": {
			flags:    []model.UserFlag{model.UserFlagEmailActivated},
			expected: 2,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.flags.Int()); diff != "" {
				t.Errorf("UserFlags.Int() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		e        *ent.User
		expected model.User
	}{
		"success": {
			e: &ent.User{
				ID:       1,
				Username: "username",
				Email:    "test-email",
				Password: "test-pass",
			},
			expected: model.User{
				ID:       1,
				Username: "username",
				Email:    "test-email",
				Password: "test-pass",
				Flags:    model.UserFlags{0},
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, model.NewUser(c.e)); diff != "" {
				t.Errorf("NewUser() (-want +got):\n%s", diff)
			}
		})
	}
}
func TestNewSignUpUser(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		usename  string
		email    string
		expected model.User
	}{
		"success": {
			usename: "test username",
			email:   "test email",
			expected: model.User{
				Username: "test username",
				Email:    "test email",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, model.NewSignUpUser(c.usename, c.email)); diff != "" {
				t.Errorf("NewSignUpUser mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUser_Exists(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		m        model.User
		expected bool
	}{
		"true": {
			m: model.User{
				ID: 1,
			},
			expected: true,
		},
		"false": {
			m:        model.User{},
			expected: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.m.Exists()); diff != "" {
				t.Errorf("User.Exists() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUser_HasCoffeeBean(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		user           model.User
		userCoffeeBean model.UserCoffeeBean
		expected       bool
	}{
		"false": {
			user: model.User{
				ID: 1,
			},
			userCoffeeBean: model.UserCoffeeBean{
				User: model.User{
					ID: 2,
				},
			},
			expected: false,
		},
		"true": {
			user: model.User{
				ID: 1,
			},
			userCoffeeBean: model.UserCoffeeBean{
				User: model.User{
					ID: 1,
				},
			},
			expected: true,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.user.HasCoffeeBean(c.userCoffeeBean)); diff != "" {
				t.Errorf("User.HasCoffeeBean() mismatch (-wan +got):\n%s", diff)
			}
		})
	}
}
