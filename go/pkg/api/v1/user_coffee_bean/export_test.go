package user_coffee_bean

import (
	"github.com/k3forx/coffee_memo/pkg/usecase/user_coffee_bean"
)

func (h *Handler) SetUsecase(uc user_coffee_bean.Usecase) {
	h.usecase = uc
}
