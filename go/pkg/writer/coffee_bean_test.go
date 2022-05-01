package writer_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeebean"
	"github.com/k3forx/coffee_memo/pkg/ent/userscoffeebean"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/writer"
	db_helper "github.com/k3forx/coffee_memo/test/db"
)

func TestCoffeeBean_Create(t *testing.T) {
	t.Parallel()

	impl := writer.NewCoffeeBeanWriter(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "create@example.com"
	})

	cases := map[string]struct {
		coffeeBean *model.CoffeeBean
		user       *model.User
	}{
		"success": {
			coffeeBean: &model.CoffeeBean{
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

			exists, err := testClient.CoffeeBean.Query().
				Where(coffeebean.IDEQ(int32(c.coffeeBean.ID))).Exist(ctx)
			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}
			if diff := cmp.Diff(true, exists); diff != "" {
				t.Errorf("coffee bean should be created")
			}

			exists, err = testClient.UsersCoffeeBean.Query().
				Where(userscoffeebean.CoffeeBeanIDEQ(int32(c.coffeeBean.ID))).Exist(ctx)
			if err != nil {
				t.Errorf("err should be nil, but got %v", err)
			}
			if diff := cmp.Diff(true, exists); diff != "" {
				t.Errorf("users coffee beans should be created")
			}
		})
	}
}

func TestCoffeeBean_DeleteByID(t *testing.T) {
	t.Parallel()

	impl := writer.NewCoffeeBeanWriter(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "delete-by-id"
	})
	coffeeBean := db_helper.InsertAndDeleteCoffeeBean(t, testClient)
	usersCoffeeBean := db_helper.InsertAndDeleteUsersCoffeeBeans(t, testClient, user, coffeeBean)

	cases := map[string]struct {
		coffeeBean model.CoffeeBean
	}{
		"success": {
			coffeeBean: model.NewCoffeeBean(coffeeBean),
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			if err := impl.DeleteByID(ctx, &c.coffeeBean); err != nil {
				t.Errorf("err should be nil but, got %q", err)
			}

			ucb, err := testClient.UsersCoffeeBean.Query().
				Where(userscoffeebean.IDEQ(int32(usersCoffeeBean.ID))).
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
