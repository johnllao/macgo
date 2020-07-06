package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
)

const (
	BuffLen = 1024
	FileSeparator = 0x1C
)

type ResponseWriter struct {
	conn net.Conn
}

func (w *ResponseWriter) Write(p []byte) (int, error) {
	var d = append(p, FileSeparator)
	return w.conn.Write(d)
}

type Server struct {
	Port       int
	MsgHandler func(io.Writer, []byte)

	listener   net.Listener
}

func (s *Server) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", "localhost:" + strconv.Itoa(s.Port))
	if err != nil {
		return err
	}
	for {
		var conn net.Conn
		conn, err = s.listener.Accept()
		if err != nil {
			log.Print("WARN: ", err)
		}

		go s.process(conn)
	}
}

func (s *Server) process (conn net.Conn) {
	var err error
	var r = bufio.NewReader(conn)
	for {
		var data []byte
		data, err = r.ReadBytes(0x1C)
		if err == io.EOF {
			break
		}
		if err != nil {
			conn.Close()
			break
		}
		var datalen = len(data)
		data = data[:datalen - 1]
		if s.MsgHandler != nil {
			var w = &ResponseWriter{
				conn: conn,
			}
			go s.MsgHandler(w, data)
		}
	}
}
