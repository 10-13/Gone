package core

type header_name = string
type header_value = string

struct ITransferData interface {
	GetRequestHeaders() []string

	SetResponseHeader(name header_name, value header_value) header_value
	GetResponseHeader(name header_name) bool, header_value
	RemoveResponseHeader(name header_name) header_value

	GetRequsetBody() string

	SetResponseBody(body string)
	GetResponseBody() string

	EndTransfer()
}

struct ITransferConnection interface {
	NextData() ITransferData

	CloseConnection()
}

struct ITransferer interface {
	NextConnection() ITransferConnection
}