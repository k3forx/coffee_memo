package user

import "github.com/k3forx/coffee_memo/pkg/usecase/user"

func (h *Handler) SetUsecase(uc user.Usecase) {
	h.usecase = uc
}
