package tcptext

type header_name = string
type header_value = string

import (
    "bufio"
    "fmt"
    "net"
    "os"
)


type TextTransferData struct {
	// TODO: Implement
}

func (self *TextTransferData) GetRequestHeaders() []string {
	panic("Not implemented.");
}

func (self *TextTransferData) SetResponseHeader(name header_name, value header_value) header_value {
	panic("Not implemented.");
}
func (self *TextTransferData) GetResponseHeader(name header_name) bool, header_value {
	panic("Not implemented.");
}
func (self *TextTransferData) RemoveResponseHeader(name header_name) header_value {
	panic("Not implemented.");
}

func (self *TextTransferData) GetRequsetBody() string {
	panic("Not implemented.");
}

func (self *TextTransferData) SetResponseBody(body string) {
	panic("Not implemented.");
}
func (self *TextTransferData) GetResponseBody() string {
	panic("Not implemented.");
}

func (self *TextTransferData) EndTransfer() {
	panic("Not implemented.");
}


type TextTransferConnection struct {
	// TODO: Implement
}

func (self *TextTransferConnection) NextData() ITransferData {
	panic("Not implemented.");
}
func (self *TextTransferConnection) IsClosed() bool {
	panic("Not implemented.");
}

func (self *TextTransferConnection) CloseConnection() {
	panic("Not implemented.");
}

type TextTransferer interface {
	// TODO: Implement
}

func (self *TextTransferer) NextConnection() ITransferConnection {
	panic("Not implemented.");
}