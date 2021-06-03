package main

import "fmt"

const (
	a = iota
	b
	c
)

var ()

type Hello interface {
}

func main() {
	fmt.Print("hello")

	var a string
	a = "string a"
	fmt.Print(a, &a)

	a += "string b"
	fmt.Println(a, &a)

	//var m2 = map[string]string{}

	//for k,v =  range m2{
	//
	//}
}
