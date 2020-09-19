package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
)

const (
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
	Ready      chan int
	Close      chan int

	listener   net.Listener

}

func NewServer(port int, msgh func(io.Writer, []byte)) *Server {
	return &Server{
		Port:       port,
		MsgHandler: msgh,
		Ready:      make(chan int),
		Close:      make(chan int),
	}
}

func (s *Server) Start() error {
	var err error
	s.listener, err = net.Listen("tcp", "localhost:" + strconv.Itoa(s.Port))
	if err != nil {
		return err
	}
	go func() {
		s.Ready <- 1
	}()

	go func() {
		for {
			var conn net.Conn
			conn, err = s.listener.Accept()
			if err != nil {
				log.Print("WARN: ", err)
			}

			go s.process(conn)
		}
		s.Close <- 1
	}()

	return nil
}

func (s *Server) process (conn net.Conn) {
	var err error
	var r = bufio.NewReader(conn)
	for {
		var data []byte
		data, err = r.ReadBytes(FileSeparator)
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
