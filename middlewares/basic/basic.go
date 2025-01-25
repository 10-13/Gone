package basic_middleware

import (
	"Gone/core"
)

type BasicMiddleware struct {
}

func (self *BasicMiddleware) Handle(context *core.ITransferData) {
	(*context).SetResponseBody("Hello world!")
}
