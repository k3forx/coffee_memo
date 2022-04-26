package coffee_bean

import (
	"github.com/k3forx/coffee_memo/pkg/usecase/coffee_bean"
)

type CoffeeBean struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FarmName    string `json:"farmName"`
	RoastDegree string `json:"roastDegree"`
	RoastedAt   string `json:"roastedAt"`
}

type CoffeeBeans struct {
	CoffeeBeans []CoffeeBean `json:"coffeeBeans"`
}

func newGetAllView(out *coffee_bean.GetAllOutput) CoffeeBeans {
	coffeeBeans := make([]CoffeeBean, len(out.CoffeeBeans))
	for i, cb := range out.CoffeeBeans {
		coffeeBeans[i] = CoffeeBean{
			ID:          cb.ID,
			Name:        cb.Name,
			FarmName:    cb.FarmName,
			RoastDegree: cb.RoastDegree.LocalizedString(),
		}
	}
	return CoffeeBeans{CoffeeBeans: coffeeBeans}
}
