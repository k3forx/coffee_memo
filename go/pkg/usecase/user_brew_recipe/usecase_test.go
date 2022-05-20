package user_brew_recipe_test

import (
	"context"
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
				userBrewRecipeWriter.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

				return injector
			},
			in: user_brew_recipe.CreateInput{
				UserID:           userID,
				UserCoffeeBeanID: userCoffeeBeanID,
			},
			res: result.OK(),
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
