package model_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/model"
)

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
