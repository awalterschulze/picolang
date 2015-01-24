
package main

import (
	"github.com/awalterschulze/picolang/fun"
	"flag"
	"os"
	
	"github.com/awalterschulze/picolang/funcs"
	

)

var port = flag.String("port", "8080", "")
var ip = flag.String("ip", "", "")

var myfuncs = map[string]interface{}{
"double":funcs.Double,"inc":funcs.Inc,"log":funcs.Log,"map":funcs.Map,
}

func main() {
	name := os.Args[1]
	addr := *ip + ":" + *port
	
	fun.Register("double", ":8080")
	
	fun.Register("inc", ":8081")
	
	fun.Register("log", ":8082")
	
	fun.Register("map", ":8083")
	
	if err := fun.Serve(addr, name, myfuncs[name]); err != nil {
		panic(err)
	}
}
