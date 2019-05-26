package main

import (
	"fmt"
)

func myfunc01() int {
	return 666
}

//给返回值起一个变量名，go推荐写法
func myfunc02() (result int) {
	return 666
}

//给返回值起一个变量名，go推荐写法
//常用写法
func myfunc03() (result int) {
	result = 666
	return
}

func main() {
	//无参返回值函数调用
	var a int
	a = myfunc01()
	fmt.Println("a = ", a)

	b := myfunc01()
	fmt.Println("b = ", b)

	c := myfunc02()
	fmt.Println("c = ", c)

	d := myfunc03()
	fmt.Println("d = ", d)
}
