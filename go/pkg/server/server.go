package server

import (
	"fmt"
	"net/http"

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
	e.GET("/",
		func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!!!!!")
		},
	)
}
