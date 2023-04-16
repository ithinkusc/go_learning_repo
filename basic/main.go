package main

import (
	"fmt"

	"rsc.io/quote"
)

func foo() {
	fmt.Println("foo is running")
}

func main() {
	fmt.Println(quote.Go())
}
