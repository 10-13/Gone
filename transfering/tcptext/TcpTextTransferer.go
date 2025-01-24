package tcptext

import (
	"Gone/core"
	. "Gone/custom_types"
)

type TextTransferData struct {
	request_headers  Headers
	response_headers Headers
	request_body     string
	response_body    string
}

func (self *TextTransferData) GetRequestHeaders() Headers {
	return self.request_headers
}

func (self *TextTransferData) SetResponseHeader(name HeaderName, value HeaderValue) HeaderValue {
	res, exists := self.response_headers[name]
	self.response_headers[name] = value
	if exists {
		return res
	}
	return HeaderValue("")
}

func (self *TextTransferData) GetResponseHeader(name HeaderName) (HeaderValue, error) {
	res, exists := self.response_headers[name]
	if !exists {
		return res, &NoSuchHeader{Code: 0, Message: string(name)}
	}
	return res, nil
}

func (self *TextTransferData) RemoveResponseHeader(name HeaderName) HeaderValue {
	res, exists := self.response_headers[name]
	delete(self.response_headers, name)
	if exists {
		return res
	}
	return HeaderValue("")
}

func (self *TextTransferData) GetRequsetBody() string {
	return self.request_body
}

func (self *TextTransferData) SetResponseBody(body string) {
	self.response_body = body
}

func (self *TextTransferData) GetResponseBody() string {
	return self.response_body
}

func (self *TextTransferData) EndTransfer() {
	panic("Not implemented.")
}

type TextTransferConnection struct {
	// TODO: Implement
}

func (self *TextTransferConnection) NextData() core.ITransferData {
	panic("Not implemented.")
}
func (self *TextTransferConnection) IsClosed() bool {
	panic("Not implemented.")
}

func (self *TextTransferConnection) CloseConnection() {
	panic("Not implemented.")
}

type TcpTextTransferer struct {
	// TODO: Implement
}

func (self *TcpTextTransferer) NextConnection() core.ITransferConnection {
	panic("Not implemented.")
}
