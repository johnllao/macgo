package client

import (
	"bufio"
	"net"

	"github.com/johnllao/macgo/pkg/sock/utils"
)

type ClientError struct {
	errmsg string
	InnerErr error
}

func (e *ClientError) Error() string {
	return e.errmsg
}

type Client struct {
	conn net.Conn
}

func Connect(addr string) (*Client, error) {
	var err error
	var conn net.Conn
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		return nil, &ClientError{ errmsg: "Connect() Unable to connect to " + addr, InnerErr: err}
	}
	var cli = &Client{
		conn: conn,
	}
	return cli, nil
}

func (c *Client) Send(data []byte, receive bool) ([]byte, error) {

	var err error

	var p []byte
	p , err = utils.Payload(data)
	if err != nil {
		return nil, &ClientError{ errmsg: "Send() Unable to read payload", InnerErr: err }
	}

	_, err = c.conn.Write(p)
	if err != nil {
		return nil, &ClientError{ errmsg: "Send() Unable to write payload to socket", InnerErr: err }
	}

	if !receive {
		return nil, nil
	}

	var r = bufio.NewReader(c.conn)
	p, err = r.ReadBytes(utils.FileSeparator)
	if err != nil {
		return nil, &ClientError{ errmsg: "Send() Unable to read response from socket", InnerErr: err }
	}

	var d []byte
	var s int
	s, d, err = utils.PayloadData(p)
	if err != nil {
		return nil, &ClientError{ errmsg: "Send() Unable to read response payload from socket", InnerErr: err }
	}

	if s != len(d) {
		return nil, &ClientError{ errmsg: "Send() Invalid response payload", InnerErr: err }
	}

	return d, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
