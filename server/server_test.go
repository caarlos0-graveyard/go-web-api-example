package server_test

import (
	"testing"
	"os"
	"io/ioutil"
	"net/http"
	"github.com/caarlos0/go-web-api-example/server"
)

func TestMain(m *testing.M) {
	s := server.New()
	go func() {
		s.ListenAndServe()
	}()

	result := m.Run()

	s.Shutdown(nil)
	os.Exit(result)
}

func TestStatus(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/status")

	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != 200 {
		t.Error("Expected 200, got", resp.StatusCode)
	} else {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		} else {
			body := string(bytes)
			if body != "OK" {
				t.Error("Expected OK, got", body)
			}
		}
	}
}
