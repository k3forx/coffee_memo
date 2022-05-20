package user_brew_recipe

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/logger"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/result"
)

func NewUsecase(injector inject.Injector) *UserBrewRecipeUsecase {
	return &UserBrewRecipeUsecase{
		injector: injector,
	}
}

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=user_brew_recipe
type Usecase interface {
	Create(ctx context.Context, in CreateInput) *result.Result
}

type UserBrewRecipeUsecase struct {
	injector inject.Injector
}

func (u *UserBrewRecipeUsecase) Create(ctx context.Context, in CreateInput) *result.Result {
	user, err := u.injector.Reader.User.GetByID(ctx, in.UserID)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if !user.Exists() {
		return result.New(result.CodeNotFound, "アカウントが存在しません")
	}

	userCoffeeBean, err := u.injector.Reader.UserCoffeeBean.GetByIDWithUser(ctx, in.UserCoffeeBeanID)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if !userCoffeeBean.Exists() {
		return result.New(result.CodeNotFound, "コーヒー豆が存在しません")
	}
	if !user.HasCoffeeBean(userCoffeeBean) {
		return result.New(result.CodeForbidden, result.CodeForbidden.String())
	}

	userBrewRecipe := model.UserBrewRecipe{
		Status: model.BrewRecipeStatusCreated,
		User: model.User{
			ID: in.UserID,
		},
		UserCoffeeBean: model.UserCoffeeBean{
			ID: in.UserCoffeeBeanID,
		},
		CoffeeBeanWeight: in.CoffeeBeanWeight,
		CoffeeBeanGrind:  in.CoffeeBeanGrind,
		LiquidWeight:     in.LiquidWeight,
		Temperature:      in.LiquidWeight,
		StepOne:          in.StepOne,
		StepTwo:          in.StepTwo,
		StepThree:        in.StepThree,
		StepFour:         in.StepFour,
		StepFive:         in.StepFive,
		Memo:             in.Memo,
	}
	if err := u.injector.Writer.UserBrewRecipe.Create(ctx, &userBrewRecipe); err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}

	return result.OK()
}
