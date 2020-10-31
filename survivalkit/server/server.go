package server

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
)

type Service struct {
	APIHandlers http.Handler
	DB          *bolt.DB
	DBPath      string
	FileServer  http.Handler
	HTTPServer  *http.Server
	Port        int
	RootPath    string
}

func (s *Service) Start() error {

	var a = http.NewServeMux()
	a.HandleFunc("/api/about", s.abouthandler)
	s.APIHandlers = a

	s.HTTPServer = &http.Server{
		Addr:    ":" + strconv.Itoa(s.Port),
		Handler: http.HandlerFunc(s.roothandler),
	}

	log.Printf("starting service. port: %d", s.Port)
	return s.HTTPServer.ListenAndServe()
}

func (s *Service) Stop() {
	if s.DB != nil {
		_ = s.DB.Close()
	}
	if s.HTTPServer != nil {
		_ = s.HTTPServer.Close()
	}
}

func (s *Service) roothandler(w http.ResponseWriter, r *http.Request) {
	var start = time.Now()

	var ssid = Session(w, r)

	if strings.HasPrefix(r.URL.Path, "/api") {
		s.APIHandlers.ServeHTTP(w, r)
		log.Printf("[%s] %s %s elapsed: %8.5f ms", ssid, r.Method, r.URL.Path, float64(time.Since(start)) / float64(time.Millisecond))
		return
	}

	s.FileServer.ServeHTTP(w, r)
	log.Printf("[%s] %s %s elapsed: %8.5f ms", ssid, r.Method, r.URL.Path, float64(time.Since(start)) / float64(time.Millisecond))
}

func (s *Service) abouthandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("About"))
}

func Start(port int, rootpath, dbpath string) (*Service, error) {

	var err error

	var s = &Service{
		DBPath:   dbpath,
		RootPath: rootpath,
		Port:     port,
	}

	log.Printf("initializing db. dbpath: %s", s.DBPath)
	s.DB, err = bolt.Open(s.DBPath, 0640, nil)
	if err != nil {
		return nil, err
	}

	if s.RootPath != "" {
		log.Printf("initializing root path. rootpath: %s", s.RootPath)
		var root = http.Dir(s.RootPath)
		s.FileServer = http.FileServer(root)
	}

	err = s.Start()
	if err != nil {
		return nil, err
	}

	return s, nil
}
