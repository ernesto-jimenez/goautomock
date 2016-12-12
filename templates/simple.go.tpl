package {{.Package}}

import (
	"fmt"
{{range $path, $name := .Imports}}
	{{$name}} "{{$path}}"{{end}}
)

// {{.Name}} mock
type {{.Name}} struct {
	calls map[string]int
	{{range .Methods}}
	{{.Name}}Func func({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}{{$type}}{{end}}) ({{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}{{$type}}{{end}})
	{{end}}
}

func New{{.Name | Capitalize}}() *{{.Name}} {
	return &{{.Name}}{
		calls: make(map[string]int),
	}
}

{{$gen := .}}
{{range .Methods}}
// {{.Name}} mocked method
func (m *{{$gen.Name}}) {{.Name}}({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}} {{$type}}{{end}}) ({{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}{{$type}}{{end}}) {
	if m.{{.Name}}Func == nil {
		panic("unexpected call to mocked method {{.Name}}")
	}
	m.calls["{{.Name}}"]++
	{{if .ReturnTypes}}return {{end}} m.{{.Name}}Func({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}}{{end}})
}

// {{.Name}}Calls returns the amount of calls to the mocked method {{.Name}}
func (m *{{$gen.Name}}) {{.Name}}TotalCalls() int {
	return m.calls["{{.Name}}"]
}
{{end}}
