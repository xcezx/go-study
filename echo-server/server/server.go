package server

import (
	"log"
	"net/http"
	"os"
	"fmt"
)

type Server struct {
	accessLogger *log.Logger
}

func New() *Server {
	logger := log.New(os.Stdout, "*** ", log.LUTC | log.LstdFlags)
	return &Server{
		accessLogger: logger,
	}
}

func (srv *Server) EchoHandler() http.Handler {
	handler := http.HandlerFunc(srv.echoHandler)
	return srv.loggingMiddleware(handler)
}

func (srv *Server) echoHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(req.URL.Query().Get("msg")))
}

func (srv *Server) PingHandler() http.Handler {
	handler := http.HandlerFunc(srv.pingHandler)
	return srv.loggingMiddleware(handler)
}

func (srv *Server) pingHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Pong")
}

func (srv *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		srv.accessLogger.Printf("method:%s\turl:%s", req.Method, req.RequestURI)
		next.ServeHTTP(w, req)
	})
}
