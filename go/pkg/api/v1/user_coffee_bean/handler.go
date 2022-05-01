package user_coffee_bean

import (
	"strconv"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/presenter"
	"github.com/k3forx/coffee_memo/pkg/result"
	"github.com/k3forx/coffee_memo/pkg/session"
	"github.com/k3forx/coffee_memo/pkg/usecase/user_coffee_bean"
	"github.com/labstack/echo/v4"
)

func NewHandler(injector inject.Injector) *Handler {
	return &Handler{
		usecase: user_coffee_bean.NewUsecase(injector),
	}
}

type Handler struct {
	usecase user_coffee_bean.Usecase
}

func Route(r *echo.Group, injector inject.Injector) {
	h := NewHandler(injector)
	r.GET("", h.GetAllByUserID)
	r.POST("", h.Create)
	r.DELETE("/:id", h.DeleteByID)
}

func (h Handler) GetAllByUserID(c echo.Context) error {
	s, err := session.GetSession(c)
	if err != nil {
		return presenter.Error(c, result.Error())
	}
	sessionUser := s.GetSessionUser()

	in := user_coffee_bean.GetAllByUserIDInput{
		UserID: sessionUser.ID,
	}
	out, res := h.usecase.GetAllByUserID(c.Request().Context(), in)
	if !res.IsOK() {
		return presenter.Error(c, res)
	}

	return presenter.JSON(c, newGetAllView(out))
}

func (h Handler) Create(c echo.Context) error {
	var req CreateRequest
	_ = c.Bind(&req)

	s, err := session.GetSession(c)
	if err != nil {
		return presenter.Error(c, result.Error())
	}
	sessionUser := s.GetSessionUser()

	in := user_coffee_bean.CreateInput{
		UserId:      sessionUser.ID,
		Name:        req.Name,
		FarmName:    req.FarmName,
		Country:     req.Country,
		RoastDegree: model.NewRoastDegree(req.RoastDegree),
	}

	if res := h.usecase.Create(c.Request().Context(), in); !res.IsOK() {
		return presenter.Error(c, res)
	}

	return presenter.Success(c)
}

func (h Handler) DeleteByID(c echo.Context) error {
	coffeeBeanID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return presenter.BadRequest(c, result.CodeBadRequest.String())
	}

	s, err := session.GetSession(c)
	if err != nil {
		return presenter.Error(c, result.Error())
	}
	sessionUser := s.GetSessionUser()

	in := user_coffee_bean.DeleteByIDInput{
		UserID:       sessionUser.ID,
		CoffeeBeanID: coffeeBeanID,
	}
	if res := h.usecase.DeleteByID(c.Request().Context(), in); !res.IsOK() {
		return presenter.Error(c, res)
	}

	return presenter.Success(c)
}
