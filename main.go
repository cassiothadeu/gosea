package main

import (
	"log"
	"net/http"

	"github.com/komand/gosea/services"
)

func main() {
	certPath := "server.pem"
	keyPath := "server.key"
	api := NewAPI(certPath, keyPath)

	http.Handle("/hello", api.Hello)
	http.Handle("/tokens", api.Tokens)

	http.Handle("/users", AddMiddleware(api.Users,
		api.Authenticate,
		api.Authorize(services.Permission("user_modify"))))

	err := http.ListenAndServeTLS(":3000", certPath, keyPath, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// AddMiddleware adds middleware to a Handler
func AddMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
