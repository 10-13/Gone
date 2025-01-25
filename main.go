package main

import (
	"Gone/hosting"
	add_content_length_middleware "Gone/middlewares/add_content_length"
	"Gone/transfering/tcptext"
	"fmt"
	"net"
)

func main() {
	host := hosting.UnsynchronisedHost{}
	listener, _ := net.Listen("tcp", "localhost:8080")
	host.Setup(
		&tcptext.TcpTextTransferer{Listener: listener},
		&add_content_length_middleware.AddContentLengthMiddleware{},
	)
	host.Start()
	var input string
	fmt.Scanln(&input)
}
