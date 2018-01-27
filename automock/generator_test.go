package automock

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	f, err := ioutil.ReadFile("../templates/testify.go.tpl")
	if err != nil {
		log.Fatal(err)
	}
	DefaultTemplate = string(f)
}

func TestNewGenerator(t *testing.T) {
	_, err := NewGenerator("io", "Writer")
	assert.NoError(t, err)
}

func TestNewGeneratorErrors(t *testing.T) {
	_, err := NewGenerator("someNonsense", "Writer")
	assert.Error(t, err)

	_, err = NewGenerator("io", "SomeWriter")
	assert.Error(t, err)
}

func TestMethods(t *testing.T) {
	g, err := NewGenerator("io", "Writer")
	assert.NoError(t, err)
	assert.Len(t, g.Methods(), 1)
}

func TestImports(t *testing.T) {
	g, err := NewGenerator("io", "Writer")
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{}, g.Imports())

	g, err = NewGenerator("net/http", "CookieJar")
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{
		"net/http": "http",
		"net/url":  "url",
	}, g.Imports())
}

type unexported interface {
	io.Reader
	Testing(*Generator)
	Test(int, ...string)
}

func TestWritesProperly(t *testing.T) {
	tests := []struct {
		pkg   string
		iface string
	}{
		{"net/http", "CookieJar"},
		{"io", "Writer"},
		{"io", "ByteScanner"},
		{"github.com/ernesto-jimenez/goautomock/automock", "unexported"},
		{".", "unexported"},
	}
	for _, test := range tests {
		t.Run(path.Join(test.pkg, test.iface), func(t *testing.T) {
			var out bytes.Buffer
			g, err := NewGenerator(test.pkg, test.iface)
			require.NoError(t, err)
			err = g.Write(&out)
			t.Log(test)
			printWithLines(t, bytes.NewBuffer(out.Bytes()))
			require.NoError(t, err)
		})
	}
}

func printWithLines(t *testing.T, txt io.Reader) {
	line := 0
	scanner := bufio.NewScanner(txt)
	for scanner.Scan() {
		line++
		t.Logf("%-4d| %s\n", line, scanner.Text())
	}
}
