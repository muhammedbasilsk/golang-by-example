package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Printf("%v, %T",  runtime.GOOS, runtime.GOOS)
}
