//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package gen

import (
	"io"
	"strconv"
	"text/template"
)

var makeTemplateString = `
build:
	docker build -t pico-docker .

delete:
	{{range .}}docker rm pico-{{.Name}}-container || true
	{{end}}docker ps

run:
	make delete
	{{range .}}docker run -d -t -p {{.Port}}:8080 --name pico-{{.Name}}-container pico-docker picoservice {{.Name}}
	{{end}}docker ps
	go run main.go
	make stop

stop:
	{{range .}}docker kill pico-{{.Name}}-container || true	
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
