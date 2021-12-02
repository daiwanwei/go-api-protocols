package middlewares

import (
	"net/http"
)

type AuthMiddleware interface {
	Authenticate() func(http.Handler) http.Handler
}

type authMiddleware struct {
	authenticate AuthenticateMiddleware
}

func NewAuthMiddleware() (middleware AuthMiddleware, err error) {
	authenticate, err := NewAuthenticateMiddleware()
	if err != nil {
		return
	}
	return &authMiddleware{
		authenticate: authenticate,
	}, nil
}

func (middleware *authMiddleware) Authenticate() func(http.Handler) http.Handler {
	return middleware.authenticate.Authenticate()
}
