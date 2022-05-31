package writer

import (
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
	"golang.org/x/net/context"
)

func NewUserBrewRecipeWriter(db *ent.Client) *UserBrewRecipeWriter {
	return &UserBrewRecipeWriter{
		db: db,
	}
}

//go:generate mockgen -source=./user_brew_recipe.go -destination=./mock/user_brew_recipe_mock.go -package=writer
type UserBrewRecipe interface {
	Create(ctx context.Context, userBrewRecipe *model.UserBrewRecipe) error
	Delete(ctx context.Context, userBrewRecipe *model.UserBrewRecipe) error
}

type UserBrewRecipeWriter struct {
	db *ent.Client
}

var _ UserBrewRecipe = (*UserBrewRecipeWriter)(nil)

func (impl *UserBrewRecipeWriter) Create(ctx context.Context, userBrewRecipe *model.UserBrewRecipe) error {
	now := time.Now().In(time.UTC)
	e, err := impl.db.UserBrewRecipe.Create().
		SetStatus(int32(userBrewRecipe.Status)).
		SetUserID(int32(userBrewRecipe.User.ID)).
		SetUserCoffeeBeanID(int32(userBrewRecipe.UserCoffeeBean.ID)).
		SetCoffeeBeanWeight(userBrewRecipe.CoffeeBeanWeight).
		SetCoffeeBeanGrind(userBrewRecipe.CoffeeBeanGrind.String()).
		SetLiquidWeight(userBrewRecipe.LiquidWeight).
		SetTemperature(userBrewRecipe.Temperature).
		SetStepOne(userBrewRecipe.StepOne).
		SetStepTwo(userBrewRecipe.StepTwo).
		SetStepThree(userBrewRecipe.StepThree).
		SetStepFour(userBrewRecipe.StepFour).
		SetStepFive(userBrewRecipe.StepFive).
		SetMemo(userBrewRecipe.Memo).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return err
	}
	*userBrewRecipe = model.NewUserBrewRecipe(e)
	return nil
}

func (impl *UserBrewRecipeWriter) Delete(ctx context.Context, userBrewRecipe *model.UserBrewRecipe) error {
	now := time.Now().In(time.UTC)
	if _, err := impl.db.UserBrewRecipe.UpdateOneID(int32(userBrewRecipe.ID)).
		SetStatus(int32(model.BrewRecipeStatusDeletedByUser)).
		SetDeletedAt(now).
		Save(ctx); err != nil {
		return err
	}
	return nil
}
