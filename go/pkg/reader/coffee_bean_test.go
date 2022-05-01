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

func TestCoffeeBean_GetByID(t *testing.T) {
	t.Parallel()

	coffeeBeanReader := reader.NewCoffeeBeanReader(testClient)
	coffeeBean := db_helper.InsertAndDeleteCoffeeBean(t, testClient)

	cases := map[string]struct {
		coffeeBeanID int
		expected     model.CoffeeBean
	}{
		"has_row": {
			coffeeBeanID: int(coffeeBean.ID),
			expected: model.CoffeeBean{
				ID:          int(coffeeBean.ID),
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
			coffeeBeanID: -111,
			expected:     model.CoffeeBean{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := coffeeBeanReader.GetByID(context.Background(), c.coffeeBeanID)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			if diff := cmp.Diff(c.expected, actual); diff != "" {
				t.Errorf("GetByID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCoffeeBean_GetByIDWithUser(t *testing.T) {
	t.Parallel()

	impl := reader.NewCoffeeBeanReader(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "get-by-id-with-user@example.com"
	})
	coffeeBean := db_helper.InsertAndDeleteCoffeeBean(t, testClient)
	_ = db_helper.InsertAndDeleteUsersCoffeeBeans(t, testClient, user, coffeeBean)

	cases := map[string]struct {
		coffeeBeanID int
		expected     model.CoffeeBean
	}{
		"has_row": {
			coffeeBeanID: int(coffeeBean.ID),
			expected: model.CoffeeBean{
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
			expected: model.CoffeeBean{
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
				cmpopts.IgnoreFields(model.CoffeeBean{}, "ID"),
				cmpopts.IgnoreFields(model.User{}, "ID"),
			}
			if diff := cmp.Diff(c.expected, actual, opts); diff != "" {
				t.Errorf("GetByIDWithUser mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCoffeeBean_GetAllByUserID(t *testing.T) {
	t.Parallel()

	coffeeBeanReader := reader.NewCoffeeBeanReader(testClient)
	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
		u.Email = "GetAllByUserID@example.com"
	})
	coffeeBean1 := db_helper.InsertAndDeleteCoffeeBean(t, testClient)
	coffeeBean2 := db_helper.InsertAndDeleteCoffeeBean(t, testClient, func(cb *ent.CoffeeBean) {
		cb.Name = "test"
		cb.FarmName = "test farm"
		cb.RoastDegree = model.RoastDegreeChinamon.String()
	})
	_ = db_helper.InsertAndDeleteUsersCoffeeBeans(t, testClient, user, coffeeBean1)
	_ = db_helper.InsertAndDeleteUsersCoffeeBeans(t, testClient, user, coffeeBean2)
	// t.Cleanup(func() {
	// 	db_helper.DeleteUsersCoffeeBeans(context.Background(), testClient, usersCoffeeBeans1)
	// 	db_helper.DeleteUsersCoffeeBeans(context.Background(), testClient, usersCoffeeBeans2)
	// 	db_helper.DeleteCoffeeBean(context.Background(), testClient, coffeeBean1)
	// 	db_helper.DeleteCoffeeBean(context.Background(), testClient, coffeeBean2)
	// 	db_helper.DeleteUser(context.Background(), testClient, user)
	// })

	cases := map[string]struct {
		userID   int
		expected []model.CoffeeBean
	}{
		"has_rows": {
			userID: int(user.ID),
			expected: []model.CoffeeBean{
				{
					Name:     "イルガチャフィ",
					FarmName: "",
					Country:  "エチオピア",
					User: model.User{
						Username:  "test-user",
						Email:     "GetAllByUserID@example.com",
						Flags:     model.UserFlags{0},
						CreatedAt: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC),
					},
					RoastDegree: model.RoastDegreeLight,
					RoastedAt:   time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
					CreatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
				},
				{
					Name:     "test",
					FarmName: "test farm",
					Country:  "エチオピア",
					User: model.User{
						Username:  "test-user",
						Email:     "GetAllByUserID@example.com",
						Flags:     model.UserFlags{0},
						CreatedAt: time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC),
					},
					RoastDegree: model.RoastDegreeChinamon,
					RoastedAt:   time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
					CreatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		"no_row": {
			userID:   -111,
			expected: []model.CoffeeBean{},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := coffeeBeanReader.GetAllByUserID(context.Background(), c.userID)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(model.CoffeeBean{}, "ID"),
				cmpopts.IgnoreFields(model.User{}, "ID"),
			}
			if diff := cmp.Diff(c.expected, actual, opts); diff != "" {
				t.Errorf("GetAllByUserID mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
