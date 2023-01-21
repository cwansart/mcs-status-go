package status_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"de.cwansart.mcss/status"
)

func TestIsOnline(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"players": {"count": 1}}`))
	}))
	defer s.Close()

	r := status.Get(s.URL)

	if r.IsOnline == false {
		t.Errorf("IsOnline should be true")
	}
	if r.PlayerCount != 1 {
		t.Errorf("PlayerCount should be 1")
	}
}

func TestIsOffline(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	url := s.URL
	s.Close()

	r := status.Get(url)

	if r.IsOnline == true {
		t.Errorf("IsOnline should be false")
	}
	if r.PlayerCount > 0 {
		t.Errorf("PlayerCount should be 0")
	}
}

func TestInvalidJsonServerResponse(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{}`))
	}))
	defer s.Close()

	r := status.Get(s.URL)

	if r.IsOnline == false {
		t.Errorf("IsOnline should be true")
	}
	if r.PlayerCount > 0 {
		t.Errorf("PlayerCount should be 0")
	}
}
