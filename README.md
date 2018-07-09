# goautomock

Automatically generate mocks from interfaces using `go generate`.

# Installation

```shell
go get github.com/ernesto-jimenez/goautomock
```

# Usage

Creating an interface in your code to mock a dependency:

[embedmd]:# (_examples/simple/example_test.go /.*goautomock.*/ $)
```go
//go:generate goautomock server

type server interface {
	Serve(string) ([]byte, error)
}

func request(s server, path string) ([]byte, error) {
	return s.Serve(path)
}

// Dummy test
func TestRequestReturnsServerError(t *testing.T) {
	m := NewServerMock()
	m.ServeFunc = func(in string) ([]byte, error) {
		return nil, errors.New("error")
	}
	_, err := request(m, "/something")
	assert.Error(t, err)
	assert.Equal(t, 1, m.ServeTotalCalls())
}
```

Here you can see what the generated mock looks like:

[embedmd]:# (_examples/simple/server_mock_test.go)
```go
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

// NewServerMock creates serverMock
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
```

Mocking an interface from the standard library:

[embedmd]:# (_examples/simple/stdlib_test.go /.*goautomock.*/ $)
```go
//go:generate goautomock io.Writer

// Dummy test using the generated mock
func TestWriter(t *testing.T) {
	m := NewWriterMock()
	expected := []byte("hello world")

	m.WriteFunc = func(b []byte) (int, error) {
		assert.Equal(t, expected, b)
		return len(b), nil
	}

	n, err := m.Write(expected)
	assert.NoError(t, err)
	assert.Equal(t, 11, n)
	assert.Equal(t, 1, m.WriteTotalCalls())
}
```

# Picking a different mock style

`goautomock` ships with different mock templates. You can see which ones with:

```shell
$ goautomock -list-templates
simple
testify
```

You can pick the template you want to use with `-template`

[embedmd]:# (_examples/testify/example_test.go /.*goautomock.*/ $)
```go
//go:generate goautomock -template=testify server

type server interface {
	Serve(string) ([]byte, error)
}

func request(s server, path string) ([]byte, error) {
	return s.Serve(path)
}

// Dummy test
func TestRequestReturnsServerError(t *testing.T) {
	m := &serverMock{}
	m.On("Serve", "/something").Return(nil, errors.New("error"))

	_, err := request(m, "/something")
	assert.Error(t, err)
	m.AssertNumberOfCalls(t, "Serve", 1)
}
```

## Defining your own mock templates

You can use your own custom template by passing a file name to `-template`. Example:

```shell
# Create a new template based on goautocomplete's testify template
$ goautomock -template=simple -print-template > mock.go.tpl

# Customise your template

# Use your template
$ goautomock -template=mock.go.tpl io.Writer
```
