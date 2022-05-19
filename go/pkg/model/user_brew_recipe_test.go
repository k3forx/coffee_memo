package model_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func TestUserBrewRecipe_Exists(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		m        model.UserBrewRecipe
		expected bool
	}{
		"exists": {
			m: model.UserBrewRecipe{
				ID: 1,
			},
			expected: true,
		},
		"does_not_exist": {
			m: model.UserBrewRecipe{
				ID: 0,
			},
			expected: false,
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, c.m.Exists()); diff != "" {
				t.Errorf("UserBrewRecipe.Exists() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNewUserBrewRecipe(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		e        *ent.UserBrewRecipe
		expected model.UserBrewRecipe
	}{
		"success": {
			e: &ent.UserBrewRecipe{
				ID:               1,
				Status:           1,
				CoffeeBeanWeight: 16,
				CoffeeBeanGrind:  string(model.CoffeeBeanGrindCoarse),
				Edges: ent.UserBrewRecipeEdges{
					UserCoffeeBean: &ent.UserCoffeeBean{
						ID: 3,
					},
				},
				CreatedAt: time.Date(2022, time.March, 2, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2022, time.March, 2, 0, 0, 0, 0, time.UTC),
			},
			expected: model.UserBrewRecipe{
				ID:     1,
				Status: model.BrewRecipeStatusCreated,
				UserCoffeeBean: model.UserCoffeeBean{
					ID: 3,
				},
				CoffeeBeanWeight: 16,
				CoffeeBeanGrind:  "coarse",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(c.expected, model.NewUserBrewRecipe(c.e)); diff != "" {
				t.Errorf("NewUserBrewRecipe mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
