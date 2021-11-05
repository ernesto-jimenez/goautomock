# doubler

Automatically generate doubles from interfaces using `go generate`.

# Installation

```shell
go get github.com/takooo/doubler
```

# Usage

Creating an interface in your code to mock a dependency:

[embedmd]:# (_examples/simple/example_test.go /.*goautomock.*/ $)
```go
//go:generate doubler

type server interface {
	Serve(input string) ([]byte, error)
}

func request(s server, path string) ([]byte, error) {
	return s.Serve(path)
}

// Dummy test
func TestRequestReturnsServerError(t *testing.T) {
	d := NewServerDouble()
	d.SetReturn(d.Serve, []byte{}, nil)
	_, err := request(m, "/something")
	assert.Error(t, err)
	assert.Equal(t, 1, m.ServeTotalCalls())
}
```

Here you can see what the generated mock looks like:

[embedmd]:# (_examples/simple/server_mock_test.go)
```go
/*
* CODE GENERATED AUTOMATICALLY WITH github.com/takaooo/doubler
* THIS FILE MUST NEVER BE EDITED MANUALLY
 */

package simple

import (
	. "github.com/sagansystems/doppelgaenger"
)

// serverMock mock
type ServerDouble struct {
	Double
}

func NewServerDouble() *ServerDouble {
	return &ServerDouble{}
}

func (d *ServerDouble) Serve(input string) ([]byte, error) {
	rets := d.CallMethod(d.Serve, input)
	inputReturn, _ := rets[0].(string)
	return inputReturn, ErrorReturn(rets[1])
}
```
