package core

type header_name = string
type header_value = string

type ITransferData interface {
	GetRequestHeaders() []string

	SetResponseHeader(name header_name, value header_value) header_value
	GetResponseHeader(name header_name) bool, header_value
	RemoveResponseHeader(name header_name) header_value

	GetRequsetBody() string

	SetResponseBody(body string)
	GetResponseBody() string

	EndTransfer()
}

type ITransferConnection interface {
	NextData() ITransferData

	CloseConnection()
}

type ITransferer interface {
	NextConnection() ITransferConnection
}