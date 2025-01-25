package core

import (
	. "Gone/custom_types"
)

type ITransferData interface {
	GetRequestHeaders() Headers

	SetResponseHeader(name HeaderName, value HeaderValue) HeaderValue
	GetResponseHeader(name HeaderName) (HeaderValue, error)
	RemoveResponseHeader(name HeaderName) HeaderValue

	GetRequsetBody() string

	SetResponseBody(body string)
	GetResponseBody() string

	EndTransfer()
}

type ITransferConnection interface {
	NextData() ITransferData
	IsClosed() bool

	CloseConnection()
}

type ITransferer interface {
	NextConnection() ITransferConnection
}
