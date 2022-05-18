package user_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/k3forx/coffee_memo/pkg/api/v1/user"
	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/result"
	userUsecase "github.com/k3forx/coffee_memo/pkg/usecase/user"
	authHelper "github.com/k3forx/coffee_memo/test/auth"
	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

const baseEndpoint = "/v1/users"

func TestHandler_GetMe(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup        func(ctrl *gomock.Controller) userUsecase.Usecase
		expectedCode int
		expectedBody map[string]interface{}
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) userUsecase.Usecase {
				u := userUsecase.NewMockUsecase(ctrl)
				out := &userUsecase.GetByIDOutput{
					User: model.User{
						ID:       1,
						Username: "test user",
						Email:    "test@example.com",
					},
				}
				u.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(out, result.OK())
				return u
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"id":       1,
				"username": "test user",
				"email":    "test@example.com",
			},
		},
		"usecase_returns_error": {
			setup: func(ctrl *gomock.Controller) userUsecase.Usecase {
				u := userUsecase.NewMockUsecase(ctrl)
				u.EXPECT().GetByID(gomock.Any(), gomock.Any()).
					Return(nil, result.New(result.CodeNotFound, "アカウントが見つかりません"))
				return u
			},
			expectedCode: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"status":  "not found",
				"message": "アカウントが見つかりません",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := inject.NewMockInjector(ctrl)
			h := user.NewHandler(injector)
			u := c.setup(ctrl)
			h.SetUsecase(u)

			e, err := authHelper.InitSessionStore(echo.New())
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			r := e.Router()
			r.Add(http.MethodGet, fmt.Sprintf("%s/me", baseEndpoint), h.GetMe)

			expectedBody, err := json.Marshal(c.expectedBody)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			apitest.New().Handler(e).Get(fmt.Sprintf("%s/me", baseEndpoint)).
				Expect(t).Status(c.expectedCode).Body(string(expectedBody)).End()
		})
	}
}
