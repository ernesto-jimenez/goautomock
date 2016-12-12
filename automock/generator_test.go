package automock

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
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
		var out bytes.Buffer
		g, err := NewGenerator(test.pkg, test.iface)
		if err != nil {
			t.Logf("%+v", test)
			t.Error(err)
			continue
		}
		err = g.Write(&out)
		if !assert.NoError(t, err) {
			fmt.Println(test)
			fmt.Println(err)
			printWithLines(bytes.NewBuffer(out.Bytes()))
		}
	}
}

func printWithLines(txt io.Reader) {
	line := 0
	scanner := bufio.NewScanner(txt)
	for scanner.Scan() {
		line++
		fmt.Printf("%-4d| %s\n", line, scanner.Text())
	}
}
