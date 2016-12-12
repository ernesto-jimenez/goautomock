/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/goautomock
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package simple

// serverMock mock
type serverMock struct {
	ServeFunc func(string) ([]byte, error)
}

// Serve mocked method
func (m *serverMock) Serve(p0 string) ([]byte, error) {
	if m.ServeFunc == nil {
		panic("unexpected call to mocked method Serve")
	}
	return m.ServeFunc(p0)
}
