package user_coffee_bean_test

import (
	"encoding/json"
	"fmt"
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
	auth_helper "github.com/k3forx/coffee_memo/test/auth"
	"github.com/labstack/echo/v4"
	"github.com/steinfletcher/apitest"
)

func TestHandler_GetAllByUserID(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		setup        func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase
		expectedCode int
		expectedBody map[string]interface{}
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				out := &userCoffeeBeanUsecase.GetAllByUserIDOutput{
					CoffeeBeans: []model.UserCoffeeBean{
						{
							ID:          1,
							Name:        "name 1",
							FarmName:    "farm 1",
							Country:     "country 1",
							RoastDegree: model.RoastDegreeChinamon,
							RoastedAt:   time.Date(2022, time.January, 10, 1, 0, 0, 0, time.UTC),
						},
						{
							ID:          2,
							Name:        "name 2",
							FarmName:    "farm 2",
							Country:     "country 2",
							RoastDegree: model.RoastDegreeLight,
							RoastedAt:   time.Date(2022, time.January, 20, 2, 0, 0, 0, time.UTC),
						},
					},
				}
				u.EXPECT().GetAllByUserID(gomock.Any(), gomock.Any()).Return(out, result.OK())
				return u
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"coffeeBeans": []interface{}{
					map[string]interface{}{
						"id":          float64(1),
						"name":        "name 1",
						"farmName":    "farm 1",
						"country":     "country 1",
						"roastDegree": "シナモンロースト",
						"roastedAt":   "2022-01-10",
					},
					map[string]interface{}{
						"id":          float64(2),
						"name":        "name 2",
						"farmName":    "farm 2",
						"country":     "country 2",
						"roastDegree": "ライトロースト",
						"roastedAt":   "2022-01-20",
					},
				},
			},
		},
		"usecase_returns_error": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				u.EXPECT().GetAllByUserID(gomock.Any(), gomock.Any()).
					Return(nil, result.New(result.CodeNotFound, "コーヒー豆が見つかりません"))
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

			ctrl := gomock.NewController(t)
			injector := inject.NewMockInjector(ctrl)
			h := user_coffee_bean.NewHandler(injector)
			u := c.setup(ctrl)
			h.SetUsecase(u)

			e, err := auth_helper.InitSessionStore(echo.New())
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			r := e.Router()
			r.Add(http.MethodGet, "/v1/coffee-beans", h.GetAllByUserID)

			expectedBody, err := json.Marshal(c.expectedBody)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			apitest.New().Handler(e).Get("/v1/coffee-beans").
				Expect(t).Status(c.expectedCode).Body(string(expectedBody)).End()
		})
	}
}

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

func TestHandler_Create(t *testing.T) {
	t.Parallel()

	const endpoint = "/v1/coffee-beans"

	cases := map[string]struct {
		setup        func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase
		body         map[string]interface{}
		expectedCode int
		expectedBody map[string]interface{}
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				u.EXPECT().Create(gomock.Any(), gomock.Any()).Return(result.OK())
				return u
			},
			body: map[string]interface{}{
				"roastedAt": "2022-01-09",
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"status":  "success",
			},
		},
		"invalid_time_format": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				return userCoffeeBeanUsecase.NewMockUsecase(ctrl)
			},
			body: map[string]interface{}{
				"roastedAt": "invalid",
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "bad request",
				"status":  "bad request",
			},
		},
		"usecase_returns_error": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				u.EXPECT().Create(gomock.Any(), gomock.Any()).Return(result.Error())
				return u
			},
			body: map[string]interface{}{
				"roastedAt": "2022-01-09",
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"message": "Error!",
				"status":  "error",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := inject.NewMockInjector(ctrl)
			h := user_coffee_bean.NewHandler(injector)
			u := c.setup(ctrl)
			h.SetUsecase(u)

			e, err := auth_helper.InitSessionStore(echo.New())
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			r := e.Router()
			r.Add(http.MethodPost, endpoint, h.Create)

			body, err := json.Marshal(c.body)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			expectedBody, err := json.Marshal(c.expectedBody)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			apitest.New().Handler(e).Post(endpoint).JSON(body).
				Expect(t).Status(c.expectedCode).Body(string(expectedBody)).End()
		})
	}
}

