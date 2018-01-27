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

	TestFunc func(int, ...string)

	TestingFunc func(*automock.Generator)
}

func NewUnexported2() *unexported2 {
	return &unexported2{
		calls: make(map[string]int),
	}
}

// Read mocked method
func (m *unexported2) Read(p0 []byte) (int, error) {
	defer func() {
		m.mut.Lock()
		m.calls["Read"]++
		m.mut.Unlock()
	}()
	if m.ReadFunc == nil {
		panic("unexpected call to mocked method Read")
	}
	return m.ReadFunc(p0)
}

// ReadTotalCalls returns the amount of calls to the mocked method Read
func (m *unexported2) ReadTotalCalls() int {
	m.mut.RLock()
	defer m.mut.RUnlock()
	return m.calls["Read"]
}

// Test mocked method
func (m *unexported2) Test(p0 int, p1 ...string) {
	defer func() {
		m.mut.Lock()
		m.calls["Test"]++
		m.mut.Unlock()
	}()
	if m.TestFunc == nil {
		panic("unexpected call to mocked method Test")
	}
	m.TestFunc(p0, p1...)
}

// TestTotalCalls returns the amount of calls to the mocked method Test
func (m *unexported2) TestTotalCalls() int {
	m.mut.RLock()
	defer m.mut.RUnlock()
	return m.calls["Test"]
}

// Testing mocked method
func (m *unexported2) Testing(p0 *automock.Generator) {
	defer func() {
		m.mut.Lock()
		m.calls["Testing"]++
		m.mut.Unlock()
	}()
	if m.TestingFunc == nil {
		panic("unexpected call to mocked method Testing")
	}
	m.TestingFunc(p0)
}

// TestingTotalCalls returns the amount of calls to the mocked method Testing
func (m *unexported2) TestingTotalCalls() int {
	m.mut.RLock()
	defer m.mut.RUnlock()
	return m.calls["Testing"]
}
