package {{.Package}}

import (
        "fmt"
{{range $path, $name := .Imports}}
        {{$name}} "{{$path}}"{{end}}
)

// {{.Name}} mock
type {{.Name}} struct {
        {{range .Methods}}
        {{.Name}}Func func({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}{{$type}}{{end}}) ({{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}{{$type}}{{end}})
        {{end}}
}

{{$gen := .}}
{{range .Methods}}
// {{.Name}} mocked method
func (m *{{$gen.Name}}) {{.Name}}({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}} {{$type}}{{end}}) ({{range $index, $type := .ReturnTypes}}{{if $index}}, {{end}}{{$type}}{{end}}) {
        if m.{{.Name}}Func == nil {
                panic("unexpected call to mocked method {{.Name}}")
        }
        {{if .ReturnTypes}}return {{end}} m.{{.Name}}Func({{range $index, $type := .ParamTypes}}{{if $index}}, {{end}}p{{$index}}{{end}})
}
{{end}}
