package user_coffee_bean_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/k3forx/coffee_memo/pkg/api/v1/user_coffee_bean"
	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/k3forx/coffee_memo/pkg/result"
	userCoffeeBeanUsecase "github.com/k3forx/coffee_memo/pkg/usecase/user_coffee_bean"
	"github.com/labstack/echo/v4"
)

func TestHandler_GetByID(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		id           string
		setup        func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase
		expectedCode int
		expectedBody map[string]interface{}
	}{
		"success": {
			id: "1",
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				in := userCoffeeBeanUsecase.GetByIDInput{
					UserCoffeeBeanID: 1,
				}
				out := &userCoffeeBeanUsecase.GetByIDOutput{
					UserCoffeeBean: model.UserCoffeeBean{
						ID:          1,
						Status:      model.CoffeeBeanStatusActive,
						Name:        "test name",
						FarmName:    "test farm",
						Country:     "Japan",
						RoastDegree: model.RoastDegreeChinamon,
						RoastedAt:   time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
						CreatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2022, time.March, 10, 0, 0, 0, 0, time.UTC),
					},
				}
				u.EXPECT().GetByID(gomock.All(), in).Return(out, result.OK())
				return u
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"id":          float64(1),
				"name":        "test name",
				"farmName":    "test farm",
				"country":     "Japan",
				"roastDegree": "chinamon",
				"roastedAt":   "2022-03-01",
			},
		},
		"invalid_id": {
			id: "",
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				return userCoffeeBeanUsecase.NewMockUsecase(ctrl)
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "bad request",
				"status":  "bad request",
			},
		},
		"usecase_returns_error": {
			id: "1",
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				in := userCoffeeBeanUsecase.GetByIDInput{
					UserCoffeeBeanID: 1,
				}
				u.EXPECT().GetByID(gomock.All(), in).Return(nil, result.New(result.CodeNotFound, "コーヒー豆が見つかりません"))
				return u
			},
			expectedCode: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"message": "コーヒー豆が見つかりません",
				"status":  "not found",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			r := httptest.NewRequest(http.MethodGet, "/v1/coffee-beans/:id", nil)
			w := httptest.NewRecorder()
			echoContext := e.NewContext(r, w)
			echoContext.SetParamNames("id")
			echoContext.SetParamValues(c.id)

			ctrl := gomock.NewController(t)
			injector := inject.NewMockInjector(ctrl)

			h := user_coffee_bean.NewHandler(injector)
			u := c.setup(ctrl)
			h.SetUsecase(u)

			err := h.GetByID(echoContext)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			if diff := cmp.Diff(c.expectedCode, w.Code); diff != "" {
				t.Errorf("GetByID() code mismatch (-want +got):\n%s", diff)
			}

			var actualBody map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &actualBody)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			if diff := cmp.Diff(c.expectedBody, actualBody); diff != "" {
				t.Errorf("GetByID() body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
