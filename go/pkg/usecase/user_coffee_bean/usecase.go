package user_coffee_bean

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/logger"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/result"
)

func NewUsecase(injector inject.Injector) *UserCoffeeBeanUsecase {
	return &UserCoffeeBeanUsecase{
		injector: injector,
	}
}

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=user_coffee_bean
type Usecase interface {
	GetAllByUserID(ctx context.Context, in GetAllByUserIDInput) (*GetAllByUserIDOutput, *result.Result)
	GetByID(ctx context.Context, in GetByIDInput) (*GetByIDOutput, *result.Result)
	Create(ctx context.Context, in CreateInput) *result.Result
	EditByID(ctx context.Context, in EditByIDInput) *result.Result
	DeleteByID(ctx context.Context, in DeleteByIDInput) *result.Result
}

type UserCoffeeBeanUsecase struct {
	injector inject.Injector
}

func (u *UserCoffeeBeanUsecase) GetAllByUserID(ctx context.Context, in GetAllByUserIDInput) (*GetAllByUserIDOutput, *result.Result) {
	user, err := u.injector.Reader.User.GetByID(ctx, in.UserID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, result.Error()
	}
	if !user.Exists() {
		return nil, result.New(result.CodeNotFound, "アカウントが存在しません。")
	}

	coffeeBeans, err := u.injector.Reader.UserCoffeeBean.GetAllByUserID(ctx, user.ID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, result.Error()
	}

	return &GetAllByUserIDOutput{CoffeeBeans: coffeeBeans}, result.OK()
}

func (u *UserCoffeeBeanUsecase) GetByID(ctx context.Context, in GetByIDInput) (*GetByIDOutput, *result.Result) {
	userCofffeeBean, err := u.injector.Reader.UserCoffeeBean.GetByID(ctx, in.UserCoffeeBeanID)
	if err != nil {
		logger.Error(ctx, err)
		return nil, result.Error()
	}
	if !userCofffeeBean.Exists() {
		return nil, result.New(result.CodeNotFound, "コーヒー豆が見つかりません")
	}
	return &GetByIDOutput{UserCoffeeBean: userCofffeeBean}, result.OK()
}

func (u *UserCoffeeBeanUsecase) Create(ctx context.Context, in CreateInput) *result.Result {
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

	userCoffeeBean := model.UserCoffeeBean{
		User: model.User{
			ID: user.ID,
		},
		Name:        in.Name,
		FarmName:    in.FarmName,
		Country:     in.Country,
		RoastDegree: in.RoastDegree,
		RoastedAt:   in.RoastedAt,
	}
	if err := u.injector.Writer.UserCoffeeBean.Create(ctx, &userCoffeeBean, &user); err != nil {
		logger.Error(ctx, err)
		return result.New(result.CodeInternalError, "コービー豆の登録に失敗しました。")
	}

	return result.OK()
}

func (u *UserCoffeeBeanUsecase) EditByID(ctx context.Context, in EditByIDInput) *result.Result {
	user, err := u.injector.Reader.User.GetByID(ctx, in.UserID)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if !user.Exists() {
		return result.New(result.CodeNotFound, "アカウントが見つかりません")
	}

	userCoffeeBean, err := u.injector.Reader.UserCoffeeBean.GetByID(ctx, in.UserCoffeeBeanID)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if !userCoffeeBean.Exists() {
		return result.New(result.CodeNotFound, "コーヒー豆が存在しません")
	}
	if userCoffeeBean.User.ID != user.ID {
		return result.New(result.CodeForbidden, "編集できません")
	}

	if err := in.validate(); err != nil {
		return result.New(result.CodeBadRequest, err.Error())
	}

	userCoffeeBean = in.Model()
	if err := u.injector.Writer.UserCoffeeBean.UpdateByID(ctx, &userCoffeeBean); err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}

	return result.OK()
}

func (u *UserCoffeeBeanUsecase) DeleteByID(ctx context.Context, in DeleteByIDInput) *result.Result {
	user, err := u.injector.Reader.User.GetByID(ctx, in.UserID)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if !user.Exists() {
		return result.New(result.CodeNotFound, "アカウントが見つかりません")
	}

	coffeeBean, err := u.injector.Reader.UserCoffeeBean.GetByIDWithUser(ctx, in.CoffeeBeanID)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if !coffeeBean.Exists() {
		return result.New(result.CodeNotFound, "コーヒー豆が存在しません")
	}
	if coffeeBean.User.ID != user.ID {
		return result.New(result.CodeForbidden, "削除できません")
	}

	coffeeBean.Status = model.CoffeeBeanStatusDeleted
	if err := u.injector.Writer.UserCoffeeBean.DeleteByID(ctx, &coffeeBean); err != nil {
		logger.Error(ctx, err)
		return result.New(result.CodeInternalError, "削除に失敗しました")
	}

	return result.OK()
}
