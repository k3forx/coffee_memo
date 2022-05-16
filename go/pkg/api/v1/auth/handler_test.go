package auth_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"

	"github.com/k3forx/coffee_memo/pkg/api/v1/auth"
	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/result"
	usecase "github.com/k3forx/coffee_memo/pkg/usecase/auth"
	auth_helper "github.com/k3forx/coffee_memo/test/auth"
)

func TestHandler_SignUp(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup              func(ctrl *gomock.Controller) *usecase.MockUsecase
		expectedStatusCode int
		expectedBody       map[string]interface{}
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) *usecase.MockUsecase {
				uc := usecase.NewMockUsecase(ctrl)

				uc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(result.OK())
				return uc
			},
			expectedStatusCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"status":  "success",
			},
		},
		"usecase_returns_error": {
			setup: func(ctrl *gomock.Controller) *usecase.MockUsecase {
				usecase := usecase.NewMockUsecase(ctrl)
				usecase.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(result.New(result.CodeForbidden, "既に使用されているメールアドレスです。"))
				return usecase
			},
			expectedStatusCode: http.StatusForbidden,
			expectedBody: map[string]interface{}{
				"message": "既に使用されているメールアドレスです。",
				"status":  "forbidden",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := inject.NewMockInjector(ctrl)
			h := auth.NewHandler(injector)
			uc := c.setup(ctrl)
			h.SetUsecase(uc)

			testReq := httptest.NewRequest(http.MethodPost, "/v1/auth/signup", nil)
			testRes := httptest.NewRecorder()

			ctx := echo.New().NewContext(testReq, testRes)
			err := h.SignUp(ctx)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}

			if diff := cmp.Diff(c.expectedStatusCode, testRes.Code); diff != "" {
				t.Errorf("SignUp() status code mismatch (-want +got):\n%s", diff)
			}

			body := map[string]interface{}{}
			_ = json.NewDecoder(testRes.Body).Decode(&body)
			if diff := cmp.Diff(c.expectedBody, body); diff != "" {
				t.Errorf("SignUp() body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandler_LogIn(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup              func(ctrl *gomock.Controller) *usecase.MockUsecase
		expectedStatusCode int
		expectedBody       map[string]interface{}
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) *usecase.MockUsecase {
				uc := usecase.NewMockUsecase(ctrl)

				uc.EXPECT().LogIn(gomock.Any(), gomock.Any()).
					Return(
						&usecase.LogInOutput{
							User: model.User{
								ID:    1,
								Email: "test@test.com",
							},
						},
						result.OK(),
					)
				return uc
			},
			expectedStatusCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"status":  "success",
			},
		},
		"usecase_returns_error": {
			setup: func(ctrl *gomock.Controller) *usecase.MockUsecase {
				usecase := usecase.NewMockUsecase(ctrl)
				usecase.EXPECT().LogIn(gomock.Any(), gomock.Any()).
					Return(nil, result.New(result.CodeBadRequest, "パスワードが違います。"))
				return usecase
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "パスワードが違います。",
				"status":  "bad request",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := inject.NewMockInjector(ctrl)
			h := auth.NewHandler(injector)
			uc := c.setup(ctrl)
			h.SetUsecase(uc)

			e, err := auth_helper.InitSessionStore(echo.New())
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			r := e.Router()
			r.Add(http.MethodPost, "/v1/auth/login", h.LogIn)

			apitest.New().Handler(e).Post("/v1/auth/login").
				Expect(t).Status(c.expectedStatusCode).End()
		})
	}
}
