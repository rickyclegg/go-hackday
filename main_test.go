package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDevs(t *testing.T) {
	t.Run("get list of devs", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/devs", nil)

		if err != nil {
			t.Fatal(err)
		}

		want := "data"
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetDevs([]byte(want)))

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		got := rr.Body.String()

		if got != want {
			t.Errorf("handler returned unexpected body: got %v want %v", got, want)
		}
	})
}
