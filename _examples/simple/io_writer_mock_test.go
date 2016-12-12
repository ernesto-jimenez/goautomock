/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package simple

// WriterMock mock
type WriterMock struct {
	WriteFunc func([]byte) (int, error)
}

// Write mocked method
func (m *WriterMock) Write(p0 []byte) (int, error) {
	if m.WriteFunc == nil {
		panic("unexpected call to mocked method Write")
	}
	return m.WriteFunc(p0)
}
