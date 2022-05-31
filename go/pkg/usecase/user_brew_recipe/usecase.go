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
	GetByID(ctx context.Context, in GetByIDInput) (*GetByIDOutput, *result.Result)
	DeleteByID(ctx context.Context, in DeleteByIDInput) *result.Result
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
		Status:           model.BrewRecipeStatusCreated,
		User:             user,
		UserCoffeeBean:   userCoffeeBean,
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

func (u *UserBrewRecipeUsecase) GetByID(ctx context.Context, in GetByIDInput) (*GetByIDOutput, *result.Result) {
	user, err := u.injector.Reader.User.GetByID(ctx, in.UserID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, result.Error()
	}
	if !user.Exists() {
		return nil, result.New(result.CodeNotFound, "アカウントが存在しません")
	}

	userBrewRecipe, err := u.injector.Reader.UserBrewRecipe.GetByID(ctx, in.UserBrewRecipeID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, result.Error()
	}
	if !userBrewRecipe.Exists() {
		return nil, result.New(result.CodeNotFound, "ドリップレシピが見つかりません")
	}
	if userBrewRecipe.User.ID != user.ID {
		return nil, result.New(result.CodeForbidden, result.CodeForbidden.String())
	}

	return &GetByIDOutput{UserBrewRecipe: userBrewRecipe}, result.OK()
}

func (u *UserBrewRecipeUsecase) DeleteByID(ctx context.Context, in DeleteByIDInput) *result.Result {
	user, err := u.injector.Reader.User.GetByID(ctx, in.UserID)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if !user.Exists() {
		return result.New(result.CodeNotFound, "アカウントが存在しません")
	}

	userBrewRecipe, err := u.injector.Reader.UserBrewRecipe.GetByID(ctx, in.UserBrewRecipeID)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if !userBrewRecipe.Exists() {
		return result.New(result.CodeNotFound, "ドリップレシピが見つかりません")
	}
	if userBrewRecipe.User.ID != user.ID {
		return result.New(result.CodeForbidden, result.CodeForbidden.String())
	}

	if err := u.injector.Writer.UserBrewRecipe.Delete(ctx, &userBrewRecipe); err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}

	return result.OK()
}
