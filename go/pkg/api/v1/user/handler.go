package user

import (
	"github.com/k3forx/coffee_memo/pkg/usecase/user"
	"github.com/labstack/echo/v4"
)

func newHandler() *Handler {
	return &Handler{}
}

type Handler struct {
}

func Route(r *echo.Group) {
	h := newHandler()
	r.GET("/", h.Get)
}

func (h Handler) Get(ctx echo.Context) error {
	u := user.NewUsecase()
	u.GetByID(ctx.Request().Context())
	return nil
}
