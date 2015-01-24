package gen

import (
	"io"
	"strconv"
	"text/template"
)

var makeTemplateString = `
build:
	docker build -t pico-docker .

run:
	{{range .}}docker rm pico-{{.Name}}-container | true
	docker run -d -t -p {{.Port}}:8080 --name pico-{{.Name}}-container pico-docker picoservice {{.Name}}
	{{end}}docker ps

stop:
	{{range .}}docker kill pico-{{.Name}}-container
	{{end}}docker ps

boot2docker-start:
	boot2docker init
	boot2docker start
	$(boot2docker shellinit)

`

var makeTemplate = template.Must(template.New("makeTemplate").Parse(makeTemplateString))

type s struct {
	Name string
	Port string
}

func Makefile(names []string, w io.Writer) error {
	vars := make([]s, len(names))
	for i, name := range names {
		vars[i] = s{name, strconv.Itoa(8080 + i)}
	}
	return makeTemplate.Execute(w, vars)
}