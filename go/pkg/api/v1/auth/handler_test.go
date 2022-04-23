package auth_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"

	"github.com/k3forx/coffee_memo/pkg/api/v1/auth"
	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/result"
	usecase "github.com/k3forx/coffee_memo/pkg/usecase/auth"
)

func TestHandler_SignUp(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		setup              func(ctrl *gomock.Controller) *usecase.MockAuthUsecase
		expectedStatusCode int
		expectedBody       map[string]interface{}
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) *usecase.MockAuthUsecase {
				usecase := usecase.NewMockAuthUsecase(ctrl)

				usecase.EXPECT().SignUp(gomock.Any(), gomock.Any()).Return(result.OK())
				return usecase
			},
			expectedStatusCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"status":  "success",
			},
		},
		"usecase_returns_error": {
			setup: func(ctrl *gomock.Controller) *usecase.MockAuthUsecase {
				usecase := usecase.NewMockAuthUsecase(ctrl)
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
