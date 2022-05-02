package user_coffee_bean_test

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
	"github.com/k3forx/coffee_memo/pkg/usecase/user_coffee_bean"
	writerMock "github.com/k3forx/coffee_memo/pkg/writer/mock"
)

func TestUsecase_GetAllByUserID(t *testing.T) {
	t.Parallel()

	err := errors.New("server error")
	const (
		userId = 1
	)
	returnedUser := model.User{
		ID: userId,
	}
	returnedCoffeeBeans := []model.UserCoffeeBean{
		{
			ID: 11,
		},
		{
			ID: 22,
		},
	}

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    user_coffee_bean.GetAllByUserIDInput
		out   *user_coffee_bean.GetAllByUserIDOutput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userId).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetAllByUserID(gomock.Any(), userId).
					Return(returnedCoffeeBeans, nil)

				return injector
			},
			in: user_coffee_bean.GetAllByUserIDInput{
				UserID: userId,
			},
			out: &user_coffee_bean.GetAllByUserIDOutput{
				CoffeeBeans: returnedCoffeeBeans,
			},
			res: result.OK(),
		},
		"error_in_getting_user": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userId).
					Return(model.User{}, err)

				return injector
			},
			in: user_coffee_bean.GetAllByUserIDInput{
				UserID: userId,
			},
			out: nil,
			res: result.Error(),
		},
		"user_does_not_exist": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userId).
					Return(model.User{}, nil)

				return injector
			},
			in: user_coffee_bean.GetAllByUserIDInput{
				UserID: userId,
			},
			out: nil,
			res: result.New(result.CodeNotFound, "アカウントが存在しません。"),
		},
		"error_in_getting_coffee_beans": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userId).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetAllByUserID(gomock.Any(), userId).
					Return(nil, err)

				return injector
			},
			in: user_coffee_bean.GetAllByUserIDInput{
				UserID: userId,
			},
			out: nil,
			res: result.Error(),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := c.setup(ctrl)
			u := user_coffee_bean.NewUsecase(injector)

			out, res := u.GetAllByUserID(context.Background(), c.in)
			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("GetAllByUserID result mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(c.out, out); diff != "" {
				t.Errorf("GetAllByUserID output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUsecase_Create(t *testing.T) {
	t.Parallel()

	const userId = 1
	err := errors.New("server error")
	returnedUser := model.User{
		ID: userId,
	}

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    user_coffee_bean.CreateInput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userId).
					Return(returnedUser, nil)

				userCoffeeBeanWriter := injector.Writer.UserCoffeeBean.(*writerMock.MockUserCoffeeBean)
				userCoffeeBeanWriter.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

				return injector
			},
			in: user_coffee_bean.CreateInput{
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
			in: user_coffee_bean.CreateInput{
				UserId: userId,
				Name:   "イルガチャフィ",
			},
			res: result.New(result.CodeBadRequest, "無効なデータです。"),
		},
		"empty_name": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				return inject.NewMockInjector(ctrl)
			},
			in: user_coffee_bean.CreateInput{
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
			in: user_coffee_bean.CreateInput{
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
			in: user_coffee_bean.CreateInput{
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

				userCoffeeBeanWriter := injector.Writer.UserCoffeeBean.(*writerMock.MockUserCoffeeBean)
				userCoffeeBeanWriter.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(err)

				return injector
			},
			in: user_coffee_bean.CreateInput{
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
			u := user_coffee_bean.NewUsecase(injector)

			res := u.Create(context.Background(), c.in)
			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("Create result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUsecase_EditByID(t *testing.T) {
	t.Parallel()

	err := errors.New("server error")
	const (
		userID           = 1
		userCoffeeBeanID = 12
	)

	returnedUser := model.User{
		ID: userID,
	}
	returnedUserCoffeeBean := model.UserCoffeeBean{
		ID:   userCoffeeBeanID,
		User: returnedUser,
	}

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    user_coffee_bean.EditByIDInput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByID(gomock.Any(), userCoffeeBeanID).
					Return(returnedUserCoffeeBean, nil)

				userCoffeeBeanWriter := injector.Writer.UserCoffeeBean.(*writerMock.MockUserCoffeeBean)
				userCoffeeBeanWriter.EXPECT().UpdateByID(
					gomock.Any(), gomock.AssignableToTypeOf(&model.UserCoffeeBean{})).Return(nil)

				return injector
			},
			in: user_coffee_bean.EditByIDInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
				Name:             "updated name",
			},
			res: result.OK(),
		},
		"error_in_getting_user": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(model.User{}, err)

				return injector
			},
			in: user_coffee_bean.EditByIDInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.Error(),
		},
		"user_does_not_exist": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(model.User{}, nil)

				return injector
			},
			in: user_coffee_bean.EditByIDInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.New(result.CodeNotFound, "アカウントが見つかりません"),
		},
		"error_in_getting_user_coffee_bean": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByID(gomock.Any(), userCoffeeBeanID).
					Return(model.UserCoffeeBean{}, err)

				return injector
			},
			in: user_coffee_bean.EditByIDInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.Error(),
		},
		"user_coffee_bean_does_not_exist": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByID(gomock.Any(), userCoffeeBeanID).
					Return(model.UserCoffeeBean{}, nil)

				return injector
			},
			in: user_coffee_bean.EditByIDInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.New(result.CodeNotFound, "コーヒー豆が存在しません"),
		},
		"user_is_different": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByID(gomock.Any(), userCoffeeBeanID).
					Return(
						model.UserCoffeeBean{
							ID: userCoffeeBeanID,
						},
						nil,
					)

				return injector
			},
			in: user_coffee_bean.EditByIDInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.New(result.CodeForbidden, "編集できません"),
		},
		"empty_name": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByID(gomock.Any(), userCoffeeBeanID).
					Return(returnedUserCoffeeBean, nil)

				return injector
			},
			in: user_coffee_bean.EditByIDInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.New(result.CodeBadRequest, "名前を入力してください"),
		},
		"error_in_updating_user_coffee_bean": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByID(gomock.Any(), userCoffeeBeanID).
					Return(returnedUserCoffeeBean, nil)

				userCoffeeBeanWriter := injector.Writer.UserCoffeeBean.(*writerMock.MockUserCoffeeBean)
				userCoffeeBeanWriter.EXPECT().UpdateByID(
					gomock.Any(), gomock.AssignableToTypeOf(&model.UserCoffeeBean{})).Return(err)

				return injector
			},
			in: user_coffee_bean.EditByIDInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
				Name:             "updated name",
			},
			res: result.Error(),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			u := user_coffee_bean.NewUsecase(c.setup(ctrl))

			res := u.EditByID(context.Background(), c.in)
			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("EditByID result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUsecase_DeleteByID(t *testing.T) {
	t.Parallel()

	err := errors.New("server error")
	const (
		userID       = 1
		coffeeBeanID = 123
	)
	returnedUser := model.User{
		ID: userID,
	}
	returnedCoffeeBean := model.UserCoffeeBean{
		ID:   coffeeBeanID,
		User: returnedUser,
	}

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    user_coffee_bean.DeleteByIDInput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), coffeeBeanID).
					Return(returnedCoffeeBean, nil)

				userCoffeeBeanWriter := injector.Writer.UserCoffeeBean.(*writerMock.MockUserCoffeeBean)
				userCoffeeBeanWriter.EXPECT().DeleteByID(gomock.Any(), gomock.AssignableToTypeOf(&model.UserCoffeeBean{})).
					Return(nil)

				return injector
			},
			in: user_coffee_bean.DeleteByIDInput{
				UserID:       userID,
				CoffeeBeanID: coffeeBeanID,
			},
			res: result.OK(),
		},
		"error_in_getting_user": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(model.User{}, err)

				return injector
			},
			in: user_coffee_bean.DeleteByIDInput{
				UserID:       userID,
				CoffeeBeanID: coffeeBeanID,
			},
			res: result.Error(),
		},
		"user_does_not_exist": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(model.User{}, nil)

				return injector
			},
			in: user_coffee_bean.DeleteByIDInput{
				UserID:       userID,
				CoffeeBeanID: coffeeBeanID,
			},
			res: result.New(result.CodeNotFound, "アカウントが見つかりません"),
		},
		"error_in_getting_coffee_bean": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), coffeeBeanID).
					Return(model.UserCoffeeBean{}, err)

				return injector
			},
			in: user_coffee_bean.DeleteByIDInput{
				UserID:       userID,
				CoffeeBeanID: coffeeBeanID,
			},
			res: result.Error(),
		},
		"coffee_bean_does_not_exist": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), coffeeBeanID).
					Return(model.UserCoffeeBean{}, nil)

				return injector
			},
			in: user_coffee_bean.DeleteByIDInput{
				UserID:       userID,
				CoffeeBeanID: coffeeBeanID,
			},
			res: result.New(result.CodeNotFound, "コーヒー豆が存在しません"),
		},
		"invald_user": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), coffeeBeanID).
					Return(
						model.UserCoffeeBean{
							ID: coffeeBeanID,
							User: model.User{
								ID: 111111,
							},
						},
						nil,
					)

				return injector
			},
			in: user_coffee_bean.DeleteByIDInput{
				UserID:       userID,
				CoffeeBeanID: coffeeBeanID,
			},
			res: result.New(result.CodeForbidden, "削除できません"),
		},
		"error_in_deleting_coffee_bean": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), coffeeBeanID).
					Return(returnedCoffeeBean, nil)

				userCoffeeBeanWriter := injector.Writer.UserCoffeeBean.(*writerMock.MockUserCoffeeBean)
				userCoffeeBeanWriter.EXPECT().DeleteByID(gomock.Any(), gomock.AssignableToTypeOf(&model.UserCoffeeBean{})).
					Return(err)

				return injector
			},
			in: user_coffee_bean.DeleteByIDInput{
				UserID:       userID,
				CoffeeBeanID: coffeeBeanID,
			},
			res: result.New(result.CodeInternalError, "削除に失敗しました"),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			u := user_coffee_bean.NewUsecase(c.setup(ctrl))

			res := u.DeleteByID(context.Background(), c.in)
			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("DeleteByID result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
