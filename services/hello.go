package services

// HelloService provides a SayHello method
type HelloService interface {
	SayHello() string
}

// NewHelloService creates a new hello world service
func NewHelloService() HelloService {
	return &helloService{}
}

type helloService struct {
}

// SayHello says hello for the hello service
func (h *helloService) SayHello() string {
	return "hello, world!"
}
