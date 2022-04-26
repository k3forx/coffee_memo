package coffee_bean

import (
	"time"

	"github.com/k3forx/coffee_memo/pkg/model"
)

type CreateInput struct {
	UserId      int
	Name        string
	FarmName    string
	Country     string
	RoastDegree model.RoastDegree
	RoastedAt   time.Time
}
