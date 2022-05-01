package server

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/k3forx/coffee_memo/pkg/api/v1/auth"
	"github.com/k3forx/coffee_memo/pkg/api/v1/user"
	"github.com/k3forx/coffee_memo/pkg/api/v1/user_coffee_bean"
	"github.com/k3forx/coffee_memo/pkg/config"
	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func NewServer(injector inject.Injector) *Server {
	e := echo.New()
	registerRoute(e, injector)
	return &Server{server: e}
}

type Server struct {
	server *echo.Echo
}

func (s *Server) Start() error {
	if err := s.server.Start(
		fmt.Sprintf(":%d", config.GetMainAppPortNumber()),
	); err != nil {
		return err
	}
	return nil
}

func registerRoute(e *echo.Echo, injector inject.Injector) {
	v1API := e.Group("/v1")

	store, err := sessions.NewRedisStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	v1API.Use(session.MiddlewareWithConfig(session.Config{Store: store}))

	{
		v1APIAuth := v1API.Group("/auth")
		auth.Route(v1APIAuth, injector)
	}
	{
		v1APIUser := v1API.Group("/users")
		user.Route(v1APIUser, injector)
	}
	{
		v1APIUserCoffeeBean := v1API.Group("/coffee-beans")
		user_coffee_bean.Route(v1APIUserCoffeeBean, injector)
	}
}
