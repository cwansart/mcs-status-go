package status

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	t.Run("test if is online", func(t *testing.T) {
		t.Parallel()

		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{"players": {"count": 1}}`))
		}))
		defer s.Close()

		r := Get(s.URL)

		if r.IsOnline == false {
			t.Errorf("IsOnline should be true")
		}
		if r.PlayerCount != 1 {
			t.Errorf("PlayerCount should be 1")
		}
	})

	t.Run("test if is offline", func(t *testing.T) {
		t.Parallel()

		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		}))
		url := s.URL
		s.Close()

		r := Get(url)

		if r.IsOnline == true {
			t.Errorf("IsOnline should be false")
		}
		if r.PlayerCount > 0 {
			t.Errorf("PlayerCount should be 0")
		}
	})

	t.Run("test invalid json response from server", func(t *testing.T) {
		t.Parallel()

		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{}`))
		}))
		defer s.Close()

		r := Get(s.URL)

		if r.IsOnline == false {
			t.Errorf("IsOnline should be true")
		}
		if r.PlayerCount > 0 {
			t.Errorf("PlayerCount should be 0")
		}
	})

	t.Run("test json unmarshalling error", func(t *testing.T) {
		t.Parallel()

		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`[ini]\nconf=test`))
		}))
		defer s.Close()

		r := Get(s.URL)

		if r.IsOnline == true {
			t.Errorf("IsOnline should be false")
		}
		if r.PlayerCount > 0 {
			t.Errorf("PlayerCount should be 0")
		}
	})
}
