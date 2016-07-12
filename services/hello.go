package services

import "net/http"

// HelloService provides a SayHello method
type HelloService interface {
	SayHello() string
}

// NewHelloService creates a new hello world service
func NewHelloService() HelloService {
	return &helloService{}
}

type helloService struct {
	H func(http.ResponseWriter, *http.Request)
}

// SayHello says hello for the hello service
func (h *helloService) SayHello() string {
	return "hello, world!"
}
