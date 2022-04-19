package auth

import (
	"net/http"
)

type SignUpView struct {
	Session string `json:"session"`
}

func newSignUpView(cookies *map[string]*http.Cookie) *SignUpView {
	var session string
	for _, c := range *cookies {
		if c.Name == "session" {
			session = c.Value
		}
	}
	return &SignUpView{Session: session}
}
