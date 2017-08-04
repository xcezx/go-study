package server

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	accessLogger *log.Logger
}

func New(accessLogger *log.Logger) *server {
	return &server{
		accessLogger: accessLogger,
	}
}

func (srv *server) EchoHandler() http.Handler {
	handler := http.HandlerFunc(srv.echoHandler)
	return srv.loggingMiddleware(handler)
}

func (srv *server) echoHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(req.URL.Query().Get("msg")))
}

func (srv *server) PingHandler() http.Handler {
	handler := http.HandlerFunc(srv.pingHandler)
	return srv.loggingMiddleware(handler)
}

func (srv *server) pingHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Pong")
}

func (srv *server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		srv.accessLogger.Printf("method:%s\turl:%s", req.Method, req.RequestURI)
		next.ServeHTTP(w, req)
	})
}
