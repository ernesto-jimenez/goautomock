/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package main

import (
	"sync"
)

// unexportedMock mock
type unexportedMock struct {
	mut   sync.RWMutex
	calls map[string]int

	ReadFunc func([]byte) (int, error)
}

func NewUnexportedMock() *unexportedMock {
	return &unexportedMock{
		calls: make(map[string]int),
	}
}

// Read mocked method
func (m *unexportedMock) Read(p0 []byte) (int, error) {
	m.mut.Lock()
	defer m.mut.Unlock()
	if m.ReadFunc == nil {
		panic("unexpected call to mocked method Read")
	}
	m.calls["Read"]++
	return m.ReadFunc(p0)
}

// ReadTotalCalls returns the amount of calls to the mocked method Read
func (m *unexportedMock) ReadTotalCalls() int {
	m.mut.RLock()
	defer m.mut.RUnlock()
	return m.calls["Read"]
}
