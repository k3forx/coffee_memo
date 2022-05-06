package user_coffee_bean

import (
	"github.com/k3forx/coffee_memo/pkg/usecase/user_coffee_bean"
)

type UserCoffeeBean struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FarmName    string `json:"farmName"`
	Country     string `json:"country"`
	RoastDegree string `json:"roastDegree"`
	RoastedAt   string `json:"roastedAt"`
}

type UserCoffeeBeans struct {
	UserCoffeeBeans []UserCoffeeBean `json:"coffeeBeans"`
}

func newGetAllView(out *user_coffee_bean.GetAllByUserIDOutput) UserCoffeeBeans {
	userCoffeeBeans := make([]UserCoffeeBean, len(out.CoffeeBeans))
	for i, cb := range out.CoffeeBeans {
		userCoffeeBeans[i] = UserCoffeeBean{
			ID:          cb.ID,
			Name:        cb.Name,
			FarmName:    cb.FarmName,
			Country:     cb.Country,
			RoastDegree: cb.RoastDegree.LocalizedString(),
			RoastedAt:   cb.RoastedAt.Format("2006-01-02"),
		}
	}
	return UserCoffeeBeans{UserCoffeeBeans: userCoffeeBeans}
}

func newGetByIDView(out *user_coffee_bean.GetByIDOutput) UserCoffeeBean {
	ucb := out.UserCoffeeBean
	return UserCoffeeBean{
		ID:          ucb.ID,
		Name:        ucb.Name,
		FarmName:    ucb.FarmName,
		Country:     ucb.Country,
		RoastDegree: ucb.RoastDegree.String(),
		RoastedAt:   ucb.RoastedAt.Format("2006-01-02"),
	}
}
