# goautomock

Automatically generate mocks from interfaces using `go generate`

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
	m := &serverMock{}
	var calls int
	m.ServeFunc = func(in string) ([]byte, error) {
		calls++
		return nil, errors.New("error")
	}
	_, err := request(m, "/something")
	assert.Error(t, err)
	assert.Equal(t, 1, calls)
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
	ServeFunc func(string) ([]byte, error)
}

// Serve mocked method
func (m *serverMock) Serve(p0 string) ([]byte, error) {
	if m.ServeFunc == nil {
		panic("unexpected call to mocked method Serve")
	}
	return m.ServeFunc(p0)
}
```

Mocking an interface from the standard library:

[embedmd]:# (_examples/simple/stdlib_test.go /.*goautomock.*/ $)
```go
//go:generate goautomock io.Writer

// Dummy test using the generated mock
func TestWriter(t *testing.T) {
	m := &WriterMock{}
	expected := []byte("hello world")

	m.WriteFunc = func(b []byte) (int, error) {
		assert.Equal(t, expected, b)
		return len(b), nil
	}

	n, err := m.Write(expected)
	assert.NoError(t, err)
	assert.Equal(t, 11, n)
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
	m := &serverMock{}
	var calls int
	m.ServeFunc = func(in string) ([]byte, error) {
		calls++
		return nil, errors.New("error")
	}
	_, err := request(m, "/something")
	assert.Error(t, err)
	assert.Equal(t, 1, calls)
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
