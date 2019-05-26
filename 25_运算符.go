package main

import (
	"fmt"
)

func main() {
	//关系运算符
	//0是假，非0就是真
	fmt.Println("4 > 3的结果是：", 4 > 3)
	fmt.Println("4 != 3的结果是：", 4 != 3)

	//逻辑运算符 非! 与&& 或||
	fmt.Println("!(4 > 3)的结果是：", !(4 > 3))
	fmt.Println("!(4 != 3)的结果是：", !(4 != 3))

	fmt.Println("true && true的结果是：", true && true)
	fmt.Println("false || true的结果是：", false || true)

	//关系运算符通常要与逻辑运算符配合使用
	a := 8
	//fmt.Println("0 <= a <= 10的结果是：", 0 <= a <= 10)//不支持这种写法
	fmt.Println("a >= 0 && a <= 10的结果是：", a >= 0 && a <= 10)
}
