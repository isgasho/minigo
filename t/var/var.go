package main

import "fmt"

var a = b
var b = 1

var x = iota
var iota = 2

var sum = b + 2

func main() {
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", x)
	fmt.Printf("%d\n", sum)
}