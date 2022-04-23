package auth

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/logger"
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
	SignUp(ctx context.Context, in SignUpInput) *result.Result
	LogIn(ctx context.Context, in LogInInput) (*LogInOutput, *result.Result)
}

type AuthUsecase struct {
	injector inject.Injector
}

func (u *AuthUsecase) SignUp(ctx context.Context, in SignUpInput) *result.Result {
	if err := in.Validate(); err != nil {
		logger.Error(ctx, err)
		return result.New(result.CodeBadRequest, err.Error())
	}

	user := model.NewSignUpUser(in.Username, in.Email)

	existingUser, err := u.injector.Reader.User.GetByEmail(ctx, in.Email)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	if existingUser.Exists() {
		return result.New(result.CodeForbidden, "既に使用されているメールアドレスです")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}
	user.Password = string(hashedPassword)

	if err := u.injector.Writer.User.Create(ctx, &user); err != nil {
		logger.Error(ctx, err)
		return result.Error()
	}

	return result.OK()
}

func (u *AuthUsecase) LogIn(ctx context.Context, in LogInInput) (*LogInOutput, *result.Result) {
	if err := in.Validate(); err != nil {
		logger.Error(ctx, err)
		return nil, result.New(result.CodeBadRequest, err.Error())
	}

	user, err := u.injector.Reader.User.GetByEmail(ctx, in.Email)
	if err != nil {
		logger.Error(ctx, err)
		return nil, result.Error()
	}
	if !user.Exists() {
		return nil, result.New(result.CodeNotFound, result.CodeNotFound.String())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, result.New(result.CodeBadRequest, "パスワードが違います。")
	}

	return &LogInOutput{User: user}, result.OK()
}
