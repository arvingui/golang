package main

import (
	"fmt"
)

type FuncType func(int, int) int

//实现加法
func Add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

//回调函数，函数有一个参数是函数类型，这个函数是回调函数
//计算器，可以进行四则运算
//多态，多种形态，调用一个接口，不同的表现，即加减乘除
func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b)
	//result = Add(a, b) //这样就把Calc函数写死了，就是去了多态得意义
	return
}

func main() {
	a := Calc(1, 1, Add)
	fmt.Println("a = ", a)

	a = Calc(10, 5, minus)
	fmt.Println("a = ", a)
}
