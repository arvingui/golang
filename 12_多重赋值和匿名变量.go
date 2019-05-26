package main

import (
	"fmt"
)

func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	//a := 10
	//b := 20

	a, b := 10, 20
	//传统方式交换a b 的值
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Printf("a = %d, b = %d\n", a, b)

	i, j := 10, 20
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

	i, j = 10, 20
	//  “_”为匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用才有优势
	//因为go函数可以返回多个值
	temp, _ = i, j
	fmt.Println("temp = ", temp)

	var c, d, e int
	c, d, e = test() //return 1 2 3
	fmt.Printf("c = %d, b = %d, e = %d\n", c, d, e)
	_, d, _ = test()
	fmt.Printf("d = %d\n", d)
}
