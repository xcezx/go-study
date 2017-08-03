package server

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupServer() *Server {
	var buf bytes.Buffer
	logger := log.New(&buf, "*** ", log.LUTC|log.LstdFlags)
	return &Server{
		accessLogger: logger,
	}
}

func Test__pingHandler(t *testing.T) {
	srv := setupServer()
	s := httptest.NewServer(http.HandlerFunc(srv.pingHandler))
	defer s.Close()

	res, err := http.Get(s.URL)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d to eq %d", res.StatusCode, http.StatusOK)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if !strings.Contains(string(body), "Pong") {
		t.Errorf("expected %s to eq Pong", body)
	}
}

func Test__echoHandler(t *testing.T) {
	srv := setupServer()
	s := httptest.NewServer(http.HandlerFunc(srv.echoHandler))
	defer s.Close()

	res, err := http.Get(s.URL + "?msg=Hello")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d to eq %d", res.StatusCode, http.StatusOK)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if !strings.Contains(string(body), "Hello") {
		t.Errorf("expected %s to eq %s", body, "Hello")
	}
}
