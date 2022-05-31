package user_brew_recipe_test

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
	"github.com/k3forx/coffee_memo/pkg/usecase/user_brew_recipe"
	writerMock "github.com/k3forx/coffee_memo/pkg/writer/mock"
)

func TestUserBrewRecipeUsecase_Create(t *testing.T) {
	t.Parallel()

	err := errors.New("server error")
	const (
		userID           = 1
		userCoffeeBeanID = 3
	)
	returnedUser := model.User{
		ID: userID,
	}
	returnedUserCoffeeBean := model.UserCoffeeBean{
		ID: userCoffeeBeanID,
		User: model.User{
			ID: userID,
		},
	}

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    user_brew_recipe.CreateInput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), userCoffeeBeanID).
					Return(returnedUserCoffeeBean, nil)

				userBrewRecipeWriter := injector.Writer.UserBrewRecipe.(*writerMock.MockUserBrewRecipe)
				userBrewRecipeWriter.EXPECT().Create(
					gomock.Any(),
					&model.UserBrewRecipe{
						Status:         model.BrewRecipeStatusCreated,
						User:           returnedUser,
						UserCoffeeBean: returnedUserCoffeeBean,
					},
				).Return(nil)

				return injector
			},
			in: user_brew_recipe.CreateInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
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
			in: user_brew_recipe.CreateInput{
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
			in: user_brew_recipe.CreateInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.New(result.CodeNotFound, "アカウントが存在しません"),
		},
		"error_in_getting_user_coffee_bean": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), userCoffeeBeanID).
					Return(model.UserCoffeeBean{}, err)

				return injector
			},
			in: user_brew_recipe.CreateInput{
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
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), userCoffeeBeanID).
					Return(model.UserCoffeeBean{}, nil)

				return injector
			},
			in: user_brew_recipe.CreateInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.New(result.CodeNotFound, "コーヒー豆が存在しません"),
		},
		"user_does_not_have_user_coffee_bean": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), userCoffeeBeanID).
					Return(model.UserCoffeeBean{ID: userCoffeeBeanID}, nil)

				return injector
			},
			in: user_brew_recipe.CreateInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.New(result.CodeForbidden, result.CodeForbidden.String()),
		},
		"error_in_creating_brew_recipe": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userCoffeeBeanReader := injector.Reader.UserCoffeeBean.(*readerMock.MockUserCoffeeBean)
				userCoffeeBeanReader.EXPECT().GetByIDWithUser(gomock.Any(), userCoffeeBeanID).
					Return(returnedUserCoffeeBean, nil)

				userBrewRecipeWriter := injector.Writer.UserBrewRecipe.(*writerMock.MockUserBrewRecipe)
				userBrewRecipeWriter.EXPECT().Create(
					gomock.Any(),
					&model.UserBrewRecipe{
						Status:         model.BrewRecipeStatusCreated,
						User:           returnedUser,
						UserCoffeeBean: returnedUserCoffeeBean,
					},
				).Return(err)

				return injector
			},
			in: user_brew_recipe.CreateInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.Error(),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := c.setup(ctrl)
			u := user_brew_recipe.NewUsecase(injector)

			res := u.Create(context.Background(), c.in)
			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("Create result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUserBrewRecipeUsecase_GetByID(t *testing.T) {
	t.Parallel()

	const (
		userID           = 123
		userBrewRecipeID = 345
	)
	returnedUser := model.User{
		ID: userID,
	}
	returnedUserBrewRecipe := model.UserBrewRecipe{
		ID: userBrewRecipeID,
		User: model.User{
			ID: userID,
		},
	}
	err := errors.New("server error")

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    user_brew_recipe.GetByIDInput
		out   *user_brew_recipe.GetByIDOutput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userBrewRecipeReader := injector.Reader.UserBrewRecipe.(*readerMock.MockUserBrewRecipe)
				userBrewRecipeReader.EXPECT().GetByID(gomock.Any(), userBrewRecipeID).
					Return(returnedUserBrewRecipe, nil)

				return injector
			},
			in: user_brew_recipe.GetByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			out: &user_brew_recipe.GetByIDOutput{
				UserBrewRecipe: returnedUserBrewRecipe,
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
			in: user_brew_recipe.GetByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			out: nil,
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
			in: user_brew_recipe.GetByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			out: nil,
			res: result.New(result.CodeNotFound, "アカウントが存在しません"),
		},
		"error_in_getting_user_brew_recipe": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userBrewRecipeReader := injector.Reader.UserBrewRecipe.(*readerMock.MockUserBrewRecipe)
				userBrewRecipeReader.EXPECT().GetByID(gomock.Any(), userBrewRecipeID).
					Return(model.UserBrewRecipe{}, err)

				return injector
			},
			in: user_brew_recipe.GetByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			out: nil,
			res: result.Error(),
		},
		"user_brew_recipe_does_not_exist": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userBrewRecipeReader := injector.Reader.UserBrewRecipe.(*readerMock.MockUserBrewRecipe)
				userBrewRecipeReader.EXPECT().GetByID(gomock.Any(), userBrewRecipeID).
					Return(model.UserBrewRecipe{}, nil)

				return injector
			},
			in: user_brew_recipe.GetByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			out: nil,
			res: result.New(result.CodeNotFound, "ドリップレシピが見つかりません"),
		},
		"user_brew_recipe_user_id_is_different": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userBrewRecipeReader := injector.Reader.UserBrewRecipe.(*readerMock.MockUserBrewRecipe)
				userBrewRecipeReader.EXPECT().GetByID(gomock.Any(), userBrewRecipeID).
					Return(
						model.UserBrewRecipe{
							ID: userBrewRecipeID,
							User: model.User{
								ID: 9,
							},
						},
						nil,
					)

				return injector
			},
			in: user_brew_recipe.GetByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			out: nil,
			res: result.New(result.CodeForbidden, result.CodeForbidden.String()),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := c.setup(ctrl)
			u := user_brew_recipe.NewUsecase(injector)

			out, res := u.GetByID(context.Background(), c.in)
			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("GetByID result mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(c.out, out); diff != "" {
				t.Errorf("GetByID output mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUserBrewRecipeUsecase_DeleteByID(t *testing.T) {
	t.Parallel()

	const (
		userID           = 123
		userBrewRecipeID = 345
	)
	returnedUser := model.User{
		ID: userID,
	}
	returnedUserBrewRecipe := model.UserBrewRecipe{
		ID: userBrewRecipeID,
		User: model.User{
			ID: userID,
		},
	}
	err := errors.New("server error")

	cases := map[string]struct {
		setup func(ctrl *gomock.Controller) inject.Injector
		in    user_brew_recipe.DeleteByIDInput
		res   *result.Result
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userBrewRecipeReader := injector.Reader.UserBrewRecipe.(*readerMock.MockUserBrewRecipe)
				userBrewRecipeReader.EXPECT().GetByID(gomock.Any(), userBrewRecipeID).
					Return(returnedUserBrewRecipe, nil)

				userBrewRecipeWriter := injector.Writer.UserBrewRecipe.(*writerMock.MockUserBrewRecipe)
				userBrewRecipeWriter.EXPECT().Delete(gomock.Any(), &returnedUserBrewRecipe).Return(nil)

				return injector
			},
			in: user_brew_recipe.DeleteByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
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
			in: user_brew_recipe.DeleteByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
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
			in: user_brew_recipe.DeleteByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			res: result.New(result.CodeNotFound, "アカウントが存在しません"),
		},
		"error_in_getting_user_brew_recipe": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userBrewRecipeReader := injector.Reader.UserBrewRecipe.(*readerMock.MockUserBrewRecipe)
				userBrewRecipeReader.EXPECT().GetByID(gomock.Any(), userBrewRecipeID).
					Return(model.UserBrewRecipe{}, err)

				return injector
			},
			in: user_brew_recipe.DeleteByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			res: result.Error(),
		},
		"user_brew_recipe_does_not_exist": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userBrewRecipeReader := injector.Reader.UserBrewRecipe.(*readerMock.MockUserBrewRecipe)
				userBrewRecipeReader.EXPECT().GetByID(gomock.Any(), userBrewRecipeID).
					Return(model.UserBrewRecipe{}, nil)

				return injector
			},
			in: user_brew_recipe.DeleteByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			res: result.New(result.CodeNotFound, "ドリップレシピが見つかりません"),
		},
		"user_brew_recipe_user_id_is_different": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userBrewRecipeReader := injector.Reader.UserBrewRecipe.(*readerMock.MockUserBrewRecipe)
				userBrewRecipeReader.EXPECT().GetByID(gomock.Any(), userBrewRecipeID).
					Return(
						model.UserBrewRecipe{
							ID: userBrewRecipeID,
							User: model.User{
								ID: 9,
							},
						},
						nil,
					)

				return injector
			},
			in: user_brew_recipe.DeleteByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			res: result.New(result.CodeForbidden, result.CodeForbidden.String()),
		},
		"error_in_deleting_user_brew_recipe": {
			setup: func(ctrl *gomock.Controller) inject.Injector {
				injector := inject.NewMockInjector(ctrl)

				userReader := injector.Reader.User.(*readerMock.MockUser)
				userReader.EXPECT().GetByID(gomock.Any(), userID).
					Return(returnedUser, nil)

				userBrewRecipeReader := injector.Reader.UserBrewRecipe.(*readerMock.MockUserBrewRecipe)
				userBrewRecipeReader.EXPECT().GetByID(gomock.Any(), userBrewRecipeID).
					Return(returnedUserBrewRecipe, nil)

				userBrewRecipeWriter := injector.Writer.UserBrewRecipe.(*writerMock.MockUserBrewRecipe)
				userBrewRecipeWriter.EXPECT().Delete(gomock.Any(), &returnedUserBrewRecipe).Return(err)

				return injector
			},
			in: user_brew_recipe.DeleteByIDInput{
				UserID:           userID,
				UserBrewRecipeID: userBrewRecipeID,
			},
			res: result.Error(),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := c.setup(ctrl)
			u := user_brew_recipe.NewUsecase(injector)

			res := u.DeleteByID(context.Background(), c.in)
			if diff := cmp.Diff(c.res, res); diff != "" {
				t.Errorf("GetByID result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
