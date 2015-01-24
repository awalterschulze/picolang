package main

import (
	"fmt"
	"github.com/awalterschulze/picolang/fun"
	"reflect"
)

func main() {
	fun.Register("double", "192.168.59.103:8080")
	fun.Register("inc", "192.168.59.103:8081")
	fun.Register("log", "192.168.59.103:8082")
	fun.Register("map", "192.168.59.103:8083")
	doubled, err := fun.Call("double", float64(4))
	if err != nil {
		panic(err)
	}
	fmt.Printf("doubled %#v\n", doubled)
	inced, err := fun.Call("inc", float64(1))
	if err != nil {
		panic(err)
	}
	fmt.Printf("inced %#v\n", inced)
	addout, err := fun.Call("map", "inc", []float64{1, 2})
	if err != nil {
		panic(err)
	}
	fmt.Printf("addOut %#v %#v\n", addout, err)
	outs, err := fun.Call("log", addout...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("output %#v %#v\n", outs, err)
	expected := []interface{}{float64(2), float64(3)}
	if !reflect.DeepEqual(outs, expected) {
		panic("wrong results after log")
	}
	if !reflect.DeepEqual(addout, expected) {
		panic("wrong results before log")
	}
}
