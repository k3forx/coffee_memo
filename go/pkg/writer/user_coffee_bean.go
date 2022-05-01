package writer

import (
	"context"
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
	"github.com/k3forx/coffee_memo/pkg/ent/usercoffeebean"
	"github.com/k3forx/coffee_memo/pkg/model"
)

func NewUserCoffeeBeanWriter(db *ent.Client) *UserCoffeeBeanWriter {
	return &UserCoffeeBeanWriter{
		db: db,
	}
}

//go:generate mockgen -source=./user_coffee_bean.go -destination=./mock/user_coffee_bean_mock.go -package=writer
type UserCoffeeBean interface {
	Create(ctx context.Context, coffeeBean *model.UserCoffeeBean, user *model.User) error
	DeleteByID(ctx context.Context, coffeeBean *model.UserCoffeeBean) error
}

type UserCoffeeBeanWriter struct {
	db *ent.Client
}

var _ UserCoffeeBean = (*UserCoffeeBeanWriter)(nil)

func (impl *UserCoffeeBeanWriter) Create(ctx context.Context, userCoffeeBean *model.UserCoffeeBean, user *model.User) error {
	now := time.Now().UTC()
	ucb, err := impl.db.UserCoffeeBean.Create().
		SetStatus(int32(userCoffeeBean.Status.Num())).
		SetUserID(int32(userCoffeeBean.User.ID)).
		SetName(userCoffeeBean.Name).
		SetFarmName(userCoffeeBean.FarmName).
		SetCountry(userCoffeeBean.Country).
		SetRoastDegree(userCoffeeBean.RoastDegree.String()).
		SetRoastedAt(userCoffeeBean.RoastedAt).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return err
	}
	*userCoffeeBean = model.NewUserCoffeeBean(ucb)
	return nil
}

func (impl *UserCoffeeBeanWriter) DeleteByID(ctx context.Context, userCoffeeBean *model.UserCoffeeBean) error {
	if _, err := impl.db.UserCoffeeBean.
		Update().
		SetStatus(int32(model.CoffeeBeanStatusDeleted.Num())).
		Where(usercoffeebean.IDEQ(int32(userCoffeeBean.ID))).
		Save(ctx); err != nil {
		return err
	}
	return nil
}
