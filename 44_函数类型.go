package main

import (
	"fmt"
)

func Add(a, b int) int { //函数名Add相当于这个函数得入口地址，所有赋值用fTest=Add即可
	return a + b
}

func minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给一个函数类型起名
//FuncType他是一个函数类型
type FuncType func(int, int) int //没有函数名字，没有{}，只有相同参数类型，相同返回值类型一样就可以赋值
func main() {
	var result int
	result = Add(1, 1) //传统调用方式
	fmt.Println("result = ", result)

	//声明一个函数类型得变量，变量名叫fTest
	var fTest FuncType
	fTest = Add            //是变量就可以赋值
	result = fTest(10, 20) //等价于Add(10, 20)
	fmt.Println("result2 = ", result)

	fTest = minus
	result = fTest(10, 5)
	fmt.Println("result3 = ", result)
}
