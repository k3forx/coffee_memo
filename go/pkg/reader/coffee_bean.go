package reader

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func NewCoffeeBeanReader(db *ent.Client) *CoffeeBeanReader {
	return &CoffeeBeanReader{
		db: db,
	}
}

//go:generate mockgen -source=./coffee_bean.go -destination=./mock/coffee_bean_mock.go -package=reader
type CoffeeBean interface {
	GetAllByUserID(ctx context.Context, userID int) ([]model.CoffeeBean, error)
}

type CoffeeBeanReader struct {
	db *ent.Client
}

var _ CoffeeBean = (*CoffeeBeanReader)(nil)

func (impl *CoffeeBeanReader) GetAllByUserID(ctx context.Context, userID int) ([]model.CoffeeBean, error) {
	return []model.CoffeeBean{}, nil
}
