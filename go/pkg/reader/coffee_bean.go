package reader

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeebean"
	"github.com/k3forx/coffee_memo/pkg/ent/userscoffeebean"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func NewCoffeeBeanReader(db *ent.Client) *CoffeeBeanReader {
	return &CoffeeBeanReader{
		db: db,
	}
}

//go:generate mockgen -source=./coffee_bean.go -destination=./mock/coffee_bean_mock.go -package=reader
type CoffeeBean interface {
	GetByID(ctx context.Context, coffeeBeanID int) (model.CoffeeBean, error)
	GetByIDWithUser(ctx context.Context, coffeeBeanID int) (model.CoffeeBean, error)
	GetAllByUserID(ctx context.Context, userID int) ([]model.CoffeeBean, error)
}

type CoffeeBeanReader struct {
	db *ent.Client
}

var _ CoffeeBean = (*CoffeeBeanReader)(nil)

func (impl *CoffeeBeanReader) GetByID(ctx context.Context, coffeeBeanID int) (model.CoffeeBean, error) {
	cb, err := impl.db.CoffeeBean.Query().Where(coffeebean.IDEQ(int32(coffeeBeanID))).Only(ctx)
	if err != nil {
		return model.CoffeeBean{}, ent.MaskNotFound(err)
	}
	return model.NewCoffeeBean(cb), nil
}

func (impl *CoffeeBeanReader) GetByIDWithUser(ctx context.Context, coffeeBeanID int) (model.CoffeeBean, error) {
	usersCoffeeBean, err := impl.db.UsersCoffeeBean.
		Query().
		Where(userscoffeebean.CoffeeBeanIDEQ(int32(coffeeBeanID))).
		WithCoffeeBean().
		Only(ctx)
	if err != nil {
		return model.CoffeeBean{}, ent.MaskNotFound(err)
	}
	return model.NewCoffeeBeanWithUser(usersCoffeeBean), nil
}

func (impl *CoffeeBeanReader) GetAllByUserID(ctx context.Context, userID int) ([]model.CoffeeBean, error) {
	es, err := impl.db.UsersCoffeeBean.
		Query().
		Where(userscoffeebean.UserIDEQ(int32(userID))).
		WithCoffeeBean().
		WithUser().
		All(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}

	ms := make([]model.CoffeeBean, len(es))
	for i, e := range es {
		ms[i] = model.NewCoffeeBeanWithUser(e)
	}
	return ms, nil
}
