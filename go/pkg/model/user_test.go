package model_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
)

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
