package coffee_bean_test

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
	"github.com/k3forx/coffee_memo/pkg/usecase/coffee_bean"
	writerMock "github.com/k3forx/coffee_memo/pkg/writer/mock"
)

func TestUsecase_Create(t *testing.T) {
	t.Parallel()

	const userId = 1
	err := errors.New("server error")
	returnedUser := model.User{
		ID: userId,
	}

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    coffee_bean.CreateInput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userId).
					Return(returnedUser, nil)

				coffeeBeanWriter := injector.Writer.CoffeeBean.(*writerMock.MockCoffeeBean)
				coffeeBeanWriter.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

				return injector
			},
			in: coffee_bean.CreateInput{
				UserId:      userId,
				Name:        "イルガチャフィ",
				RoastDegree: model.RoastDegreeChinamon,
			},
			res: result.OK(),
		},
		"invalid_roast_degree": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				return inject.NewMockInjector(ctrl)
			},
			in: coffee_bean.CreateInput{
				UserId: userId,
				Name:   "イルガチャフィ",
			},
			res: result.New(result.CodeBadRequest, "無効なデータです。"),
		},
		"empty_name": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				return inject.NewMockInjector(ctrl)
			},
			in: coffee_bean.CreateInput{
				UserId:      userId,
				RoastDegree: model.RoastDegreeChinamon,
			},
			res: result.New(result.CodeBadRequest, "無効なデータです。"),
		},
		"error_in_getting_user": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userId).
					Return(model.User{}, err)

				return injector
			},
			in: coffee_bean.CreateInput{
				UserId:      userId,
				Name:        "イルガチャフィ",
				RoastDegree: model.RoastDegreeChinamon,
			},
			res: result.Error(),
		},
		"user_does_not_found": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userId).
					Return(model.User{}, nil)

				return injector
			},
			in: coffee_bean.CreateInput{
				UserId:      userId,
				Name:        "イルガチャフィ",
				RoastDegree: model.RoastDegreeChinamon,
			},
			res: result.New(result.CodeNotFound, "アカウントが見つかりません。"),
		},
		"error_in_creating_coffee_bean": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userId).
					Return(returnedUser, nil)

				coffeeBeanWriter := injector.Writer.CoffeeBean.(*writerMock.MockCoffeeBean)
				coffeeBeanWriter.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(err)

				return injector
			},
			in: coffee_bean.CreateInput{
				UserId:      userId,
				Name:        "イルガチャフィ",
				RoastDegree: model.RoastDegreeChinamon,
			},
			res: result.New(result.CodeInternalError, "コービー豆の登録に失敗しました。"),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := c.setup(ctrl)
			u := coffee_bean.NewUsecase(injector)

			res := u.Create(context.Background(), c.in)
			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("Create result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
