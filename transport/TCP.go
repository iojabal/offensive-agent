package transport

import "net"

type TCPTransport struct {
	conn net.Conn
	addr string
}

func NewTCPTransport(address string) *TCPTransport {
	return &TCPTransport{
		addr: address,
	}
}
func (t *TCPTransport) Connect() error {
	conn, err := net.Dial("tcp", t.addr)
	if err != nil {
		return err
	}
	t.conn = conn
	return nil
}
func (t *TCPTransport) Send(data []byte) (int, error) {
	return t.conn.Write(data)
}
func (t *TCPTransport) Read() ([]byte, error) {
	buf := make([]byte, 4096)
	n, err := t.conn.Read(buf)
	return buf[:n], err
}
func (t *TCPTransport) Close() error {
	return t.conn.Close()
}
