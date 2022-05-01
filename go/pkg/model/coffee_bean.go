package model

import (
	"time"

	"github.com/k3forx/coffee_memo/pkg/ent"
)

type RoastDegree string

const (
	RoastDegreeUnknown  RoastDegree = ""
	RoastDegreeLight    RoastDegree = "light"
	RoastDegreeChinamon RoastDegree = "chinamon"
	RoastDegreeMedium   RoastDegree = "medium"
	RoastDegreeHigh     RoastDegree = "high"
	RoastDegreeCity     RoastDegree = "city"
	RoastDegreeFullCity RoastDegree = "fullcity"
	RoastDegreeFrench   RoastDegree = "french"
	RoastDegreeItalian  RoastDegree = "italian"
)

var roastDegreeStringMaps = map[RoastDegree]string{
	RoastDegreeUnknown:  "",
	RoastDegreeLight:    "light",
	RoastDegreeChinamon: "chinamon",
	RoastDegreeMedium:   "medium",
	RoastDegreeHigh:     "high",
	RoastDegreeCity:     "city",
	RoastDegreeFullCity: "fullcity",
	RoastDegreeFrench:   "french",
	RoastDegreeItalian:  "italian",
}

var roastDegreeLocalizedStringMaps = map[RoastDegree]string{
	RoastDegreeUnknown:  "",
	RoastDegreeLight:    "ライト",
	RoastDegreeChinamon: "シナモン",
	RoastDegreeMedium:   "ミディアム",
	RoastDegreeHigh:     "ハイ",
	RoastDegreeCity:     "シティ",
	RoastDegreeFullCity: "フルシティ",
	RoastDegreeFrench:   "フレンチ",
	RoastDegreeItalian:  "イタリアン",
}

func (rd RoastDegree) String() string {
	return string(rd)
}

func (rd RoastDegree) LocalizedString() string {
	for t, localizedString := range roastDegreeLocalizedStringMaps {
		if rd == t && rd.Valid() {
			return localizedString + "ロースト"
		}
	}
	return roastDegreeLocalizedStringMaps[RoastDegreeUnknown]
}

func (rd RoastDegree) Valid() bool {
	return rd != RoastDegreeUnknown
}

func NewRoastDegree(str string) RoastDegree {
	for t, s := range roastDegreeStringMaps {
		if str == s {
			return t
		}
	}
	return RoastDegreeUnknown
}

type CoffeeBeanStatus int

const (
	CoffeeBeanStatusActive CoffeeBeanStatus = iota
	CoffeeBeanStatusDeleted
)

func (cbs CoffeeBeanStatus) Num() int {
	return int(cbs)
}

type CoffeeBean struct {
	ID          int
	Status      CoffeeBeanStatus
	Name        string
	FarmName    string
	Country     string
	User        User
	RoastDegree RoastDegree
	RoastedAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (cb *CoffeeBean) Exists() bool {
	return cb.ID > 0
}

func NewCoffeeBean(e *ent.CoffeeBean) CoffeeBean {
	return CoffeeBean{
		ID:          int(e.ID),
		Name:        e.Name,
		FarmName:    e.FarmName,
		Country:     e.Country,
		User:        User{},
		RoastDegree: NewRoastDegree(e.RoastDegree),
		RoastedAt:   e.RoastedAt,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func NewCoffeeBeanWithUser(e *ent.UsersCoffeeBean) CoffeeBean {
	cb := e.Edges.CoffeeBean
	u := e.Edges.User

	flags := make([]UserFlag, len(userFlagsList))
	for i, f := range userFlagsList {
		if u.Flags&f.Num() == f.Num() {
			flags[i] = f
		}
	}
	return CoffeeBean{
		ID:       int(e.ID),
		Name:     cb.Name,
		FarmName: cb.FarmName,
		Country:  cb.Country,
		User: User{
			ID:        int(u.ID),
			Username:  u.Username,
			Email:     u.Email,
			Flags:     flags,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		RoastDegree: NewRoastDegree(cb.RoastDegree),
		RoastedAt:   cb.RoastedAt,
		CreatedAt:   cb.CreatedAt,
		UpdatedAt:   cb.UpdatedAt,
	}

}
