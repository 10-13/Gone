package hosting

import "Gone/core"

type UnsynchronisedHost struct {
	trs core.ITransferer
	hnd core.IMiddleware

	running bool
}

func (self *UnsynchronisedHost) Setup(trs core.ITransferer, hnd core.IMiddleware) bool {
	if self.running {
		return false
	}

	self.trs = trs
	self.hnd = hnd
	return true
}

func (self *UnsynchronisedHost) RunConnection(conn core.ITransferConnection) {
	for !conn.IsClosed() {
		data := conn.NextData()

		self.hnd.Handle(data)

		data.EndTransfer()
	}
}

func (self *UnsynchronisedHost) RunServer() {
	for self.running {
		conn := self.trs.NextConnection()

		go self.RunConnection(conn)
	}
}

func (self *UnsynchronisedHost) Start() bool {
	if self.running {
		return false
	}

	go self.RunServer()
	self.running = true
	return true
}

func (self UnsynchronisedHost) Stop() bool {
	if !self.running {
		return false
	}

	self.running = false
	return true
}