func TestHandler_EditByID(t *testing.T) {
	t.Parallel()

	const baseEndpoint = "/v1/coffee-beans"

	cases := map[string]struct {
		setup        func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase
		id           string
		body         map[string]interface{}
		expectedCode int
		expectedBody map[string]interface{}
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				u.EXPECT().EditByID(gomock.Any(), gomock.Any()).Return(result.OK())
				return u
			},
			id: "1",
			body: map[string]interface{}{
				"roastedAt": "2022-01-09",
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"status":  "success",
			},
		},
		"invalid_id": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				return userCoffeeBeanUsecase.NewMockUsecase(ctrl)
			},
			id: "invalid-id",
			body: map[string]interface{}{
				"roastedAt": "2022-01-09",
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "bad request",
				"status":  "bad request",
			},
		},
		"invalid_time_format": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				return userCoffeeBeanUsecase.NewMockUsecase(ctrl)
			},
			id: "1",
			body: map[string]interface{}{
				"roastedAt": "invalid",
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "bad request",
				"status":  "bad request",
			},
		},
		"usecase_returns_error": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				u.EXPECT().EditByID(gomock.Any(), gomock.Any()).Return(result.Error())
				return u
			},
			id: "1",
			body: map[string]interface{}{
				"roastedAt": "2022-01-09",
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"message": "Error!",
				"status":  "error",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := inject.NewMockInjector(ctrl)
			h := user_coffee_bean.NewHandler(injector)
			u := c.setup(ctrl)
			h.SetUsecase(u)

			e, err := auth_helper.InitSessionStore(echo.New())
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			r := e.Router()
			r.Add(http.MethodPut, fmt.Sprintf("%s/:id", baseEndpoint), h.EditByID)

			body, err := json.Marshal(c.body)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			expectedBody, err := json.Marshal(c.expectedBody)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			apitest.New().Handler(e).Put(fmt.Sprintf("%s/%s", baseEndpoint, c.id)).JSON(body).
				Expect(t).Status(c.expectedCode).Body(string(expectedBody)).End()
		})
	}
}

func TestHandler_DeleteByID(t *testing.T) {
	t.Parallel()

	const baseEndpoint = "/v1/coffee-beans"

	cases := map[string]struct {
		setup        func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase
		id           string
		body         map[string]interface{}
		expectedCode int
		expectedBody map[string]interface{}
	}{
		"success": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				u.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(result.OK())
				return u
			},
			id: "1",
			body: map[string]interface{}{
				"roastedAt": "2022-01-09",
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "success",
				"status":  "success",
			},
		},
		"invalid_id": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				return userCoffeeBeanUsecase.NewMockUsecase(ctrl)
			},
			id: "invalid-id",
			body: map[string]interface{}{
				"roastedAt": "2022-01-09",
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "bad request",
				"status":  "bad request",
			},
		},
		"usecase_returns_error": {
			setup: func(ctrl *gomock.Controller) userCoffeeBeanUsecase.Usecase {
				u := userCoffeeBeanUsecase.NewMockUsecase(ctrl)
				u.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(result.Error())
				return u
			},
			id: "1",
			body: map[string]interface{}{
				"roastedAt": "2022-01-09",
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"message": "Error!",
				"status":  "error",
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			injector := inject.NewMockInjector(ctrl)
			h := user_coffee_bean.NewHandler(injector)
			u := c.setup(ctrl)
			h.SetUsecase(u)

			e, err := auth_helper.InitSessionStore(echo.New())
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			r := e.Router()
			r.Add(http.MethodDelete, fmt.Sprintf("%s/:id", baseEndpoint), h.DeleteByID)

			body, err := json.Marshal(c.body)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			expectedBody, err := json.Marshal(c.expectedBody)
			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
			}
			apitest.New().Handler(e).Delete(fmt.Sprintf("%s/%s", baseEndpoint, c.id)).JSON(body).
				Expect(t).Status(c.expectedCode).Body(string(expectedBody)).End()
		})
	}
}
