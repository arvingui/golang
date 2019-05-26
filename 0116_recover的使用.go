/*
运行时panic异常一旦被引发就会导致程序崩溃，不过Go语言为我们提供了专用于“拦截”运行时panic的内建函
数--recover，他可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。
func recover() interface{}

注意：recover只有在defer调用的函数中有效。
*/
package main

import "fmt"

func test1() {
	fmt.Println("11111111111111111111")
}

func test2(x int) {
	//设置recover，这样当下面数组越界造成程序崩溃时，recover可以恢复程序正常运行
	defer func() {
		//recover()  //只写这行只能恢复程序，但是看不到错误信息，下面可以恢复并打印
		//fmt.Println(recover())//但这样有一个问题，就是当test2函数没有数组越界是也不能正常执行test2

		//所以要做个判断
		if err := recover(); err != nil { //产生了panic异常
			fmt.Println(err)
		}

	}() //别忘了（），调用此匿名函数

	var a [10]int
	a[x] = 111 //当x为20的时候，导致数组越界，产生一个panic，导致程序崩溃
}

func test3() {
	fmt.Println("3333333333333333333")
}

func main() {
	test1()
	test2(20)
	test3()
}
