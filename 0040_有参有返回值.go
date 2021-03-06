package main

import (
	"fmt"
)

func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return //又返回值得函数，必须通过return返回
}

func main() {
	max, min := MaxAndMin(10, 20)
	fmt.Printf("max = %d, min = %d\n", max, min)

	//通过匿名函数丢弃某个返回值
	a, _ := MaxAndMin(10, 20)
	fmt.Printf("a = %d\n", a)
}
