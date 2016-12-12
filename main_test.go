package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//go:generate go build .
//go:generate ./goautomock -template=testify io.Writer
//go:generate ./goautomock -template=testify -pkg=io ByteScanner
//go:generate ./goautomock -template=testify net/http.CookieJar

func TestMockedIOWriter(t *testing.T) {
	m := &WriterMock{}
	expected := []byte("hello")
	m.On("Write", expected).Return(5, nil)
	n, err := m.Write(expected)
	assert.Equal(t, 5, n)
	assert.Equal(t, nil, err)
}

func TestMockedCookieJar(t *testing.T) {
	jar := &CookieJarMock{}
	cookie := http.Cookie{Name: "hello", Value: "World"}
	jar.On("Cookies", mock.AnythingOfType("*url.URL")).Return([]*http.Cookie{&cookie}).Once()
	c := http.Client{Jar: jar}
	c.Get("http://localhost")

	jar.On("Cookies", mock.AnythingOfType("*url.URL")).Return(nil).Once()
	c.Get("http://localhost")
}

func TestMockByteScanner(t *testing.T) {
	var s io.ByteScanner
	m := &ByteScannerMock{}
	s = m
	m.On("ReadByte").Return(byte('_'), nil)
	b, err := s.ReadByte()
	assert.Equal(t, byte('_'), b)
	assert.Equal(t, nil, err)
}

type unexported interface {
	io.Reader
}

//go:generate ./goautomock unexported

func TestUnexported(t *testing.T) {
	m := &unexportedMock{}
	var calls int
	m.ReadFunc = func(b []byte) (int, error) {
		calls++
		return 0, nil
	}
	m.Read([]byte{})
	assert.Equal(t, 1, calls)
}
