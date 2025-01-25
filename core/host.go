package core

type IMiddleware interface {
	Handle(context *ITransferData)
}

type IHost interface {
	Setup(trs ITransferer, hnd IMiddleware) bool
	Start() bool
	Stop() bool
}
