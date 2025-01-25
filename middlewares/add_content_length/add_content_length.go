package add_content_length_middleware

import (
	"Gone/core"
	"Gone/custom_types"
	router_middleware "Gone/middlewares/router"
	"strconv"
)

type AddContentLengthMiddleware struct {
}

func (self *AddContentLengthMiddleware) Handle(context *core.ITransferData) {
	underlying := router_middleware.RouterMiddleware{}
	underlying.Handle(context)
	(*context).SetResponseHeader(
		"Content-Length",
		custom_types.HeaderValue(strconv.Itoa(len((*context).GetResponseBody()))),
	)
}
