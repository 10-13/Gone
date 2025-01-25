package basic_middleware

import (
	"Gone/core"
	"Gone/custom_types"
	"strconv"
)

type BasicMiddleware struct {
}

func (self *BasicMiddleware) Handle(context core.ITransferData) {
	context.SetResponseBody("Hello world!")
	context.SetResponseHeader("Content-Length", custom_types.HeaderValue(strconv.Itoa(len("Hello world!"))))
}
