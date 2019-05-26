package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Printf("内部: a = %d, b = %d\n", a, b)
	}(a, b) //这是已经把此时的a、b值传过去了，只是没有调用而已

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)
}

func main01() {
	a := 10
	b := 20

	defer func() {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}() //()代表调用此匿名函数

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %c", a, b)
}
