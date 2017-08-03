package main

import (
	"log"
	"net/http"

	"github.com/xcezx/go-study/echo-server/server"
)

func main() {
	srv := server.New()

	http.Handle("/echo", srv.EchoHandler())
	http.Handle("/ping", srv.PingHandler())

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("http.ListenAndServer faild: %v", err)
	}
}
