package model_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func TestRoastDegree_String(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		m        model.RoastDegree
		expected string
	}{
		"unknown": {
			m:        model.RoastDegreeUnknown,
			expected: "",
		},
		"light": {
			m:        model.RoastDegreeLight,
			expected: "light",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.m.String()); diff != "" {
				t.Errorf("RoastDegree.String() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRoastDegree_LocalizedString(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		m        model.RoastDegree
		expected string
	}{
		"unknown": {
			m:        model.RoastDegreeUnknown,
			expected: "",
		},
		"light": {
			m:        model.RoastDegreeLight,
			expected: "ライトロースト",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.m.LocalizedString()); diff != "" {
				t.Errorf("RoastDegree.LocalizedString() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRoastDegree_Valid(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		m        model.RoastDegree
		expected bool
	}{
		"valid": {
			m:        model.RoastDegreeChinamon,
			expected: true,
		},
		"invalid": {
			m:        model.RoastDegreeUnknown,
			expected: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.m.Valid()); diff != "" {
				t.Errorf("RoastDegree.Valid() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNewRoastDegree(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		str      string
		expected model.RoastDegree
	}{
		"unknown": {
			str:      "",
			expected: model.RoastDegreeUnknown,
		},
		"light": {
			str:      "light",
			expected: model.RoastDegreeLight,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, model.NewRoastDegree(c.str)); diff != "" {
				t.Errorf("NewRoastDegree() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCoffeeBean_Exists(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		m        model.CoffeeBean
		expected bool
	}{
		"exist": {
			m: model.CoffeeBean{
				ID: 1,
			},
			expected: true,
		},
		"not_exist": {
			m:        model.CoffeeBean{},
			expected: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.m.Exists()); diff != "" {
				t.Errorf("CoffeeBean.Exists() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
