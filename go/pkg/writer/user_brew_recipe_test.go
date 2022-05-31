package writer_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/writer"
	dbHelper "github.com/k3forx/coffee_memo/test/db"
)

func TestUserBrewRecipeWriter_Create(t *testing.T) {
	impl := writer.NewUserBrewRecipeWriter(testClient)

	user := dbHelper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "test-user-brew-recipe-writer-create"
	})
	userCoffeeBean := dbHelper.InsertAndDeleteUserCoffeeBean(t, testClient, user, func(cb *ent.UserCoffeeBean) {
		cb.UserID = user.ID
	})
	now := time.Now().In(time.UTC)

	cases := map[string]struct {
		m           model.UserBrewRecipe
		expectedErr error
	}{
		"success": {
			m: model.UserBrewRecipe{
				Status: model.BrewRecipeStatusCreated,
				User: model.User{
					ID: int(user.ID),
				},
				UserCoffeeBean: model.UserCoffeeBean{
					ID: int(userCoffeeBean.ID),
				},
				CoffeeBeanWeight: 15,
				CoffeeBeanGrind:  model.CoffeeBeanGrindCoarse,
				LiquidWeight:     225,
				Temperature:      90,
				StepOne:          "step 1",
				StepTwo:          "step 2",
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expectedErr: nil,
		},
		// "error": {
		// 	m: &model.UserBrewRecipe{
		// 		Status: model.BrewRecipeStatusCreated,
		// 		User: model.User{
		// 			ID: 0,
		// 		},
		// 		UserCoffeeBean: model.UserCoffeeBean{
		// 			ID: 0,
		// 		},
		// 		CoffeeBeanWeight: 15,
		// 		CoffeeBeanGrind:  model.CoffeeBeanGrindCoarse,
		// 		LiquidWeight:     225,
		// 		Temperature:      90,
		// 		StepOne:          "step 1",
		// 		StepTwo:          "step 2",
		// 		CreatedAt:        now,
		// 		UpdatedAt:        now,
		// 	},
		// 	expectedErr: errors.New("ent: constraint failed: Error 1452: Cannot add or update a child row: a foreign key constraint fails (`coffee_app_test`.`user_brew_recipes`, CONSTRAINT `fkey_user_drip_recipes_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))"),
		// },
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			if err := impl.Create(ctx, &c.m); err != nil {
				t.Errorf("err should be nil but, got %q", err)
			}

			userBrewRecipeEntity, err := testClient.UserBrewRecipe.Get(ctx, int32(c.m.ID))
			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}

			actualUserBrewRecipe := model.NewUserBrewRecipe(userBrewRecipeEntity)
			opts := cmp.Options{
				cmpopts.EquateApproxTime(time.Minute),
			}
			if diff := cmp.Diff(c.m, actualUserBrewRecipe, opts); diff != "" {
				t.Errorf("user brew recipe mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUserBrewRecipeWriter_Delete(t *testing.T) {
	impl := writer.NewUserBrewRecipeWriter(testClient)

	user := dbHelper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "test-user-brew-recipe-writer-create"
	})
	userCoffeeBean := dbHelper.InsertAndDeleteUserCoffeeBean(t, testClient, user, func(cb *ent.UserCoffeeBean) {
		cb.UserID = user.ID
	})
	now := time.Now().In(time.UTC)
	userBrewRecipe := dbHelper.InsertAndDeleteUserBrewRecipe(t, testClient, user, userCoffeeBean, func(ubr *ent.UserBrewRecipe) {
		ubr.CreatedAt = now
		ubr.UpdatedAt = now
	})

	cases := map[string]struct {
		m        model.UserBrewRecipe
		expected model.UserBrewRecipe
	}{
		"success": {
			m: model.NewUserBrewRecipe(userBrewRecipe),
			expected: model.UserBrewRecipe{
				ID:     int(userBrewRecipe.ID),
				Status: model.BrewRecipeStatusDeletedByUser,
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
				Memo:             "good taste!",
				DeletedAt:        now,
				CreatedAt:        now,
				UpdatedAt:        now,
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()

			if err := impl.Delete(ctx, &c.m); err != nil {
				t.Errorf("err should be nil but, got %q", err)
			}

			userBrewRecipeEntity, err := testClient.UserBrewRecipe.Get(ctx, int32(c.m.ID))
			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}

			actualUserBrewRecipe := model.NewUserBrewRecipe(userBrewRecipeEntity)
			opts := cmp.Options{
				cmpopts.EquateApproxTime(time.Minute),
			}
			if diff := cmp.Diff(c.expected, actualUserBrewRecipe, opts); diff != "" {
				t.Errorf("user brew recipe mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
