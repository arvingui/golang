/*
err接口是做异常处理，很常用,但是通常处理普通错误，但如果遇到一些数组访问越界、空指针引用等，
这些运行时错误会引起panic异常，panic一般作为报告致命错误的一种方式。
一般而言，当panic异常发生时，程序会中断运行，并立即执行goroutine
*/
package main

import "fmt"

func test1() {
	fmt.Println("11111111111111111111")
}

func test2() {
	fmt.Println("2222222222222222222")
	//显示调用panic函数，导致程序中断
	panic("this is a panic test")
}

func test3() {
	fmt.Println("3333333333333333333")
}

func main() {
	test1()
	test2()
	test3()
}
