/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package main

// unexportedMock mock
type unexportedMock struct {
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
	if m.ReadFunc == nil {
		panic("unexpected call to mocked method Read")
	}
	m.calls["Read"]++
	return m.ReadFunc(p0)
}

// ReadCalls returns the amount of calls to the mocked method Read
func (m *unexportedMock) ReadTotalCalls() int {
	return m.calls["Read"]
}
