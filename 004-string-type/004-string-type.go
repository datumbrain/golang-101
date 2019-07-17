package main

import "fmt"

func main() {
	s := "Hello, world!"

	// raw string
	rawString := `Hey
	this is a new string `

	fmt.Println(rawString)

	bytes := []byte(s)

	for i := 0; i < len(bytes); i++ {
		fmt.Printf("%v", bytes[i])
		if i != (len(bytes) - 1) {
			fmt.Print(", ")
		}
	}

	fmt.Println()

	for i, v := range bytes {
		fmt.Printf("At index position we have %d we have hex %x.\n", i, v)
	}
}
