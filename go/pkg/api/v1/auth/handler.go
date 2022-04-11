package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/presenter"
	"github.com/k3forx/coffee_memo/pkg/result"
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
	r.GET("/", h.Get)
}

func (h Handler) SignUp(ctx echo.Context) error {
	var req SignUpRequest

	// Ignore error because we can catch errors by Validate method
	_ = ctx.Bind(&req)

	// u := auth.NewUsecase(h.injector)
	// in := auth.SignUpInput{
	// 	Username: req.Username,
	// 	Email:    req.Email,
	// 	Password: req.Password,
	// }
	// if res := u.SignUp(ctx.Request().Context(), in); !res.IsOK() {
	// 	// return presenter.Error(ctx, res)
	// 	presenter.Error(ctx, res)
	// }
	sess, err := session.Get("session", ctx)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
	}
	sess.Options = &sessions.Options{
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	//テキトウな値
	sess.Values["foo"] = "bar"
	//ログインしました
	sess.Values["auth"] = true
	if err := sess.Save(ctx.Request(), ctx.Response()); err != nil {
		return presenter.Error(ctx, result.Error())
	}

	return ctx.NoContent(http.StatusOK)
	// return presenter.Success(ctx)
}

func (h Handler) Get(ctx echo.Context) error {
	sess, err := session.Get("session", ctx)
	if err != nil {
		return err
	}
	b, ok := sess.Values["auth"].(bool)
	if !ok {
		fmt.Println("not ok")
	}
	fmt.Printf("b: %+v\n", b)

	res, ok := sess.Values["foo"].(string)
	if !ok {
		fmt.Println("not ok")
	}
	fmt.Printf("res: %+v\n", res)

	return nil
}
