package main

import (
	"log"
	"net/http"
)

func main() {
	certPath := "server.pem"
	keyPath := "server.key"
	api := NewAPI(certPath, keyPath)

	http.HandleFunc("/hello", api.Hello.ServeHTTP)
	http.HandleFunc("/tokens", api.Tokens.ServeHTTP)

	// TODO: Wrap with authentication and authorization middleware
	http.HandleFunc("/users", api.Users.ServeHTTP)

	err := http.ListenAndServeTLS(":3000", certPath, keyPath, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
