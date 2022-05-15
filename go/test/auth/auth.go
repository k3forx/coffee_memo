package auth

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func InitSessionStore(e *echo.Echo) (*echo.Echo, error) {
	store, err := sessions.NewRedisStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		return nil, err
	}
	e.Use(session.MiddlewareWithConfig(session.Config{Store: store}))
	return e, nil
}
