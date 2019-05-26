package main

import (
	"fmt"
)

func test(a int) {
	if a == 1 {
		fmt.Println("a = ", a)
		return //终止函数调用
	}

	//调用函数自身
	test(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	test(3)
	fmt.Println("main")
}
