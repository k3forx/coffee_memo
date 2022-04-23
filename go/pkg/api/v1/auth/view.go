package auth

import (
	"net/http"
)

type LogInView struct {
	Session string `json:"session"`
}

func newLogInView(cookies *map[string]*http.Cookie) *LogInView {
	var session string
	for _, c := range *cookies {
		if c.Name == "session" {
			session = c.Value
		}
	}
	return &LogInView{Session: session}
}
