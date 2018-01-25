/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package main

import (
	"sync"

	automock "github.com/ernesto-jimenez/goautomock/automock"
)

// unexported2 mock
type unexported2 struct {
	mut   sync.RWMutex
	calls map[string]int

	ReadFunc func([]byte) (int, error)

	TestingFunc func(*automock.Generator)
}

func NewUnexported2() *unexported2 {
	return &unexported2{
		calls: make(map[string]int),
	}
}

// Read mocked method
func (m *unexported2) Read(p0 []byte) (int, error) {
	m.mut.Lock()
	defer m.mut.Unlock()
	if m.ReadFunc == nil {
		panic("unexpected call to mocked method Read")
	}
	m.calls["Read"]++
	return m.ReadFunc(p0)
}

// ReadTotalCalls returns the amount of calls to the mocked method Read
func (m *unexported2) ReadTotalCalls() int {
	m.mut.RLock()
	defer m.mut.RUnlock()
	return m.calls["Read"]
}

// Testing mocked method
func (m *unexported2) Testing(p0 *automock.Generator) {
	m.mut.Lock()
	defer m.mut.Unlock()
	if m.TestingFunc == nil {
		panic("unexpected call to mocked method Testing")
	}
	m.calls["Testing"]++
	m.TestingFunc(p0)
}

// TestingTotalCalls returns the amount of calls to the mocked method Testing
func (m *unexported2) TestingTotalCalls() int {
	m.mut.RLock()
	defer m.mut.RUnlock()
	return m.calls["Testing"]
}
