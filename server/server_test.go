package server

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	s := New()
	go func() {
		s.ListenAndServe()
	}()

	result := m.Run()

	s.Shutdown(nil)
	os.Exit(result)
}

func TestStatus(t *testing.T) {
	resp, err := http.Get("http://localhost:3000/status")

	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, 200, "status could should be 200")

	bytes, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)
	assert.Equal(t, string(bytes), "OK", "body should say OK")
}
