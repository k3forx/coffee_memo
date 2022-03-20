package user

import (
	"context"
	"fmt"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/result"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserUsecase) SignUp(ctx context.Context, in SignUpInput) *result.Result {
	user := model.User{
		Username: in.Username,
		Email:    in.Email,
	}

	existingUser, err := u.injector.Reader.User.GetByEmail(ctx, in.Email)
	fmt.Printf("err: %+v\n", err)
	if err != nil {
		return result.Error()
	}
	if existingUser.Exists() {
		return result.New(result.CodeForbidden, "既に使用されているメールアドレスです")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	fmt.Printf("err: %+v\n", err)
	if err != nil {
		return result.Error()
	}
	user.Password = string(hashedPassword)

	if err := u.injector.Writer.User.Create(ctx, user); err != nil {
		fmt.Printf("err: %+v\n", err)
		return result.Error()
	}

	return result.OK()
}
