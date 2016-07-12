package handlers

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/komand/gosea/services"
)

// Tokens exposes an API to the tokens service
type Tokens struct {
	Service services.TokenService
}

// NewTokens creates new handler for tokens
func NewTokens(s services.TokenService) *Tokens {
	return &Tokens{s}
}

// Handler will return tokens
func (t *Tokens) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		user, ok := context.Get(req, "user").(*services.User)
		if !ok {
			http.Error(w, "Failed to get user from request", http.StatusBadRequest)
		}
		token, err := t.Service.Get(user)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		}
		w.Write([]byte(token))
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

}
