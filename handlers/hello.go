package handlers

import (
	"net/http"

	"github.com/komand/gosea/services"
)

// Hello exposes an api for the hello service
type Hello struct {
	Service services.HelloService
}

// NewHello creates a new handler for hello
func NewHello(s services.HelloService) *Hello {
	return &Hello{s}
}

// Handler handles hello requests
func (h *Hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		s := h.Service.SayHello()
		w.Write([]byte(s))
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
