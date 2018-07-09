package {{.Package}}

import (
	"fmt"
	"sync"
{{range $path, $name := .Imports}}
	{{$name}} "{{$path}}"{{end}}
)

// {{.Name}} mock
type {{.Name}} struct {
	mut sync.RWMutex
	calls map[string]int
	{{range .Methods}}
	{{.Name}}Func func({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}{{$type}}{{end}}) ({{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}{{$type}}{{end}})
	{{end}}
}

// New{{.Name | Capitalize}} creates {{.Name}}
func New{{.Name | Capitalize}}() *{{.Name}} {
	return &{{.Name}}{
		calls: make(map[string]int),
	}
}

{{$gen := .}}
{{range .Methods}}
// {{.Name}} mocked method
func (m *{{$gen.Name}}) {{.Name}}({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}} {{$type}}{{end}}) ({{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}{{$type}}{{end}}) {
	defer func() {
		m.mut.Lock()
		m.calls["{{.Name}}"]++
		m.mut.Unlock()
	}()
	if m.{{.Name}}Func == nil {
		panic("unexpected call to mocked method {{.Name}}")
	}
	{{if .ReturnTypes}}return {{end}} m.{{.Name}}Func({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}}{{end}}{{if .Variadic}}...{{end}})
}

// {{.Name}}TotalCalls returns the amount of calls to the mocked method {{.Name}}
func (m *{{$gen.Name}}) {{.Name}}TotalCalls() int {
	m.mut.RLock()
	defer m.mut.RUnlock()
	return m.calls["{{.Name}}"]
}
{{end}}
