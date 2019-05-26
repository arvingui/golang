package main

import (
	"fmt"
)

func main() {
	var a bool
	fmt.Println("a = ", a)

	a = true
	fmt.Println("a = ", a)

	c := false
	fmt.Printf("c type = %T", c)
}
