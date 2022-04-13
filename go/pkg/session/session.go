package session

import (
	"github.com/gorilla/sessions"
	"github.com/k3forx/coffee_memo/pkg/model"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const sessionKey = "session"

var defaultOptions = &sessions.Options{
	MaxAge:   86400 * 7,
	HttpOnly: true,
}

func NewSession(c echo.Context) (*Session, error) {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return nil, err
	}
	sess.Options = defaultOptions
	return &Session{session: sess}, nil
}

type Session struct {
	session *sessions.Session
}

func (s *Session) SetSessionUser(u *model.User) {
	s.session.Values["id"] = u.ID
	s.session.Values["email"] = u.Email
}

func (s *Session) GetSessionUser() model.User {
	if s.session == nil {
		return model.User{}
	}

	id, ok := s.session.Values["id"].(int)
	if !ok {
		return model.User{}
	}

	email, ok := s.session.Values["email"].(string)
	if !ok {
		return model.User{}
	}
	return model.User{ID: id, Email: email}
}

func (s *Session) Save(c echo.Context) error {
	if err := s.session.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	return nil
}

func GetSession(c echo.Context) (*Session, error) {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return nil, err
	}
	return &Session{session: sess}, nil
}
