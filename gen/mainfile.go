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
