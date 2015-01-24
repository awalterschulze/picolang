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
	"fmt"
	"github.com/awalterschulze/picolang/fun"
	"github.com/awalterschulze/picolang/funcs"
	"time"
)

func main() {
	logAddr := "127.0.0.1:8080"
	addAddr := "127.0.0.1:8081"
	mapAddr := "127.0.0.1:8082"
	incAddr := "127.0.0.1:8083"
	fun.Register("log", logAddr)
	fun.Register("add", addAddr)
	fun.Register("map", mapAddr)
	fun.Register("inc", incAddr)
	go func() {
		if err := fun.Serve(logAddr, "log", funcs.Log); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := fun.Serve(addAddr, "add", funcs.Add); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := fun.Serve(incAddr, "inc", funcs.Inc); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := fun.Serve(mapAddr, "map", funcs.Map); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)
	//addout, err := fun.Call("add", int64(11), int64(2))
	addout, err := fun.Call("map", "inc", []float64{1, 2})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	outs, err := fun.Call("log", addout...)
	fmt.Printf("output %#v %#v\n", outs, err)
}
