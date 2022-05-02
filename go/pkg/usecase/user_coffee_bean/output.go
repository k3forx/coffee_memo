package user_coffee_bean

import "github.com/k3forx/coffee_memo/pkg/model"

type GetAllByUserIDOutput struct {
	CoffeeBeans []model.UserCoffeeBean
}

type GetByIDOutput struct {
	UserCoffeeBean model.UserCoffeeBean
}
