package db_helper

import (
	"context"
	"testing"
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/usercoffeebean"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func newUserCoffeeBean(user *ent.User) *ent.UserCoffeeBean {
	return &ent.UserCoffeeBean{
		UserID:      user.ID,
		Status:      int32(model.CoffeeBeanStatusActive.Num()),
		Name:        "イルガチャフィ",
		FarmName:    "",
		Country:     "エチオピア",
		RoastDegree: "light",
		RoastedAt:   time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
	}
}

func InsertUserCoffeeBean(client *ent.Client, userCoffeeBean *ent.UserCoffeeBean) (*ent.UserCoffeeBean, error) {
	cb, err := client.UserCoffeeBean.Create().
		SetUserID(userCoffeeBean.UserID).
		SetStatus(userCoffeeBean.Status).
		SetName(userCoffeeBean.Name).
		SetFarmName(userCoffeeBean.FarmName).
		SetCountry(userCoffeeBean.Country).
		SetRoastDegree(userCoffeeBean.RoastDegree).
		SetRoastedAt(userCoffeeBean.RoastedAt).
		SetCreatedAt(userCoffeeBean.CreatedAt).
		SetUpdatedAt(userCoffeeBean.UpdatedAt).
		Save(context.Background())
	return cb, err
}

func DeleteUserCoffeeBean(ctx context.Context, client *ent.Client, userCoffeeBean *ent.UserCoffeeBean) {
	client.UserCoffeeBean.Delete().Where(usercoffeebean.IDEQ(userCoffeeBean.ID))
}

func InsertAndDeleteUserCoffeeBean(tb testing.TB, client *ent.Client, user *ent.User, setters ...func(cb *ent.UserCoffeeBean)) *ent.UserCoffeeBean {
	tb.Helper()

	coffeeBean := newUserCoffeeBean(user)
	for _, setter := range setters {
		setter(coffeeBean)
	}

	cb, err := InsertUserCoffeeBean(client, coffeeBean)
	if err != nil {
		tb.Fatalf("InsertAndDeleteCoffeeBeans failed: %+v\n", err)
	}
	tb.Cleanup(func() {
		DeleteUserCoffeeBean(context.Background(), client, cb)
	})

	return cb
}
