package reader_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/reader"
	db_helper "github.com/k3forx/coffee_memo/test/db"
)

func TestUserCoffeeBean_GetByID(t *testing.T) {
	t.Parallel()

	impl := reader.NewUserCoffeeBeanReader(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "get-by-id@example.com"
	})
	userCoffeeBean := db_helper.InsertAndDeleteUserCoffeeBean(t, testClient, user)

	cases := map[string]struct {
		id       int
		expected model.UserCoffeeBean
	}{
		"has_row": {
			id: int(userCoffeeBean.ID),
			expected: model.UserCoffeeBean{
				ID: int(userCoffeeBean.ID),
				User: model.User{
					ID: int(user.ID),
				},
				Name:        "イルガチャフィ",
				FarmName:    "",
				Country:     "エチオピア",
				RoastDegree: model.RoastDegreeLight,
				RoastedAt:   time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
				CreatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
			},
		},
		"no_row": {
			id:       -111,
			expected: model.UserCoffeeBean{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := impl.GetByID(context.Background(), c.id)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("GetByID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUserCoffeeBean_GetByIDWithUser(t *testing.T) {
	t.Parallel()

	impl := reader.NewUserCoffeeBeanReader(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "get-by-id-with-user@example.com"
	})
	coffeeBean := db_helper.InsertAndDeleteUserCoffeeBean(t, testClient, user)

	cases := map[string]struct {
		coffeeBeanID int
		expected     model.UserCoffeeBean
	}{
		"has_row": {
			coffeeBeanID: int(coffeeBean.ID),
			expected: model.UserCoffeeBean{
				Status:   model.CoffeeBeanStatusActive,
				Name:     "イルガチャフィ",
				FarmName: "",
				Country:  "エチオピア",
				User: model.User{
					Username:  "test-user",
					Email:     "get-by-id-with-user@example.com",
					Flags:     model.UserFlags{0},
					CreatedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC),
				},
				RoastDegree: model.RoastDegreeLight,
				RoastedAt:   time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
				CreatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
			},
		},
		"no_row": {
			coffeeBeanID: -111,
			expected: model.UserCoffeeBean{
				User: model.User{
					Flags: nil,
				},
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := impl.GetByIDWithUser(context.Background(), c.coffeeBeanID)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(model.UserCoffeeBean{}, "ID"),
				cmpopts.IgnoreFields(model.User{}, "ID"),
			}
			if diff := cmp.Diff(c.expected, actual, opts); diff != "" {
				t.Errorf("GetByIDWithUser mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUserCoffeeBean_GetAllByUserID(t *testing.T) {
	t.Parallel()

	impl := reader.NewUserCoffeeBeanReader(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "GetAllByUserID@example.com"
	})
	_ = db_helper.InsertAndDeleteUserCoffeeBean(t, testClient, user)
	_ = db_helper.InsertAndDeleteUserCoffeeBean(t, testClient, user, func(cb *ent.UserCoffeeBean) {
		cb.Name = "test"
		cb.FarmName = "test farm"
		cb.RoastDegree = model.RoastDegreeChinamon.String()
	})
	// t.Cleanup(func() {
	// 	db_helper.DeleteUsersCoffeeBeans(context.Background(), testClient, usersCoffeeBeans1)
	// 	db_helper.DeleteUsersCoffeeBeans(context.Background(), testClient, usersCoffeeBeans2)
	// 	db_helper.DeleteCoffeeBean(context.Background(), testClient, coffeeBean1)
	// 	db_helper.DeleteCoffeeBean(context.Background(), testClient, coffeeBean2)
	// 	db_helper.DeleteUser(context.Background(), testClient, user)
	// })

	cases := map[string]struct {
		userID   int
		expected []model.UserCoffeeBean
	}{
		"has_rows": {
			userID: int(user.ID),
			expected: []model.UserCoffeeBean{
				{
					Name:        "イルガチャフィ",
					FarmName:    "",
					Country:     "エチオピア",
					User:        model.User{},
					RoastDegree: model.RoastDegreeLight,
					RoastedAt:   time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
					CreatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
				},
				{
					Name:        "test",
					FarmName:    "test farm",
					Country:     "エチオピア",
					User:        model.User{},
					RoastDegree: model.RoastDegreeChinamon,
					RoastedAt:   time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
					CreatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		"no_row": {
			userID:   -111,
			expected: []model.UserCoffeeBean{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := impl.GetAllByUserID(context.Background(), c.userID)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(model.UserCoffeeBean{}, "ID"),
				cmpopts.IgnoreFields(model.User{}, "ID"),
			}
			if diff := cmp.Diff(c.expected, actual, opts); diff != "" {
				t.Errorf("GetAllByUserID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
