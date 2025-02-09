package transporter

import (
	"net"

	"github.com/Ehco1996/ehco/internal/lb"
	"github.com/Ehco1996/ehco/internal/logger"
)

type Mwss struct {
	raw *Raw
	mtp *mwssTransporter
}

func (s *Mwss) GetOrCreateBufferCh(uaddr *net.UDPAddr) *BufferCh {
	return s.raw.GetOrCreateBufferCh(uaddr)
}

func (s *Mwss) HandleUDPConn(uaddr *net.UDPAddr, local *net.UDPConn) {
	s.raw.HandleUDPConn(uaddr, local)
}

func (s *Mwss) HandleTCPConn(c *net.TCPConn, remote *lb.Node) error {
	defer c.Close()
	mwsc, err := s.mtp.Dial(remote.Address + "/mwss/")
	if err != nil {
		return err
	}
	defer mwsc.Close()
	logger.Infof("[mwss] HandleTCPConn from:%s to:%s", c.RemoteAddr(), remote.Address)
	return transport(c, mwsc, remote.Label)
}

func (s *Mwss) GetRemote() *lb.Node {
	return s.raw.GetRemote()
}

func (s *Mwss) LimitByIp(c *net.TCPConn) error {
	return s.raw.LimitByIp(c)
}
