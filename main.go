package main

import (
	"Gone/hosting"
	basic_middleware "Gone/middlewares/basic"
	"Gone/transfering/tcptext"
	"net"
)

func main() {
	host := hosting.UnsynchronisedHost{}
	listener, _ := net.Listen("tcp", "localhost:8080")
	host.Setup(&tcptext.TcpTextTransferer{Listener: listener}, &basic_middleware.BasicMiddleware{})
	host.Start()
}
