package access

type HelloApp struct {
}

func NewHelloApp() *HelloApp {
	return &HelloApp{}
}

func (h *HelloApp) SayHello() (interface{}, error) {
	return "Hello, this is gin-practice!", nil
}
