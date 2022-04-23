package auth

import (
	"github.com/k3forx/coffee_memo/pkg/usecase/auth"
)

func (h *Handler) SetUsecase(uc auth.Usecase) {
	h.usecase = uc
}
