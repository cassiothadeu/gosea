package main

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/komand/gosea/handlers"
	"github.com/komand/gosea/services"
)

// API holds the api handlers
type API struct {
	encryptionKey []byte
	AclService    services.ACLService

	Hello  *handlers.Hello
	Tokens *handlers.Tokens
	Users  *handlers.Users
}

// NewAPI creates a new API
func NewAPI(certPath, keyPath string) *API {
	aclService := services.NewACLService()
	tokenService := services.NewTokenService()
	userService := services.NewUserService()
	helloService := services.NewHelloService()

	return &API{
		AclService: aclService,
		Tokens:     handlers.NewTokens(tokenService),
		Hello:      handlers.NewHello(helloService),
		Users:      handlers.NewUsers(userService),
	}
}

// Middleware
func (a *API) getToken(token *jwt.Token) (interface{}, error) {
	return []byte("secret"), nil
}

// Authenticate provides Authentication middleware for handlers
func (a *API) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string

		// Get token from the Authorization header
		// format: Authorization: Bearer <token>
		tokens, ok := r.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
		}

		// If the token is empty...
		if token == "" {
			// If we get here, the required token is missing
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		// Now parse the token
		parsedToken, err := jwt.Parse(token, a.getToken)
		if err != nil {
			// Could not parse the token
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		if jwt.SigningMethodHS256.Alg() != parsedToken.Header["alg"] {
			// Could not validate token algorithm
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		if !parsedToken.Valid {
			// Token is invalid
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		// Everything worked! Set the user in the context.
		context.Set(r, "user", parsedToken)

		next.ServeHTTP(w, r)
	})
}

// Authorize provides authorization middleware for our handlers
func (a *API) Authorize(permissions ...services.Permission) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := context.Get(r, "user").(*services.User)

			if user == nil {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			}

			for _, permission := range permissions {
				if err := a.AclService.CheckPermission(user, permission); err != nil {
					http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

// SecureHeaders adds secure headers to the API
func (a *API) SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Add security headers here
		next.ServeHTTP(w, r)
	})
}
