package user_test

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
	"github.com/k3forx/coffee_memo/pkg/usecase/user"
)

func TestUsecase_GetByID(t *testing.T) {
	t.Parallel()

	const userID = 1

	returnedUser := model.User{
		ID: userID,
	}

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    user.GetByIDInput
		out   *user.GetByIDOutput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).Return(returnedUser, nil)

				return injector
			},
			in: user.GetByIDInput{
				UserID: userID,
			},
			out: &user.GetByIDOutput{
				User: returnedUser,
			},
			res: result.OK(),
		},
		"error_in_getting_user": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).Return(model.User{}, errors.New("server error"))

				return injector
			},
			in: user.GetByIDInput{
				UserID: userID,
			},
			out: nil,
			res: result.Error(),
		},
		"user_is_not_found": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).Return(model.User{}, nil)

				return injector
			},
			in: user.GetByIDInput{
				UserID: userID,
			},
			out: nil,
			res: result.New(result.CodeNotFound, "アカウントが見つかりません"),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			u := user.NewUsecase(c.setup(ctrl))

			out, res := u.GetByID(context.Background(), c.in)

			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("Usecase.GetByID() result mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(c.out, out); diff != "" {
				t.Errorf("Usecase.GetByID() output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
