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
	"github.com/awalterschulze/picolang/compile"
)

func main() {
	var langStr = `
import map "github.com/awalterschulze/picolang/funcs.Map"
		import inc "github.com/awalterschulze/picolang/funcs.Inc"

	let a = inc(2)

	println(a)

	let c = map(inc, {2.3, 4})

	println(c)
`
	compile.Compile(langStr, "./out/", "192.168.59.103")
}
