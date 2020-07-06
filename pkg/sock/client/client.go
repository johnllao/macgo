package client

import (
	"bufio"
	"net"
)

const (
	FileSeparator = 0x1C
)

type Client struct {
	conn net.Conn
}

func Connect(addr string) (*Client, error) {
	var err error
	var conn net.Conn
	conn, err = net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	var cli = &Client{
		conn: conn,
	}
	return cli, nil
}

func (c *Client) Send(data []byte) ([]byte, error) {

	var err error
	var m = append(data, FileSeparator)
	_, err = c.conn.Write(m)
	if err != nil {
		return nil, err
	}

	var r = bufio.NewReader(c.conn)
	var res []byte
	res, err = r.ReadBytes(FileSeparator)
	if err != nil {
		return nil, err
	}

	return res[:len(res) - 1], nil
}

func (c *Client) Close() {
	c.conn.Close()
}
