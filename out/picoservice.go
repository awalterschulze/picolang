
package main

import (
	"github.com/awalterschulze/picolang/fun"
	"flag"
	"os"
	
	"github.com/awalterschulze/picolang/funcs"
	

)

var port = flag.String("port", "8080", "")
var ip = flag.String("ip", "127.0.0.1", "")

var myfuncs = map[string]interface{}{
"inc":funcs.Inc,"log":funcs.Log,"map":funcs.Map,
}

func main() {
	name := os.Args[0]
	addr := *ip + ":" + *port
	
	fun.Register("inc", "127.0.0.1:8080")
	
	fun.Register("log", "127.0.0.1:8080")
	
	fun.Register("map", "127.0.0.1:8080")
	
	if err := fun.Serve(addr, name, myfuncs[name]); err != nil {
		panic(err)
	}
}
