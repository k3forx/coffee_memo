package user

import (
	"strconv"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/presenter"
	"github.com/k3forx/coffee_memo/pkg/usecase/user"
	echo "github.com/labstack/echo/v4"
)

func newHandler(injector inject.Injector) *Handler {
	return &Handler{
		injector: injector,
	}
}

type Handler struct {
	injector inject.Injector
}

func Route(r *echo.Group, injector inject.Injector) {
	h := newHandler(injector)
	r.GET("/:id", h.Get)
}

func (h Handler) Get(ctx echo.Context) error {
	u := user.NewUsecase(h.injector)
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return presenter.BadRequest(ctx, err.Error())
	}
	in := user.GetByIDInput{
		UserID: userID,
	}
	out, res := u.GetByID(ctx.Request().Context(), in)
	if !res.IsOK() {
		return presenter.Error(ctx, res)
	}

	return presenter.JSON(ctx, out)
}
