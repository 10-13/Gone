package hosting

type UnsynchronisedHost struct {
	trs ITransferer
	hnd IMiddleware

	running bool
}

func (self UnsynchronisedHost*) Setup(trs ITransferer, hnd IMiddleware) bool {
	self.trs = trs;
	self.hnd = hnd;
}

func (self UnsynchronisedHost*) RunConnection(conn ITransferConnection) {
	for !conn.IsClosed() {
		var data := conn.NextData();

		self.hnd.Handle(data);

		data.EndTransfer();
	}
}

func (self UnsynchronisedHost*) RunServer() {
	for running {
		var conn ITransferConnection := self.trs.NextConnection();

		go self.RunConnection(conn)
	}
}

func (self UnsynchronisedHost*) Start() bool {
	if running {
		return false;
	}

	go self.RunServer();
	running = true;
	return true;
}

func (self UnsynchronisedHost*) Stop() bool {
	if !running {
		return false;
	}

	running = false;
	return true;
}
