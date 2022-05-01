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

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=coffee_bean
type Usecase interface {
	GetAllByUserID(ctx context.Context, in GetAllByUserIDInput) (*GetAllByUserIDOutput, *result.Result)
	Create(ctx context.Context, in CreateInput) *result.Result
}

type CoffeeBeanUsecase struct {
	injector inject.Injector
}

func (u *CoffeeBeanUsecase) GetAllByUserID(ctx context.Context, in GetAllByUserIDInput) (*GetAllByUserIDOutput, *result.Result) {
	user, err := u.injector.Reader.User.GetByID(ctx, in.UserID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, result.Error()
	}
	if !user.Exists() {
		return nil, result.New(result.CodeNotFound, "アカウントが存在しません。")
	}

	coffeeBeans, err := u.injector.Reader.CoffeeBean.GetAllByUserID(ctx, user.ID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, result.Error()
	}

	return &GetAllByUserIDOutput{CoffeeBeans: coffeeBeans}, result.OK()
}

func (u *CoffeeBeanUsecase) Create(ctx context.Context, in CreateInput) *result.Result {
	if !in.RoastDegree.Valid() || in.Name == "" {
		return result.New(result.CodeBadRequest, "無効なデータです。")
	}

	user, err := u.injector.Reader.User.GetByID(ctx, in.UserId)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if !user.Exists() {
		return result.New(result.CodeNotFound, "アカウントが見つかりません。")
	}

	coffeeBean := model.CoffeeBean{
		Name:        in.Name,
		FarmName:    in.FarmName,
		Country:     in.Country,
		RoastDegree: in.RoastDegree,
		RoastedAt:   in.RoastedAt,
	}
	if err := u.injector.Writer.CoffeeBean.Create(ctx, &coffeeBean, &user); err != nil {
		logger.Error(ctx, err)
		return result.New(result.CodeInternalError, "コービー豆の登録に失敗しました。")
	}

	return result.OK()
}
