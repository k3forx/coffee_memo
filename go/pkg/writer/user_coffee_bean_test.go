package writer_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/usercoffeebean"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/writer"
	db_helper "github.com/k3forx/coffee_memo/test/db"
)

func TestUserCoffeeBean_Create(t *testing.T) {
	t.Parallel()

	impl := writer.NewUserCoffeeBeanWriter(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "create@example.com"
	})

	cases := map[string]struct {
		coffeeBean *model.UserCoffeeBean
		user       *model.User
	}{
		"success": {
			coffeeBean: &model.UserCoffeeBean{
				User: model.User{
					ID: int(user.ID),
				},
				Name:        "test name",
				FarmName:    "test farm name",
				Country:     "test country",
				RoastDegree: model.RoastDegreeFrench,
				RoastedAt:   time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC),
				CreatedAt:   time.Date(2022, time.February, 20, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, time.February, 20, 0, 0, 0, 0, time.UTC),
			},
			user: &model.User{
				ID: int(user.ID),
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			if err := impl.Create(ctx, c.coffeeBean, c.user); err != nil {
				t.Errorf("err should be nil, but got %q\n", err)
			}

			exists, err := testClient.UserCoffeeBean.Query().
				Where(usercoffeebean.IDEQ(int32(c.coffeeBean.ID))).Exist(ctx)
			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}
			if diff := cmp.Diff(true, exists); diff != "" {
				t.Errorf("coffee bean should be created")
			}
		})
	}
}

func TestUserCoffeeBean_UpdateByID(t *testing.T) {
	t.Parallel()

	impl := writer.NewUserCoffeeBeanWriter(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "update-by-id@example.com"
	})
	ucb := db_helper.InsertAndDeleteUserCoffeeBean(t, testClient, user)
	userCoffeeBean := model.NewUserCoffeeBean(ucb)

	cases := map[string]struct {
		userCoffeeBean model.UserCoffeeBean
	}{
		"success": {
			userCoffeeBean: model.UserCoffeeBean{
				ID: userCoffeeBean.ID,
				User: model.User{
					ID: int(user.ID),
				},
				Name:      "updated name",
				FarmName:  "updated farm name",
				Country:   "updated country",
				RoastedAt: time.Now(),
				CreatedAt: time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Now(),
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			userCoffeeBean = c.userCoffeeBean
			if err := impl.UpdateByID(ctx, &userCoffeeBean); err != nil {
				t.Errorf("err should be nil but, got %q", err)
			}

			ucb, err := testClient.UserCoffeeBean.Get(ctx, int32(userCoffeeBean.ID))
			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}
			opts := cmp.Options{
				cmpopts.EquateApproxTime(time.Minute),
			}
			if diff := cmp.Diff(c.userCoffeeBean, model.NewUserCoffeeBean(ucb), opts); diff != "" {
				t.Errorf("Update mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUserCoffeeBean_DeleteByID(t *testing.T) {
	t.Parallel()

	impl := writer.NewUserCoffeeBeanWriter(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "delete-by-id"
	})
	userCoffeeBean := db_helper.InsertAndDeleteUserCoffeeBean(t, testClient, user)

	cases := map[string]struct {
		userCoffeeBean model.UserCoffeeBean
	}{
		"success": {
			userCoffeeBean: model.NewUserCoffeeBean(userCoffeeBean),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			if err := impl.DeleteByID(ctx, &c.userCoffeeBean); err != nil {
				t.Errorf("err should be nil but, got %q", err)
			}

			ucb, err := testClient.UserCoffeeBean.Query().
				Where(usercoffeebean.IDEQ(int32(userCoffeeBean.ID))).
				Only(ctx)
			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}
			if diff := cmp.Diff(model.CoffeeBeanStatusDeleted.Num(), int(ucb.Status)); diff != "" {
				t.Errorf("status should be deleted, but got %v", ucb.Status)
			}
		})
	}
}
