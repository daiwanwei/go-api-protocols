package middlewares

import (
	"context"
	"go-api-protocols/business/services"
	"go-api-protocols/utils/security"
	"net/http"
	"strings"
)

type AuthenticateMiddleware interface {
	Authenticate() func(http.Handler) http.Handler
}

type authenticateMiddleware struct {
	auth AuthService
}

func NewAuthenticateMiddleware() (middleware AuthenticateMiddleware, err error) {
	service, err := services.GetService()
	if err != nil {
		return nil, err
	}
	return &authenticateMiddleware{auth: service.Auth}, nil
}

func (middleware *authenticateMiddleware) Authenticate() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientToken := r.Header.Get("Authorization")
			if clientToken == "" {
				next.ServeHTTP(w, r)
				return
			}

			extractedToken := strings.Split(clientToken, "Bearer ")

			if len(extractedToken) == 2 {
				clientToken = strings.TrimSpace(extractedToken[1])
			} else {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			name, err := security.ExtractName(clientToken)
			if err != nil {
				return
			}
			user, err := middleware.auth.FindAuthByName(context.Background(), name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if user == nil {
				http.Error(w, "user not found", http.StatusUnauthorized)
				return
			}

			isValid, err := security.ValidateToken(clientToken, user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			if !isValid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			authentication := &authentication{user.GetName(), user.GetAuthorities()}
			ctx := context.WithValue(r.Context(), authenticationKey, authentication)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

type contextKey string

var authenticationKey = contextKey("Authentication")

func GetAuthenticationForContext(ctx context.Context) security.Authentication {
	if auth, ok := ctx.Value(authenticationKey).(security.Authentication); !ok {
		return nil
	} else {
		return auth
	}
}

////type AuthKey string
//
//var AuthKey="Authentication"

type AuthService interface {
	FindAuthByName(ctx context.Context, name string) (security.Authentication, error)
}

type authentication struct {
	Name        string
	Authorities []string
}

func (auth *authentication) GetName() string {
	return auth.Name
}

func (auth *authentication) GetAuthorities() []string {
	return auth.Authorities
}
