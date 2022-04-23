package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/presenter"
	"github.com/k3forx/coffee_memo/pkg/result"
	"github.com/k3forx/coffee_memo/pkg/session"
	"github.com/k3forx/coffee_memo/pkg/usecase/auth"
)

func NewHandler(injector inject.Injector) *Handler {
	return &Handler{
		usecase: auth.NewUsecase(injector),
	}
}

type Handler struct {
	usecase auth.Usecase
}

func Route(r *echo.Group, injector inject.Injector) {
	h := NewHandler(injector)
	r.POST("/signup", h.SignUp)
	r.POST("/login", h.LogIn)
}

func (h Handler) SignUp(c echo.Context) error {
	var req SignUpRequest

	// Ignore error because we can catch errors by Validate method
	_ = c.Bind(&req)

	in := auth.SignUpInput{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	if res := h.usecase.SignUp(c.Request().Context(), in); !res.IsOK() {
		return presenter.Error(c, res)
	}

	return presenter.Success(c)
}

func (h Handler) LogIn(c echo.Context) error {
	var req LogInRequest

	// Ignore error because we can catch errors by validate method
	_ = c.Bind(&req)

	in := auth.LogInInput{
		Email:    req.Email,
		Password: req.Password,
	}

	out, res := h.usecase.LogIn(c.Request().Context(), in)
	if !res.IsOK() {
		return presenter.Error(c, res)
	}

	sess, _ := session.NewSession(c)
	sess.SetSessionUser(&out.User)
	if err := sess.Save(c); err != nil {
		return presenter.Error(c, result.Error())
	}

	cookies := c.Response().Writer.Header().Get("Set-Cookie")
	cookieMaps := map[string]*http.Cookie{}
	for _, c := range (&http.Request{Header: http.Header{"Cookie": {cookies}}}).Cookies() {
		cookieMaps[c.Name] = c
	}

	return presenter.JSON(c, newLogInView(&cookieMaps))
}
