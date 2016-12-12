package {{.Package}}

import (
	"fmt"
	mock "github.com/stretchr/testify/mock"
{{range $path, $name := .Imports}}
	{{$name}} "{{$path}}"{{end}}
)

// {{.Name}} mock
type {{.Name}} struct {
	mock.Mock
}

{{$gen := .}}
{{range .Methods}}
// {{.Name}} mocked method
func (m *{{$gen.Name}}) {{.Name}}({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}} {{$type}}{{end}}) ({{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}{{$type}}{{end}}) {
{{if .ReturnTypes}}
	ret := m.Called({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}}{{end}})
	{{range $index, $type := .ReturnTypes}}
	var r{{$index}} {{$type}}
	switch res := ret.Get({{$index}}).(type) {
	case nil:
	case {{$type}}:
		r{{$index}} = res
	default:
		panic(fmt.Sprintf("unexpected type: %v", res))
	}
	{{end}}
	return {{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}r{{$index}}{{end}}
{{else}}
	m.Called({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}}{{end}})
{{end}}
}
{{end}}
