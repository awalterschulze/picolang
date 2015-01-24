
package main

import (
	"github.com/awalterschulze/picolang/fun"
	"fmt"
)

var _ = fmt.Printf
var _ = fun.Call

func main() {
	fun.Register("add", "192.168.59.103:8080")
	fun.Register("inc", "192.168.59.103:8081")
	fun.Register("map", "192.168.59.103:8082")
	fun.Register("reduce", "192.168.59.103:8083")
	
	c, err := fun.Call("reduce","add", 0.000000, []interface{}{1.000000, 2.000000, 3.000000, 4.000000})
if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n",c)
	
}

