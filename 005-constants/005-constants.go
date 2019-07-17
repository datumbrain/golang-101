package main

import "fmt"

const (
	a uint8 = 132
	b float32 = 42.78
	c string = "Fahad"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
