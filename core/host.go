package core

struct IMiddleware interface {
	Handle(context ITransferData)
}

struct IHost interface {
	Setup(trs ITransferer, hnd IMiddleware) bool
	Start() bool
	Stop() bool
}