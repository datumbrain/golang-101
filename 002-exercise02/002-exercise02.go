package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true

	fmt.Printf("%v, %v, %v\n", x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T, %T, %T\n", x, y, z)
}