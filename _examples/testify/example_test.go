package testify

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:generate goautomock -template=testify server

type server interface {
	Serve(string) ([]byte, error)
}

func request(s server, path string) ([]byte, error) {
	return s.Serve(path)
}

// Dummy test
func TestRequestReturnsServerError(t *testing.T) {
	m := &serverMock{}
	m.On("Serve", "/something").Return(nil, errors.New("error"))

	_, err := request(m, "/something")
	assert.Error(t, err)
	m.AssertNumberOfCalls(t, "Serve", 1)
}
