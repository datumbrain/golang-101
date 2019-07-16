package _01_hello

import "fmt"

func main() {
	fmt.Println("Hello world!")
	var z int = 10

	fmt.Println(z)
	n := 10
	for i := 0; i < n; i++ {
		fmt.Print(i)
		if i != n-1 {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}

	foo() //
}

func foo() {
	fmt.Println("This is my first (second, actually) function.")
}
