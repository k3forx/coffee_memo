package reader

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/usercoffeebean"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func NewUserCoffeeBeanReader(db *ent.Client) *UserCoffeeBeanReader {
	return &UserCoffeeBeanReader{
		db: db,
	}
}

//go:generate mockgen -source=./user_coffee_bean.go -destination=./mock/user_coffee_bean_mock.go -package=reader
type UserCoffeeBean interface {
	GetByID(ctx context.Context, id int) (model.UserCoffeeBean, error)
	GetByIDWithUser(ctx context.Context, id int) (model.UserCoffeeBean, error)
	GetAllByUserID(ctx context.Context, userID int) ([]model.UserCoffeeBean, error)
}

type UserCoffeeBeanReader struct {
	db *ent.Client
}

var _ UserCoffeeBean = (*UserCoffeeBeanReader)(nil)

func (impl *UserCoffeeBeanReader) GetByID(ctx context.Context, id int) (model.UserCoffeeBean, error) {
	ucb, err := impl.db.UserCoffeeBean.Query().
		Where(usercoffeebean.IDEQ(int32(id))).Only(ctx)
	if err != nil {
		return model.UserCoffeeBean{}, ent.MaskNotFound(err)
	}
	return model.NewUserCoffeeBean(ucb), nil
}

func (impl *UserCoffeeBeanReader) GetByIDWithUser(ctx context.Context, id int) (model.UserCoffeeBean, error) {
	usersCoffeeBean, err := impl.db.UserCoffeeBean.Query().
		Where(usercoffeebean.IDEQ(int32(id))).
		WithUser().
		Only(ctx)
	if err != nil {
		return model.UserCoffeeBean{}, ent.MaskNotFound(err)
	}
	return model.NewUserCoffeeBean(usersCoffeeBean), nil
}

func (impl *UserCoffeeBeanReader) GetAllByUserID(ctx context.Context, userID int) ([]model.UserCoffeeBean, error) {
	es, err := impl.db.UserCoffeeBean.
		Query().
		Where(usercoffeebean.UserIDEQ(int32(userID))).
		All(ctx)
	if err != nil {
		return nil, ent.MaskNotFound(err)
	}

	ms := make([]model.UserCoffeeBean, len(es))
	for i, e := range es {
		ms[i] = model.NewUserCoffeeBean(e)
	}
	return ms, nil
}
