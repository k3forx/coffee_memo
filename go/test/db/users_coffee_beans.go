package db_helper

import (
	"context"
	"testing"
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/userscoffeebean"
)

func InsertUsersCoffeeBeans(tb testing.TB, client *ent.Client, user *ent.User, coffeeBean *ent.CoffeeBean) (*ent.UsersCoffeeBean, error) {
	ucb, err := client.UsersCoffeeBean.
		Create().
		SetCoffeeBean(coffeeBean).
		SetUser(user).
		SetCreatedAt(time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC)).
		SetUpdatedAt(time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC)).
		Save(context.Background())
	return ucb, err
}

func DeleteUsersCoffeeBeansByUser(ctx context.Context, client *ent.Client, user *ent.User) {
	client.UsersCoffeeBean.Delete().Where(userscoffeebean.UserIDEQ(user.ID))
}

func DeleteUsersCoffeeBeans(ctx context.Context, client *ent.Client, usersCoffeeBeans *ent.UsersCoffeeBean) {
	client.UsersCoffeeBean.Delete().Where(userscoffeebean.IDEQ(usersCoffeeBeans.ID))
}

func InsertAndDeleteUsersCoffeeBeans(tb testing.TB, client *ent.Client, user *ent.User, coffeeBean *ent.CoffeeBean) *ent.UsersCoffeeBean {
	tb.Helper()

	ucb, err := InsertUsersCoffeeBeans(tb, client, user, coffeeBean)
	if err != nil {
		tb.Fatalf("InsertUsersCoffeeBeans failed: %+v\n", err)
	}
	tb.Cleanup(func() {
		DeleteUsersCoffeeBeans(context.Background(), client, ucb)
	})

	return ucb
}
