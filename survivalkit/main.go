package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/johnllao/macgo/survivalkit/server"
)

var (
	dbpath   string
	logfile  *os.File
	logpath  string
	port 	 int
	rootpath string
)

func init() {
	flag.StringVar(&dbpath, "dbpath", "", "path of the boltdb data")
	flag.StringVar(&logpath, "logpath", "stderr", "path of the log file")
	flag.StringVar(&rootpath, "rootpath", "", "path of the static files")
	flag.IntVar(&port, "port", 8080, "service port number")
	flag.Parse()

	if logpath != "stderr" {
		var err error
		logfile, err = os.Open(logpath)
		if err == nil {
			log.SetOutput(logfile)
		}
	}
}

func main() {
	var err error
	var s *server.Service
	s, err = server.Start(port, rootpath, dbpath)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	var c = make(chan os.Signal)
	signal.Notify(c, os.Kill, syscall.SIGTERM)
	go func() {
		<-c

		s.Stop()

		if logfile != nil {
			_ = logfile.Close()
		}
		os.Exit(0)
	}()
}
