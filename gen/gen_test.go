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
	"bytes"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"testing"
)

func TestGen(t *testing.T) {
	funcs := map[string]string{
		"inc": "github.com/awalterschulze/picolang/funcs.Inc",
		"map": "github.com/awalterschulze/picolang/funcs.Map",
	}
	buf := bytes.NewBuffer(nil)
	err := Go(funcs, "127.0.0.1", buf)
	if err != nil {
		t.Error(err)
	}
	data := buf.Bytes()
	t.Logf("%v", (string(data)))
	os.Mkdir("./testdata/", 0777)
	if err := ioutil.WriteFile("./testdata/main.go", data, 0777); err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "./testdata/main.go", nil, parser.ImportsOnly); err != nil {
		t.Fatal(err)
	}
}
