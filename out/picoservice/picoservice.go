
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
"inc":funcs.Inc,"map":funcs.Map,
}

func main() {
	name := os.Args[1]
	addr := *ip + ":" + *port
	
	fun.Register("inc", "192.168.59.103:8080")
	
	fun.Register("map", "192.168.59.103:8081")
	
	if err := fun.Serve(addr, name, myfuncs[name]); err != nil {
		panic(err)
	}
}
