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
		if rd == t {
			return localizedString + "ロースト"
		}
	}
	return roastDegreeLocalizedStringMaps[RoastDegreeUnknown]
}

func NewRoastDegree(str string) RoastDegree {
	for t, s := range roastDegreeStringMaps {
		if str == s {
			return t
		}
	}
	return RoastDegreeUnknown
}

type CoffeeBean struct {
	ID          int
	User        User
	Name        string
	FarmName    string
	Country     string
	RoastDegree RoastDegree
	RoastedAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCoffeeBean(e *ent.CoffeeBean) CoffeeBean {
	return CoffeeBean{
		ID:          int(e.ID),
		User:        User{},
		Name:        e.Name,
		FarmName:    e.FarmName,
		Country:     e.Country,
		RoastDegree: NewRoastDegree(e.RoastedDegree),
		RoastedAt:   e.RoastedAt,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
