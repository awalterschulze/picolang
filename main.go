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

package main

import (
	"bytes"
	"github.com/awalterschulze/picolang/gen"
	"github.com/awalterschulze/picolang/lang"
	"github.com/awalterschulze/picolang/lang/ast"
	"io/ioutil"
	"sort"
)

var langStr = `
import map "github.com/awalterschulze/picolang/funcs.Map"
		import inc "github.com/awalterschulze/picolang/funcs.Inc"

	let a = inc(2)

	println(a)

	let c = map(inc, {2.3, 4})

	println(c)
`

func main() {
	statements, err := lang.Parse(langStr)
	if err != nil {
		panic(err)
	}
	myfuncs := map[string]string{}
	notimports := []ast.Statement{}
	imports := []*ast.Import{}
	for _, statement := range statements {
		if imp, ok := statement.(*ast.Import); ok {
			imports = append(imports, imp)
			myfuncs[imp.Name] = imp.Path
		} else {
			notimports = append(notimports, statement)
		}
	}
	for _, notimport := range notimports {
		notimport.SetImports(imports)
	}
	gobuf := bytes.NewBuffer(nil)
	if err := gen.Go(myfuncs, gobuf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./out/picoservice/picoservice.go", gobuf.Bytes(), 0777); err != nil {
		panic(err)
	}
	names := make([]string, 0, len(myfuncs))
	for name, _ := range myfuncs {
		names = append(names, name)
	}
	sort.Strings(names)
	makebuf := bytes.NewBuffer(nil)
	if err := gen.Makefile(names, makebuf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./out/Makefile", makebuf.Bytes(), 0777); err != nil {
		panic(err)
	}
	mainbuf := bytes.NewBuffer(nil)
	if err := gen.Mainfile("192.168.59.103", names, notimports, mainbuf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./out/main.go", mainbuf.Bytes(), 0777); err != nil {
		panic(err)
	}
}
