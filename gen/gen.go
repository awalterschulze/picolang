package gen

import (
	"io"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var goTemplateString = `
package main

import (
	"github.com/awalterschulze/picolang/fun"
	"flag"
	"os"
	{{range .Imports}}
	{{.}}
	{{end}}

)

var port = flag.String("port", "8080", "")
var ip = flag.String("ip", "", "")

var myfuncs = map[string]interface{}{
{{range .Funcs}}"{{.Name}}":{{.Function}},{{end}}
}

func main() {
	name := os.Args[1]
	addr := *ip + ":" + *port
	{{range .Funcs}}
	fun.Register("{{.Name}}", "{{.Addr}}")
	{{end}}
	if err := fun.Serve(addr, name, myfuncs[name]); err != nil {
		panic(err)
	}
}
`

var goTemplate = template.Must(template.New("goTemplate").Parse(goTemplateString))

type fun struct {
	Name     string
	Function string
	Addr     string
	Import   string
}

func Go(funcs map[string]string, ip string, w io.Writer) error {
	names := make([]string, 0, len(funcs))
	for name, _ := range funcs {
		names = append(names, name)
	}
	sort.Strings(names)
	funs := make([]fun, len(names))
	for i, name := range names {
		funs[i] = newFun(name, funcs[name], ip, i)
	}
	imports := uniqImports(funs)
	vars := struct {
		Funcs   []fun
		Imports []string
	}{funs, imports}
	return goTemplate.Execute(w, vars)
}

func uniqImports(funs []fun) []string {
	uniq := make([]string, 0, len(funs))
	set := make(map[string]struct{})
	for _, f := range funs {
		if _, ok := set[f.Import]; ok {
			continue
		}
		set[f.Import] = struct{}{}
		uniq = append(uniq, f.Import)
	}
	sort.Strings(uniq)
	return uniq
}

func newFun(name string, function string, ip string, num int) fun {
	fs := strings.Split(function, ".")
	imp := `"` + strings.Join(fs[:len(fs)-1], ".") + `"`
	calls := strings.Split(function, "/")
	call := calls[len(calls)-1]
	port := strconv.Itoa(8080 + num)
	addr := ip + ":" + port
	return fun{Name: name, Function: call, Import: imp, Addr: addr}
}
