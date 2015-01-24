
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
"add":funcs.Add,"inc":funcs.Inc,"map":funcs.Map,"reduce":funcs.Reduce,
}

func main() {
	name := os.Args[1]
	addr := *ip + ":" + *port
	
	fun.Register("add", "192.168.59.103:8080")
	
	fun.Register("inc", "192.168.59.103:8081")
	
	fun.Register("map", "192.168.59.103:8082")
	
	fun.Register("reduce", "192.168.59.103:8083")
	
	if err := fun.Serve(addr, name, myfuncs[name]); err != nil {
		panic(err)
	}
}
