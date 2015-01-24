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

package fun

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type Params struct {
	Params   []interface{}
	StackErr []string
}

func (this *Params) Error() string {
	return strings.Join(this.StackErr, "\n")
}

func NewEncodingError(err error) *Params {
	return &Params{
		StackErr: []string{"encoding error: " + err.Error()},
	}
}

func NewDecodingError(name string, input []byte, err error) *Params {
	return &Params{
		StackErr: []string{"calling " + name + " gave decoding error: " + err.Error() + " given input " + string(input)},
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
	params.StackErr = append(params.StackErr, funName)
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
				params.StackErr = []string{i.(Error).Error()}
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

func Serve(addr string, name string, fun interface{}) error {
	funValue := reflect.ValueOf(fun)
	http.HandleFunc("/"+name+"/", func(w http.ResponseWriter, r *http.Request) {
		paramStr := r.FormValue("params")
		params := &Params{}
		if err := json.Unmarshal([]byte(paramStr), params); err != nil {
			w.Write(Encode(NewEncodingError(err)))
			return
		}
		if len(params.StackErr) > 0 {
			w.Write(Encode(AddStackError(params, name)))
			return
		}
		w.Write(Encode(ToParams(funValue.Call(ToValues(params)))))
		return
	})
	println("serving " + name + " at " + addr + "/" + name + " ...")
	return http.ListenAndServe(addr, nil)
}

func Call(name string, params *Params) *Params {
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

func ToInterfaces(params *Params) ([]interface{}, error) {
	if len(params.StackErr) > 0 {
		return params.Params, params
	}
	return params.Params, nil
}

func CallMeMaybe(name string, args ...interface{}) ([]interface{}, error) {
	params := Args(args...)
	outs := Call(name, params)
	return ToInterfaces(outs)
}

var funcMap = map[string]string{}

func Register(name string, addr string) error {
	if _, ok := funcMap[name]; ok {
		return fmt.Errorf("conflicing function names for %s", name)
	}
	funcMap[name] = addr
	return nil
}
