package user_coffee_bean

import (
	"errors"
	"time"

	"github.com/k3forx/coffee_memo/pkg/model"
)

type GetAllByUserIDInput struct {
	UserID int
}

type GetByIDInput struct {
	UserCoffeeBeanID int
}

type CreateInput struct {
	UserId      int
	Name        string
	FarmName    string
	Country     string
	RoastDegree model.RoastDegree
	RoastedAt   time.Time
}

type EditByIDInput struct {
	UserID           int
	UserCoffeeBeanID int
	Name             string
	FarmName         string
	Country          string
	RoastDegree      model.RoastDegree
	RoastedAt        time.Time
}

func (in EditByIDInput) validate() error {
	switch {
	case in.Name == "":
		return errors.New("名前を入力してください")
	}
	return nil
}

func (in EditByIDInput) Model() model.UserCoffeeBean {
	return model.UserCoffeeBean{
		ID: in.UserCoffeeBeanID,
		User: model.User{
			ID: in.UserID,
		},
		Name:        in.Name,
		FarmName:    in.FarmName,
		Country:     in.Country,
		RoastDegree: in.RoastDegree,
		RoastedAt:   in.RoastedAt,
	}
}

type DeleteByIDInput struct {
	UserID       int
	CoffeeBeanID int
}
