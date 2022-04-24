package coffee_bean

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/logger"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/result"
)

func NewUsecase(injector inject.Injector) *CoffeeBeanUsecase {
	return &CoffeeBeanUsecase{
		injector: injector,
	}
}

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=auth
type Usecase interface {
	Create(ctx context.Context, in CreateInput) *result.Result
}

type CoffeeBeanUsecase struct {
	injector inject.Injector
}

func (u *CoffeeBeanUsecase) Create(ctx context.Context, in CreateInput) *result.Result {
	coffeeBean := model.CoffeeBean{
		User: model.User{
			ID: in.UserId,
		},
		Name:        in.Name,
		FarmName:    in.FarmName,
		Country:     in.Country,
		RoastDegree: in.RoastDegree,
		RoastedAt:   in.RoastedAt,
	}

	if err := u.injector.Writer.CoffeeBean.Create(ctx, &coffeeBean); err != nil {
		logger.Error(ctx, err)
		return result.New(result.CodeInternalError, "コービー豆の登録に失敗しました。")
	}

	return result.OK()
}
