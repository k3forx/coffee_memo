package auth

import (
	echo "github.com/labstack/echo/v4"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/presenter"
	"github.com/k3forx/coffee_memo/pkg/usecase/auth"
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
	r.POST("/signup", h.SignUp)
}

func (h Handler) SignUp(ctx echo.Context) error {
	var req SignUpRequest

	// Ignore error because we can catch errors by Validate method
	_ = ctx.Bind(&req)

	u := auth.NewUsecase(h.injector)
	in := auth.SignUpInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if res := u.SignUp(ctx.Request().Context(), in); !res.IsOK() {
		return presenter.Error(ctx, res)
	}

	return presenter.Success(ctx)
}
