package user_brew_recipe

import (
	"time"

	"github.com/k3forx/coffee_memo/pkg/model"
)

type CreateInput struct {
	UserID           int
	UserCoffeeBeanID int
	CoffeeBeanWeight float64
	CoffeeBeanGrind  model.CoffeeBeanGrind
	LiquidWeight     float64
	Temperature      float64
	StepOne          string
	StepTwo          string
	StepThree        string
	StepFour         string
	StepFive         string
	Memo             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type GetByIDInput struct {
	UserID           int
	UserBrewRecipeID int
}
