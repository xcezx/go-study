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

func TestPingHandler(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "*** ", log.LUTC|log.LstdFlags)
	srv := &Server{
		accessLogger: logger,
	}
	handler := srv.PingHandler()

	s := httptest.NewServer(handler)
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

func TestEchoHandler(t *testing.T) {
	t.Skip()
}
