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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Params struct {
	Params []interface{}
	Error  []string
}

func NewEncodingError(err error) *Params {
	return &Params{
		Error: []string{"encoding error: " + err.Error()},
	}
}

func NewDecodingError(name string, input []byte, err error) *Params {
	return &Params{
		Error: []string{"calling " + name + " gave decoding error: " + err.Error() + " given input " + string(input)},
	}
}

func Encode(params *Params) []byte {
	data, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	return data
}

func AddStackError(params *Params, funName string) *Params {
	params.Error = append(params.Error, funName)
	return params
}

func ToValues(params *Params) []reflect.Value {
	values := make([]reflect.Value, len(params.Params))
	for i := range params.Params {
		values[i] = reflect.ValueOf(params.Params[i])
	}
	return values
}

type Error interface {
	Error() string
}

var errTyp = reflect.TypeOf((*Error)(nil)).Elem()

func ToParams(values []reflect.Value) *Params {
	params := &Params{Params: make([]interface{}, 0, len(values))}
	for _, v := range values {
		i := v.Interface()
		if v.Type().Implements(errTyp) {
			if i != nil {
				params.Error = []string{i.(Error).Error()}
				return params
			}
		} else {
			params.Params = append(params.Params, i)
		}
	}
	return params
}

func Args(values ...interface{}) *Params {
	params := &Params{Params: make([]interface{}, 0, len(values))}
	for _, v := range values {
		params.Params = append(params.Params, v)
	}
	return params
}

func serve(addr string, name string, fun interface{}) error {
	funValue := reflect.ValueOf(fun)
	http.HandleFunc("/"+name+"/", func(w http.ResponseWriter, r *http.Request) {
		paramStr := r.FormValue("params")
		params := &Params{}
		if err := json.Unmarshal([]byte(paramStr), params); err != nil {
			w.Write(Encode(NewEncodingError(err)))
			return
		}
		if len(params.Error) > 0 {
			w.Write(Encode(AddStackError(params, name)))
			return
		}
		w.Write(Encode(ToParams(funValue.Call(ToValues(params)))))
		return
	})
	println("serving " + name + " at " + addr + "/" + name + " ...")
	return http.ListenAndServe(addr, nil)
}

func log(s interface{}) {
	fmt.Printf("%v\n", s)
}

func inc(i float64) float64 {
	return i + 1
}

func add(i, j float64) (float64, error) {
	if i > 10 {
		return 0, fmt.Errorf("wtf")
	}
	return i + j, nil
}

type StackErr struct {
	errs []string
}

func (this *StackErr) Error() string {
	return strings.Join(this.errs, "\n")
}

func mapper(name string, list interface{}) (interface{}, error) {
	listValue := reflect.ValueOf(list)
	l := listValue.Len()
	results := make([]interface{}, 0, l)
	for i := 0; i < l; i++ {
		params := call(name, Args(listValue.Index(i).Interface()))
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

var funcMap = map[string]string{
	"log": "127.0.0.1:8080",
	"add": "127.0.0.1:8081",
	"map": "127.0.0.1:8082",
	"inc": "127.0.0.1:8083",
}

func call(name string, params *Params) *Params {
	addr := funcMap[name]
	enc := Encode(params)
	values := url.Values{}
	values.Add("params", string(enc))
	getUrl := "http://" + addr + "/" + name + "/?" + values.Encode()
	fmt.Printf("getting %s\n", getUrl)
	resp, err := http.Get(getUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	outs := &Params{}
	if err := json.Unmarshal(body, outs); err != nil {
		return NewDecodingError(getUrl, body, err)
	}
	return outs
}

func main() {
	go func() {
		if err := serve(funcMap["log"], "log", log); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := serve(funcMap["add"], "add", add); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := serve(funcMap["inc"], "inc", inc); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := serve(funcMap["map"], "map", mapper); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	//addout := call("add", Args(int64(11), int64(2)))
	addout := call("map", Args("inc", []float64{1, 2}))
	if len(addout.Error) > 0 {
		fmt.Printf("Error: %#v\n", addout.Error)
		return
	}
	outs := call("log", addout)
	fmt.Printf("output %#v\n", outs)
}
