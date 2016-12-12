/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package simple

// serverMock mock
type serverMock struct {
	calls map[string]int

	ServeFunc func(string) ([]byte, error)
}

func NewServerMock() *serverMock {
	return &serverMock{
		calls: make(map[string]int),
	}
}

// Serve mocked method
func (m *serverMock) Serve(p0 string) ([]byte, error) {
	if m.ServeFunc == nil {
		panic("unexpected call to mocked method Serve")
	}
	m.calls["Serve"]++
	return m.ServeFunc(p0)
}

// ServeCalls returns the amount of calls to the mocked method Serve
func (m *serverMock) ServeTotalCalls() int {
	return m.calls["Serve"]
}
