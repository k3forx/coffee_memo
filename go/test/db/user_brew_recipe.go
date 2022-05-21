package db_helper

import (
	"context"
	"testing"
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/userbrewrecipe"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func newUserBrewRecipe(user *ent.User, userCoffeeBean *ent.UserCoffeeBean) *ent.UserBrewRecipe {
	return &ent.UserBrewRecipe{
		Status:           int32(model.BrewRecipeStatusCreated),
		UserID:           user.ID,
		UserCoffeeBeanID: userCoffeeBean.ID,
		CoffeeBeanWeight: 16,
		CoffeeBeanGrind:  model.CoffeeBeanGrindCoarse.String(),
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
	}
}
func InsertUserBrewRecipe(client *ent.Client, userBrewRecipe *ent.UserBrewRecipe) (*ent.UserBrewRecipe, error) {
	ubr, err := client.UserBrewRecipe.Create().
		SetStatus(userBrewRecipe.Status).
		SetUserID(userBrewRecipe.UserID).
		SetUserCoffeeBeanID(userBrewRecipe.UserCoffeeBeanID).
		SetCoffeeBeanWeight(userBrewRecipe.CoffeeBeanWeight).
		SetCoffeeBeanGrind(userBrewRecipe.CoffeeBeanGrind).
		SetLiquidWeight(userBrewRecipe.LiquidWeight).
		SetTemperature(userBrewRecipe.Temperature).
		SetStepOne(userBrewRecipe.StepOne).
		SetStepTwo(userBrewRecipe.StepTwo).
		SetStepThree(userBrewRecipe.StepThree).
		SetStepFour(userBrewRecipe.StepFour).
		SetStepFive(userBrewRecipe.StepFive).
		SetMemo(userBrewRecipe.Memo).
		SetCreatedAt(userBrewRecipe.CreatedAt).
		SetUpdatedAt(userBrewRecipe.UpdatedAt).
		SetDeletedAt(userBrewRecipe.DeletedAt).
		Save(context.Background())
	return ubr, err
}

func DeleteUserBrewRecipe(ctx context.Context, client *ent.Client, userBrewRecipe *ent.UserBrewRecipe) {
	client.UserBrewRecipe.Delete().Where(userbrewrecipe.IDEQ(userBrewRecipe.ID))
}

func InsertAndDeleteUserBrewRecipe(tb testing.TB, client *ent.Client, user *ent.User, userCoffeeBean *ent.UserCoffeeBean, setters ...func(ubr *ent.UserBrewRecipe)) *ent.UserBrewRecipe {
	tb.Helper()

	recipe := newUserBrewRecipe(user, userCoffeeBean)
	for _, setter := range setters {
		setter(recipe)
	}

	ubr, err := InsertUserBrewRecipe(client, recipe)
	if err != nil {
		tb.Fatalf("InsertUserBrewRecipe failed: %+v\n", err)
	}
	tb.Cleanup(func() {
		DeleteUserBrewRecipe(context.Background(), client, ubr)
	})

	return ubr
}
