package server

import (
	"fmt"

	"github.com/k3forx/coffee_memo/pkg/api/v1/user"
	"github.com/k3forx/coffee_memo/pkg/config"
	"github.com/labstack/echo/v4"
)

func NewServer() *Server {
	e := echo.New()
	registerRoute(e)
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

func registerRoute(e *echo.Echo) {
	v1API := e.Group("/v1")
	{
		v1APIUser := v1API.Group("/users")
		user.Route(v1APIUser)
	}
}