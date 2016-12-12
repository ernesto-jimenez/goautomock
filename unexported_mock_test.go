/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package main

// unexportedMock mock
type unexportedMock struct {
	ReadFunc func([]byte) (int, error)
}

// Read mocked method
func (m *unexportedMock) Read(p0 []byte) (int, error) {
	if m.ReadFunc == nil {
		panic("unexpected call to mocked method Read")
	}
	return m.ReadFunc(p0)
}
