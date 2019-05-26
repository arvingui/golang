package main

import (
	"fmt"
)

func main() {

	//不同类型变量的声明（定义）
	//var a int
	//var b float64  传统定义方法

	var (
		a int
		b float64
	)

	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//const i int = 10
	//const j float64 = 3.14
	const (
		i int     = 10
		j float64 = 3.14
	)

	/*const (
		i      = 10        //也可以自动推导类型
		j  = 3.14
	)*/

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
}
