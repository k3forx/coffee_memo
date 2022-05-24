package reader_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/reader"
	db_helper "github.com/k3forx/coffee_memo/test/db"
)

func TestUserBrewRecipeReader_GetByID(t *testing.T) {
	t.Parallel()

	impl := reader.NewUserBrewRecipeReader(testClient)

	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "user-brew-recipe-get-by-id"
	})
	userCoffeeBean := db_helper.InsertAndDeleteUserCoffeeBean(t, testClient, user)
	userBrewRecipe := db_helper.InsertAndDeleteUserBrewRecipe(t, testClient, user, userCoffeeBean)

	cases := map[string]struct {
		userBrewRecipeID int
		expected         model.UserBrewRecipe
	}{
		"has_row": {
			userBrewRecipeID: int(userBrewRecipe.ID),
			expected: model.UserBrewRecipe{
				ID:     int(userBrewRecipe.ID),
				Status: model.BrewRecipeStatusCreated,
				User: model.User{
					ID: int(user.ID),
				},
				UserCoffeeBean: model.UserCoffeeBean{
					ID: int(userCoffeeBean.ID),
				},
				CoffeeBeanWeight: 16,
				CoffeeBeanGrind:  model.CoffeeBeanGrindCoarse,
				LiquidWeight:     240,
				Temperature:      92,
				StepOne:          "step 1",
				StepTwo:          "step 2",
				StepThree:        "",
				StepFour:         "",
				StepFive:         "",
				Memo:             "good taste!",
				CreatedAt:        time.Date(2022, time.March, 3, 0, 0, 0, 0, time.UTC),
				UpdatedAt:        time.Date(2022, time.March, 3, 0, 0, 0, 0, time.UTC),
				DeletedAt:        time.Time{},
			},
		},
		"no_row": {
			userBrewRecipeID: -100,
			expected:         model.UserBrewRecipe{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := impl.GetByID(context.Background(), c.userBrewRecipeID)
			if err != nil {
				t.Errorf("err should be nil, but got: %q", err)
			}
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("UserBrewRecipeReader.GetByID() return mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
