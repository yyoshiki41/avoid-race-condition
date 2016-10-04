func (srv *Server) Serve(l net.Listener) error {
	....

	for {
		rw, e := l.Accept()
		....

		c := srv.newConn(rw)
		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve()
	}
}
