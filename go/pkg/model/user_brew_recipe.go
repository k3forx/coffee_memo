package model

import (
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
)

type BrewRecipeStatus int

const (
	BrewRecipeStatusUnknown BrewRecipeStatus = iota
	BrewRecipeStatusCreated
)

type CoffeeBeanGrind string

const (
	CoffeeBeanGrindUnknown CoffeeBeanGrind = "unknown"
	CoffeeBeanGrindCoarse  CoffeeBeanGrind = "coarse"
)

var coffeeBeanGrindStringMap = map[CoffeeBeanGrind]string{
	CoffeeBeanGrindUnknown: "unknown",
	CoffeeBeanGrindCoarse:  "coarse",
}

func (m CoffeeBeanGrind) String() string {
	return string(m)
}

func NewCoffeeBeanGrind(str string) CoffeeBeanGrind {
	for k, v := range coffeeBeanGrindStringMap {
		if str == v {
			return k
		}
	}
	return CoffeeBeanGrindUnknown
}

type UserBrewRecipe struct {
	ID               int
	Status           BrewRecipeStatus
	User             User
	UserCoffeeBean   UserCoffeeBean
	CoffeeBeanWeight float64
	CoffeeBeanGrind  CoffeeBeanGrind
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
	DeletedAt        time.Time
}

func (m *UserBrewRecipe) Exists() bool {
	return m.ID > 0
}

func NewUserBrewRecipe(e *ent.UserBrewRecipe) UserBrewRecipe {
	u := e.Edges.User
	var user User
	if u != nil {
		user = NewUser(u)
	}
	user.ID = int(e.UserID)

	ucb := e.Edges.UserCoffeeBean
	var userCoffeeBean UserCoffeeBean
	if ucb != nil {
		userCoffeeBean = NewUserCoffeeBean(ucb)
	}
	userCoffeeBean.ID = int(e.UserCoffeeBeanID)

	userBrewRecipe := UserBrewRecipe{
		ID:               int(e.ID),
		Status:           BrewRecipeStatus(e.Status),
		User:             user,
		UserCoffeeBean:   userCoffeeBean,
		CoffeeBeanWeight: e.CoffeeBeanWeight,
		CoffeeBeanGrind:  NewCoffeeBeanGrind(e.CoffeeBeanGrind),
		LiquidWeight:     e.LiquidWeight,
		Temperature:      e.Temperature,
		StepOne:          e.StepOne,
		StepTwo:          e.StepTwo,
		StepThree:        e.StepThree,
		StepFour:         e.StepFour,
		StepFive:         e.StepFive,
		Memo:             e.Memo,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
		DeletedAt:        e.DeletedAt,
	}

	return userBrewRecipe
}
