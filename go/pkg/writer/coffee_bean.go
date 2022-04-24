package writer

import (
	"context"
	"fmt"
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/transaction"
)

func NewCoffeeBeanWriter(db *ent.Client) *CoffeeBeanWriter {
	return &CoffeeBeanWriter{
		db: db,
	}
}

//go:generate mockgen -source=./coffee_bean.go -destination=./mock/coffee_bean_mock.go -package=writer
type CoffeeBean interface {
	Create(ctx context.Context, coffeeBean *model.CoffeeBean) error
}

type CoffeeBeanWriter struct {
	db *ent.Client
}

var _ CoffeeBean = (*CoffeeBeanWriter)(nil)

func (impl *CoffeeBeanWriter) Create(ctx context.Context, coffeeBean *model.CoffeeBean) error {
	now := time.Now().UTC()

	tx, err := impl.db.Tx(ctx)
	if err != nil {
		return err
	}

	b, err := tx.CoffeeBean.Create().
		SetName(coffeeBean.Name).
		SetFarmName(coffeeBean.FarmName).
		SetCountry(coffeeBean.Country).
		SetRoastedDegree(coffeeBean.RoastDegree.String()).
		SetRoastedAt(coffeeBean.RoastedAt).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return transaction.Rollback(tx, fmt.Errorf("failed to create coffee beans: %w", err))
	}
	*coffeeBean = model.NewCoffeeBean(b)

	return tx.Commit()
}
