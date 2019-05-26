package main

import (
	"fmt"
)

func main() {
	a := 10
	str := "mike"

	//匿名函数，没有函数名字
	f1 := func() {
		fmt.Println("a = ", a) //能够捕获到外面的变量
		fmt.Println("str = ", str)
	}

	f1()

	//给一个函数类型起别名
	type FuncType func() //函数没有参数，没有返回值

	//声明变量
	var f2 FuncType
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	//匿名函数定义好后必须调用，不然会报错
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)

	}() //后面圆括号代表使用此匿名函数，是参数列表，里面可以加参数

	//带参数得匿名函数，同时调用
	f3 := func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
	f3(1, 2)

	//定义匿名函数，同时调用
	func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(10, 20)

	//匿名函数，有参数有返回值
	xMax, yMin := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return //有返回值函数必须return
	}(10, 20)

	fmt.Printf("xMax = %d, yMin = %d\n", xMax, yMin)
}
