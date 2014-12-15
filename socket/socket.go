package socket

import (
	"net"
	"time"
)

type TSocket struct {
	conn    net.Conn
	addr    net.Addr
	timeout time.Duration
}

func NewTSocket(hostPort string) (*TSocket, error) {
	return NewTSocketTimeout(hostPort, 0)
}

func NewTSocketTimeout(hostPort string, timeout time.Duration) (*TSocket, error) {
	addr, err := net.ResolveIPAddr("tcp", hostPort)
	if err != nil {
		return nil, err
	}
	return NewTSocketFromAddrTimeout(addr, timeout)
}

func NewTSocketFromAddrTimeout(addr net.Addr, timeout time.Duration) *TSocket {
	return &TSocket{addr: addr, timeout: timeout}
}

func NewTSocketFromConnTimeout(conn net.Conn, timeout time.Duration) *TSocket {
	return &TSocket{conn: conn, addr: conn.RemoteAddr(), timeout: timeout}
}

func (t *TSocket) SetTimeout(timeout time.Duration) {
	t.timeout = timeout
}

func (t *TSocket) Open() {

}

func (t *TSocket) IsOpen() {
	if t.conn == nil {
	}
}
