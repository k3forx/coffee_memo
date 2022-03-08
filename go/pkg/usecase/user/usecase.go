package user

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/result"
)

func NewUsecase(injector inject.Injector) *UserUsecase {
	return &UserUsecase{
		injector: injector,
	}
}

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=user
type Usecase interface {
	GetByID(ctx context.Context, in GetByIDInput) (*GetByIDOutput, *result.Result)
}

type UserUsecase struct {
	injector inject.Injector
}

var _ Usecase = (*UserUsecase)(nil)

func (u *UserUsecase) GetByID(ctx context.Context, in GetByIDInput) (*GetByIDOutput, *result.Result) {
	user, err := u.injector.Reader.User.GetByID(ctx, in.UserID)
	if err != nil {
		return nil, result.Error()
	}
	if !user.Exists() {
		return nil, result.New(result.CodeNotFound, "user is not found")
	}

	return &GetByIDOutput{User: user}, result.OK()
}
