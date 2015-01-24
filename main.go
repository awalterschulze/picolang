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
	"io/ioutil"
	"sort"
)

func main() {
	myfuncs := map[string]string{
		"inc": "github.com/awalterschulze/picolang/funcs.Inc",
		"map": "github.com/awalterschulze/picolang/funcs.Map",
		"log": "github.com/awalterschulze/picolang/funcs.Log",
	}
	gobuf := bytes.NewBuffer(nil)
	if err := gen.Go(myfuncs, gobuf); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./out/picoservice.go", gobuf.Bytes(), 0777); err != nil {
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
}
