package tcptext

import (
	"Gone/core"
	. "Gone/custom_types"
	"io"
	"net"
	"strings"
)

type TextTransferData struct {
	http_method      string
	url              string
	request_headers  Headers
	response_headers Headers
	request_body     string
	response_body    string
	connection       *TextTransferConnection
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
	res := "HTTP/1.1 200 OK"
	for k, v := range self.response_headers {
		res += "\n"
		res += string(k) + ": " + string(v)
	}
	res += "\n\n"
	res += self.response_body
	self.connection.TcpConnection.Write([]byte(res))
}

type TextTransferConnection struct {
	TcpConnection net.Conn
	closed        bool
}

func (self *TextTransferConnection) NextData() core.ITransferData {
	req := ""
	BUF_SIZE := 1024
	buf := make([]byte, BUF_SIZE)
	for {
		n, err := self.TcpConnection.Read(buf)
		if err != nil && err != io.EOF {
			panic("something went wrong in next data method")
		}
		req += string(buf[:n])
		if n < BUF_SIZE {
			break
		}
	}
	lines := strings.Split(req, "\n")

	http_method := strings.Split(lines[0], " ")[0]
	url := strings.Split(lines[0], " ")[1]
	headers := Headers{}

	i := 1
	for ; i < len(lines) && lines[i] != ""; i++ {
		split_idx := strings.Index(lines[i], ": ")
		if split_idx == -1 {
			break
		}
		headers[HeaderName(lines[i][:split_idx])] = HeaderValue(lines[i][split_idx+2:])
	}

	body := ""
	i++
	for ; i < len(lines); i++ {
		body += lines[i]
		if i+1 < len(lines) {
			body += "\n"
		}
	}

	return &TextTransferData{
		http_method:      http_method,
		url:              url,
		request_headers:  headers,
		response_headers: Headers{},
		request_body:     body,
		response_body:    "",
		connection:       self,
	}
}

func (self *TextTransferConnection) IsClosed() bool {
	return self.closed
}

func (self *TextTransferConnection) CloseConnection() {
	self.closed = true
	self.TcpConnection.Close()
}

type TcpTextTransferer struct {
	Listener net.Listener
}

func (self *TcpTextTransferer) NextConnection() core.ITransferConnection {
	connection, _ := self.Listener.Accept()
	return &TextTransferConnection{
		TcpConnection: connection,
		closed:        false,
	}
}
