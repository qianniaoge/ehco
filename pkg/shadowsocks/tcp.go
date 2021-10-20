package shadowsocks

import (
	"net"
)

// StartTCPServer Listen on addr for incoming connections.
func StartTCPServer(addr string, shadow func(net.Conn) net.Conn) error {
	l, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		return err
	}
	for {
		c, err := l.Accept()
		if err != nil {
			println("accept error", err.Error())
			continue
		}
		go func() {
			defer c.Close()
			c = shadow(c)
		}()
	}
}
