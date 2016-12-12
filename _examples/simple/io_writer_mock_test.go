/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package simple

// WriterMock mock
type WriterMock struct {
	calls map[string]int

	WriteFunc func([]byte) (int, error)
}

func NewWriterMock() *WriterMock {
	return &WriterMock{
		calls: make(map[string]int),
	}
}

// Write mocked method
func (m *WriterMock) Write(p0 []byte) (int, error) {
	if m.WriteFunc == nil {
		panic("unexpected call to mocked method Write")
	}
	m.calls["Write"]++
	return m.WriteFunc(p0)
}

// WriteCalls returns the amount of calls to the mocked method Write
func (m *WriterMock) WriteTotalCalls() int {
	return m.calls["Write"]
}
