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

package funcs

import (
	"fmt"
	"github.com/awalterschulze/picolang/fun"
	"reflect"
	"strings"
)

type StackErr struct {
	errs []string
}

func (this *StackErr) Error() string {
	return strings.Join(this.errs, "\n")
}

func Log(s interface{}) {
	fmt.Printf("%v\n", s)
}

func Inc(i float64) float64 {
	return i + 1
}

func Add(i, j float64) (float64, error) {
	if i > 10 {
		return 0, fmt.Errorf("wtf")
	}
	return i + j, nil
}

func Map(name string, list interface{}) (interface{}, error) {
	listValue := reflect.ValueOf(list)
	l := listValue.Len()
	results := make([]interface{}, 0, l)
	for i := 0; i < l; i++ {
		params := fun.Call(name, fun.Args(listValue.Index(i).Interface()))
		if len(params.Error) > 0 {
			return nil, &StackErr{params.Error}
		}
		if len(params.Params) > 1 {
			panic("not implemented")
		}
		if len(params.Params) == 1 {
			results = append(results, params.Params[0])
		}
	}
	return results, nil
}
