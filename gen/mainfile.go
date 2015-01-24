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
	"github.com/awalterschulze/picolang/lang/ast"
	"io"
	"strconv"
	"text/template"
)

var mainTemplateString = `
package main

import (
	"github.com/awalterschulze/picolang/fun"
	"fmt"
)

var _ = fmt.Printf
var _ = fun.Call

func main() {
	{{range .Registers}}fun.Register("{{.Name}}", "{{.Addr}}")
	{{end}}
	{{range .Statements}}{{.Generate}}
	{{end}}
}

`

var mainTemplate = template.Must(template.New("mainTemplate").Parse(mainTemplateString))

type temps struct {
	Statements []ast.Statement
	Registers  []register
}

type register struct {
	Name string
	Addr string
}

func Mainfile(ip string, names []string, statements []ast.Statement, w io.Writer) error {
	regs := make([]register, len(names))
	for i, name := range names {
		regs[i] = register{name, ip + ":" + strconv.Itoa(8080+i)}
	}
	return mainTemplate.Execute(w, &temps{statements, regs})
}
