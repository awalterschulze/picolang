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
	"github.com/awalterschulze/picolang/fun"
	"log"
)

func Log(s ...interface{}) interface{} {
	log.Printf("%v\n", s)
	return s
}

func Inc(i float64) float64 {
	return i + 1
}

func Double(i float64) float64 {
	return i * 2
}

func Add(i, j float64) float64 {
	return i + j
}

func Map(name string, list []interface{}) ([]interface{}, error) {
	results := make([]interface{}, 0, len(list))
	for i := 0; i < len(list); i++ {
		values, err := fun.Call(name, list[i])
		if err != nil {
			return nil, err
		}
		if len(values) > 1 {
			panic("not implemented")
		}
		if len(values) == 1 {
			results = append(results, values[0])
		}
	}
	return results, nil
}

func Reduce(name string, initial interface{}, list []interface{}) (interface{}, error) {
	result := initial
	for i := 0; i < len(list); i++ {
		res, err := fun.Call(name, result, list[i])
		if err != nil {
			return nil, err
		}
		if len(res) == 1 {
			result = res[0]
		} else {
			panic("not implemented")
		}
	}
	return result, nil
}
