package main

import (
	"fmt"
	"runtime"
)

func main() {
	specs := fmt.Sprintf( "%v - %v", runtime.GOOS, runtime.GOARCH)
	fmt.Println(specs)
}