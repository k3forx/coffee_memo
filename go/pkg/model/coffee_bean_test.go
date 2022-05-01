package model_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/ent"
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

func TestNewCoffeeBean(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		e        *ent.CoffeeBean
		expected model.CoffeeBean
	}{
		"not_empty": {
			e: &ent.CoffeeBean{
				ID:          int32(1),
				Name:        "name",
				FarmName:    "farm name",
				Country:     "country",
				RoastDegree: "light",
				RoastedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
				CreatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
			},
			expected: model.CoffeeBean{
				ID:          1,
				Name:        "name",
				FarmName:    "farm name",
				Country:     "country",
				RoastDegree: model.RoastDegreeLight,
				RoastedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
				CreatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
			},
		},
		"empty": {
			e:        &ent.CoffeeBean{},
			expected: model.CoffeeBean{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, model.NewCoffeeBean(c.e)); diff != "" {
				t.Errorf("NewCoffeeBean mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNewCoffeeBeanWithUser(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		e        *ent.UsersCoffeeBean
		expected model.CoffeeBean
	}{
		"not_empty": {
			e: &ent.UsersCoffeeBean{
				ID:           int32(1),
				UserID:       3,
				CoffeeBeanID: 12,
				CreatedAt:    time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
				UpdatedAt:    time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
				Edges: ent.UsersCoffeeBeanEdges{
					User: &ent.User{
						ID: 3,
					},
					CoffeeBean: &ent.CoffeeBean{
						ID:          12,
						Name:        "name",
						FarmName:    "farm name",
						Country:     "country",
						RoastDegree: "light",
						RoastedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
						CreatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			expected: model.CoffeeBean{
				ID: 1,
				User: model.User{
					ID:    3,
					Flags: model.UserFlags{0},
				},
				Name:        "name",
				FarmName:    "farm name",
				Country:     "country",
				RoastDegree: model.RoastDegreeLight,
				RoastedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
				CreatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
			},
		},
		"empty": {
			e: &ent.UsersCoffeeBean{
				Edges: ent.UsersCoffeeBeanEdges{
					User:       &ent.User{},
					CoffeeBean: &ent.CoffeeBean{},
				},
			},
			expected: model.CoffeeBean{
				User: model.User{
					Flags: model.UserFlags{0},
				},
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, model.NewCoffeeBeanWithUser(c.e)); diff != "" {
				t.Errorf("NewCoffeeBeanWithUser mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
