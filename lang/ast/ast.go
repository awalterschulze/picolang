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

package ast

import (
	"fmt"
	"strconv"
	"strings"
)

type Statements []Statement

type Statement interface {
	Generate() string
	SetImports([]*Import)
}

type Let struct {
	Var  string
	Expr *Expr
}

func (this *Let) Generate() string {
	first := this.Var + ", err := " + this.Expr.Generate() + "\n"
	iffy := `if err != nil {
		panic(err)
	}`
	return first + iffy
}

func (this *Let) SetImports(imps []*Import) {
	this.Expr.SetImports(imps)
}

type Println struct {
	Expr *Expr
}

func (this *Println) Generate() string {
	return "fmt.Printf(\"%+v\\n\"," + this.Expr.Generate() + ")"
}

func (this *Println) SetImports(imps []*Import) {
	this.Expr.SetImports(imps)
}

type Import struct {
	Name string
	Path string
}

func (this *Import) Generate() string {
	paths := strings.Split(this.Path, ".")
	return fmt.Sprintf("\"%s\"", strings.Join(paths[:len(paths)-1], "."))
}

func (this *Import) SetImports([]*Import) {}

type Expr struct {
	Name    string
	Params  []*Param
	Imports []*Import
}

func (this *Expr) Generate() string {
	if this.Params == nil {
		return this.Name
	}
	if len(this.Params) == 0 {
		return "fun.Call(" + strconv.Quote(this.Name) + ")"
	}
	strs := make([]string, len(this.Params))
	for i := range this.Params {
		strs[i] = this.Params[i].Generate()
	}
	return "fun.Call(" + strconv.Quote(this.Name) + "," + strings.Join(strs, ", ") + ")"
}

func (this *Expr) SetImports(imps []*Import) {
	this.Imports = imps
	for _, param := range this.Params {
		param.SetImports(imps)
	}
}

type Param struct {
	Id      *string
	Str     *string
	Float   *float64
	List    []*Param
	Imports []*Import
}

func (this *Param) Generate() string {
	if this.Id != nil {
		for _, imp := range this.Imports {
			if imp.Name == *this.Id {
				return strconv.Quote(*this.Id)
			}
		}
		return *this.Id
	}
	if this.Str != nil {
		return strconv.Quote(*this.Str)
	}
	if this.Float != nil {
		return fmt.Sprintf("%f", *this.Float)
	}
	strs := make([]string, len(this.List))
	for i := range this.List {
		strs[i] = this.List[i].Generate()
	}
	return "[]interface{}{" + strings.Join(strs, ", ") + "}"
}

func (this *Param) SetImports(imps []*Import) {
	this.Imports = imps
	for _, param := range this.List {
		param.SetImports(imps)
	}
}
