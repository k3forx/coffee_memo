package writer

import (
	"context"
	"fmt"
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/userscoffeebean"
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
	Create(ctx context.Context, coffeeBean *model.CoffeeBean, user *model.User) error
}

type CoffeeBeanWriter struct {
	db *ent.Client
}

var _ CoffeeBean = (*CoffeeBeanWriter)(nil)

func (impl *CoffeeBeanWriter) Create(ctx context.Context, coffeeBean *model.CoffeeBean, user *model.User) error {
	now := time.Now().UTC()

	tx, err := impl.db.Tx(ctx)
	if err != nil {
		return err
	}

	b, err := tx.CoffeeBean.Create().
		SetName(coffeeBean.Name).
		SetFarmName(coffeeBean.FarmName).
		SetCountry(coffeeBean.Country).
		SetRoastDegree(coffeeBean.RoastDegree.String()).
		SetRoastedAt(coffeeBean.RoastedAt).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return transaction.Rollback(tx, fmt.Errorf("failed to create coffee beans: %w", err))
	}
	*coffeeBean = model.NewCoffeeBean(b)

	_, err = tx.UsersCoffeeBean.Create().
		SetStatus(int32(model.CoffeeBeanStatusActive.Num())).
		SetUserID(int32(user.ID)).
		SetCoffeeBeanID(int32(coffeeBean.ID)).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return transaction.Rollback(tx, fmt.Errorf("failed to create users coffee beans: %w", err))
	}

	return tx.Commit()
}

func (impl *CoffeeBeanWriter) DeleteByID(ctx context.Context, coffeeBean *model.CoffeeBean) error {
	if _, err := impl.db.UsersCoffeeBean.
		Update().
		SetStatus(int32(model.CoffeeBeanStatusDeleted.Num())).
		Where(userscoffeebean.CoffeeBeanIDEQ(int32(coffeeBean.ID))).
		Save(ctx); err != nil {
		return err
	}
	return nil
}
