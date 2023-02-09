package access

import (
	"gin-practice/libs/error"
	"github.com/gin-gonic/gin"
)

type HelloApp struct {
}

func NewHelloApp() *HelloApp {
	return &HelloApp{}
}

func (h *HelloApp) SayHello(ctx *gin.Context) (interface{}, *error.Error) {
	return "Hi, this is LuWu ^_^", nil
}

func (h *HelloApp) SayFailed(ctx *gin.Context) (interface{}, *error.Error) {
	return nil, &error.Error{
		Code: 1001,
		Msg:  "Hi, i'm failed 0_0",
	}
}
