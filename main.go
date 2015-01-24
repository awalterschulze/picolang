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
	"flag"
	"github.com/awalterschulze/picolang/compile"
	"io/ioutil"
	"log"
)

var ip = flag.String("ip", "127.0.0.1", "ip where dockers will be deployed")
var out = flag.String("o", "./out/", "folder where Dockerfile is and generated code should code")

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	compile.Compile(string(data), *out, *ip)
}
