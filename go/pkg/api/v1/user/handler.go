package user

import (
	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/presenter"
	"github.com/k3forx/coffee_memo/pkg/result"
	"github.com/k3forx/coffee_memo/pkg/session"
	"github.com/k3forx/coffee_memo/pkg/usecase/user"
	"github.com/labstack/echo/v4"
)

func NewHandler(injector inject.Injector) *Handler {
	return &Handler{
		usecase: user.NewUsecase(injector),
	}
}

type Handler struct {
	usecase user.Usecase
}

func Route(r *echo.Group, injector inject.Injector) {
	h := NewHandler(injector)
	r.GET("/me", h.GetMe)
}

func (h Handler) GetMe(c echo.Context) error {
	s, err := session.GetSession(c)
	if err != nil {
		return presenter.Error(c, result.Error())
	}
	sessionUser := s.GetSessionUser()

	in := user.GetByIDInput{
		UserID: sessionUser.ID,
	}
	out, res := h.usecase.GetByID(c.Request().Context(), in)
	if !res.IsOK() {
		return presenter.Error(c, res)
	}

	return presenter.JSON(c, newGetByIDView(out))
}
