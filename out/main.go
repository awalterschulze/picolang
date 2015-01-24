
package main

import (
	"github.com/awalterschulze/picolang/fun"
	"fmt"
)

var _ = fmt.Printf
var _ = fun.Call

func main() {
	fun.Register("inc", "192.168.59.103:8080")
	fun.Register("map", "192.168.59.103:8081")
	
	a, err := fun.Call("inc",2.000000)
if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n",a)
	c, err := fun.Call("map","inc", []interface{}{2.300000, 4.000000})
if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n",c)
	
}

