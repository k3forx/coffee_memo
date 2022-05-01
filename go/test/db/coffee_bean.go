package db_helper

import (
	"context"
	"testing"
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeebean"
)

func newCoffeeBean() *ent.CoffeeBean {
	return &ent.CoffeeBean{
		Name:        "イルガチャフィ",
		FarmName:    "",
		Country:     "エチオピア",
		RoastDegree: "light",
		RoastedAt:   time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2022, time.April, 29, 0, 0, 0, 0, time.UTC),
	}
}

func InsertCoffeeBean(client *ent.Client, coffeeBean *ent.CoffeeBean) (*ent.CoffeeBean, error) {
	cb, err := client.CoffeeBean.Create().
		SetName(coffeeBean.Name).
		SetFarmName(coffeeBean.FarmName).
		SetCountry(coffeeBean.Country).
		SetRoastDegree(coffeeBean.RoastDegree).
		SetRoastedAt(coffeeBean.RoastedAt).
		SetCreatedAt(coffeeBean.CreatedAt).
		SetUpdatedAt(coffeeBean.UpdatedAt).
		Save(context.Background())
	return cb, err
}

func DeleteCoffeeBean(ctx context.Context, client *ent.Client, coffeeBean *ent.CoffeeBean) {
	client.CoffeeBean.Delete().Where(coffeebean.IDEQ(coffeeBean.ID))
}

func InsertAndDeleteCoffeeBean(tb testing.TB, client *ent.Client, setters ...func(cb *ent.CoffeeBean)) *ent.CoffeeBean {
	tb.Helper()

	coffeeBean := newCoffeeBean()
	for _, setter := range setters {
		setter(coffeeBean)
	}

	cb, err := InsertCoffeeBean(client, coffeeBean)
	if err != nil {
		tb.Fatalf("InsertAndDeleteCoffeeBeans failed: %+v\n", err)
	}
	tb.Cleanup(func() {
		DeleteCoffeeBean(context.Background(), client, cb)
	})

	return cb
}
