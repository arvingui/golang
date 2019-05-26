package main

import (
	"fmt"
)

func main() {
	//1
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}

	//2
	if a := 11; a == 11 {
		fmt.Println("a == 11")
	} else {
		fmt.Println("a != 11")
	}

	a = 12
	if a == 12 {
		fmt.Println("a == 12")
	} else if a > 12 {
		fmt.Println("a > 10")
	} else if a < 12 {
		fmt.Println("a < 12 ")
	} else {
		fmt.Println("这种情况是不可能的")
	}
}
