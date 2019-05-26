package main

import "fmt"

func main() {
	//go语言中定义时数组的下标必须是常量，哪怕是已经赋值的变量都不行
	var a [10]int
	var b [5]int

	fmt.Printf("len(a) = %d, len(b) = %d\n", len(a), len(b))

	a[0] = 1
	i := 1
	a[i] = 2

	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	for i, data := range a {
		fmt.Printf("a[%d] = %d\n", i, data)
	}
}
