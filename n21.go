package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	n := 42

	fmt.Printf("%d \n %b \n %#x\n", n, n, n)
}
