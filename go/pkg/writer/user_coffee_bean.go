package writer

import (
	"context"

	"github.com/k3forx/coffee_memo/pkg/ent"
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

func (impl *UserCoffeeBeanWriter) Create(ctx context.Context, coffeeBean *model.UserCoffeeBean, user *model.User) error {
	// now := time.Now().UTC()

	// tx, err := impl.db.Tx(ctx)
	// if err != nil {
	// 	return err
	// }

	// b, err := tx.CoffeeBean.Create().
	// 	SetName(coffeeBean.Name).
	// 	SetFarmName(coffeeBean.FarmName).
	// 	SetCountry(coffeeBean.Country).
	// 	SetRoastDegree(coffeeBean.RoastDegree.String()).
	// 	SetRoastedAt(coffeeBean.RoastedAt).
	// 	SetCreatedAt(now).
	// 	SetUpdatedAt(now).
	// 	Save(ctx)
	// if err != nil {
	// 	return transaction.Rollback(tx, fmt.Errorf("failed to create coffee beans: %w", err))
	// }
	// *coffeeBean = model.NewCoffeeBean(b)

	// _, err = tx.UsersCoffeeBean.Create().
	// 	SetStatus(int32(model.CoffeeBeanStatusActive.Num())).
	// 	SetUserID(int32(user.ID)).
	// 	SetCoffeeBeanID(int32(coffeeBean.ID)).
	// 	SetCreatedAt(now).
	// 	SetUpdatedAt(now).
	// 	Save(ctx)
	// if err != nil {
	// 	return transaction.Rollback(tx, fmt.Errorf("failed to create users coffee beans: %w", err))
	// }

	// return tx.Commit()
	return nil
}

func (impl *UserCoffeeBeanWriter) DeleteByID(ctx context.Context, coffeeBean *model.UserCoffeeBean) error {
	// if _, err := impl.db.UsersCoffeeBean.
	// 	Update().
	// 	SetStatus(int32(model.CoffeeBeanStatusDeleted.Num())).
	// 	Where(userscoffeebean.CoffeeBeanIDEQ(int32(coffeeBean.ID))).
	// 	Save(ctx); err != nil {
	// 	return err
	// }
	return nil
}
