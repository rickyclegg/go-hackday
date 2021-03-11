package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetDevs(t *testing.T) {
	t.Run("get list of devs", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/devs", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetDevs)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		want := []string{
			"Rick",
			"Lisa",
			"Osh",
			"Drilon",
			"Tom",
		}
		var got []string

		json.NewDecoder(rr.Body).Decode(&got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("handler returned unexpected body: got %v want %v", got, want)
		}
	})
}
