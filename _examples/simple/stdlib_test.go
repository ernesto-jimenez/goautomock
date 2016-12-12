package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
