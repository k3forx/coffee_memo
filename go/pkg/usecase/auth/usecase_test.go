package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/model"
	readerMock "github.com/k3forx/coffee_memo/pkg/reader/mock"
	"github.com/k3forx/coffee_memo/pkg/result"
	"github.com/k3forx/coffee_memo/pkg/usecase/auth"
	writerMock "github.com/k3forx/coffee_memo/pkg/writer/mock"
)

func TestUsecase_SignUp(t *testing.T) {
	t.Parallel()

	const email = "test@test.jp"

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    auth.SignUpInput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByEmail(gomock.Any(), email).
					Return(model.User{}, nil)

				userWriter := injector.Writer.User.(*writerMock.MockUser)
				userWriter.EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(&model.User{})).
					DoAndReturn(func(_ context.Context, u *model.User) error {
						u.ID = 1
						u.Email = email
						return nil
					})

				return injector
			},
			in: auth.SignUpInput{
				Email:    email,
				Password: "testtesttest",
				Username: "test user",
			},
			res: result.OK(),
		},
		"validation_failed_by_empty_username": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)
				return injector
			},
			in: auth.SignUpInput{
				Email:    email,
				Password: "testtesttest",
			},
			res: result.New(result.CodeBadRequest, "Username: ユーザー名を指定してください."),
		},
		"error_in_getting_email": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByEmail(gomock.Any(), email).
					Return(model.User{}, errors.New("server error"))

				return injector
			},
			in: auth.SignUpInput{
				Email:    email,
				Password: "testtesttest",
				Username: "test user",
			},
			res: result.Error(),
		},
		"user_is_already_registered": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByEmail(gomock.Any(), email).
					Return(model.User{ID: 1}, nil)

				return injector
			},
			in: auth.SignUpInput{
				Email:    email,
				Password: "testtesttest",
				Username: "test user",
			},
			res: result.New(result.CodeForbidden, "既に使用されているメールアドレスです"),
		},
		"error_in_creating_user": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByEmail(gomock.Any(), email).
					Return(model.User{}, nil)

				userWriter := injector.Writer.User.(*writerMock.MockUser)
				userWriter.EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(&model.User{})).Return(errors.New("server error"))

				return injector
			},
			in: auth.SignUpInput{
				Email:    email,
				Password: "testtesttest",
				Username: "test user",
			},
			res: result.Error(),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			u := auth.NewUsecase(c.setup(ctrl))

			res := u.SignUp(context.Background(), c.in)

			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("SignUp() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
