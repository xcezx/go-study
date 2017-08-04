package main

import (
	"log"
	"net/http"
	"os"

	"github.com/xcezx/go-study/echo-server/server"
)

func main() {
	logger := log.New(os.Stdout, "*** ", log.LUTC|log.LstdFlags)
	srv := server.New(logger)

	http.Handle("/echo", srv.EchoHandler())
	http.Handle("/ping", srv.PingHandler())

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("http.ListenAndServer faild: %v", err)
	}
}
