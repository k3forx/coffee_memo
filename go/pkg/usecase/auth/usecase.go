package auth

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/result"
	"golang.org/x/crypto/bcrypt"
)

func NewUsecase(injector inject.Injector) *AuthUsecase {
	return &AuthUsecase{
		injector: injector,
	}
}

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=auth
type Usecase interface {
}

type AuthUsecase struct {
	injector inject.Injector
}

var _ Usecase = (*AuthUsecase)(nil)

func (u *AuthUsecase) SignUp(ctx context.Context, in SignUpInput) (*SignUpOutput, *result.Result) {
	if err := in.Validate(); err != nil {
		// TODO: add error log
		// logger.Error(ctx, err)
		return nil, result.New(result.CodeBadRequest, err.Error())
	}

	user := model.NewSignUpUser(in.Username, in.Email)

	existingUser, err := u.injector.Reader.User.GetByEmail(ctx, in.Email)
	if err != nil {
		// TODO: add error log
		// logger.Error(ctx, err)
		return nil, result.Error()
	}
	if existingUser.Exists() {
		return nil, result.New(result.CodeForbidden, "既に使用されているメールアドレスです")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		// TODO: add error log
		// logger.Error(ctx, err)
		return nil, result.Error()
	}
	user.Password = string(hashedPassword)

	if err := u.injector.Writer.User.Create(ctx, &user); err != nil {
		// TODO: add error log
		// logger.Error(ctx, err)
		return nil, result.Error()
	}

	return &SignUpOutput{User: user}, result.OK()
}
