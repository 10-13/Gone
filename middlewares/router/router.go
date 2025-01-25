package router_middleware

import (
	"Gone/core"
	basic_middleware "Gone/middlewares/basic"
)

type RouterMiddleware struct {
}

func (self *RouterMiddleware) Handle(context *core.ITransferData) {
	handler := basic_middleware.BasicMiddleware{}
	handler.Handle(context)
}
