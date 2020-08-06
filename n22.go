package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Printf("%v, %T\n",  runtime.GOOS, runtime.GOOS)

	a := 1
	b := 2

	fmt.Println( a == b )
	fmt.Println( a <= b )

	fmt.Println( a <= b )
	fmt.Println( a >= b )
	fmt.Println( a < b )

	fmt.Println( a > b )
}
