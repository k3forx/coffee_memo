package presenter

import (
	"net/http"

	"github.com/k3forx/coffee_memo/pkg/result"
	echo "github.com/labstack/echo/v4"
)

type v1SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func newV1SuccessResponse() *v1SuccessResponse {
	return &v1SuccessResponse{
		Status:  "success",
		Message: "success",
	}
}

type v1ErrResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func newV1ErrResponse(res *result.Result) *v1ErrResponse {
	return &v1ErrResponse{
		Status:  res.Code.String(),
		Message: res.Message,
	}
}

func Success(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, newV1SuccessResponse())
}

func JSON(ctx echo.Context, body interface{}) error {
	return ctx.JSON(http.StatusOK, body)
}

func BadRequest(ctx echo.Context, msg string) error {
	res := result.New(result.CodeBadRequest, msg)
	return ctx.JSON(res.Code.ToStatusCode(), newV1ErrResponse(res))
}

func Error(ctx echo.Context, res *result.Result) error {
	return ctx.JSON(res.Code.ToStatusCode(), newV1ErrResponse(res))
}
