package tcptext

import (
	"Gone/core"
	. "Gone/custom_types"
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
}

func (self *TextTransferConnection) NextData() core.ITransferData {
	req := ""
	buf := make([]byte, 1024)
	for {
		n, err := self.TcpConnection.Read(buf)
		if err != nil {
			panic("something went wrong in next data method")
		}
		if n == 0 {
			break
		}
		req += string(buf[:n])
	}
	lines := strings.Split(req, "\n")

	http_method := strings.Split(lines[0], " ")[0]
	url := strings.Split(lines[0], " ")[1]
	headers := Headers{}

	i := 1
	for ; i < len(lines) && lines[i] != ""; i++ {
		split_idx := strings.Index(lines[1], ": ")
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
