package simple

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:generate goautomock server

type server interface {
	Serve(string) ([]byte, error)
}

func request(s server, path string) ([]byte, error) {
	return s.Serve(path)
}

// Dummy test
func TestRequestReturnsServerError(t *testing.T) {
	m := &serverMock{}
	var calls int
	m.ServeFunc = func(in string) ([]byte, error) {
		calls++
		return nil, errors.New("error")
	}
	_, err := request(m, "/something")
	assert.Error(t, err)
	assert.Equal(t, 1, calls)
}
