package user

import (
	"fmt"
	"log"
	"strconv"

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
	r.GET("/:id", h.Get)
}

func (h Handler) Get(ctx echo.Context) error {
	u := user.NewUsecase()
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	in := user.GetByIDInput{
		UserID: userID,
	}
	out := u.GetByID(ctx.Request().Context(), in)
	fmt.Printf("output: %+v\n", out)

	return nil
}
