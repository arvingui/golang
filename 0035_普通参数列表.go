package main

import (
	"fmt"
)

//这时就不需要var a int了
func MyFunc01(a int) {
	fmt.Println("a = ", a)
}

//func MyFunc02(a, b int)
func MyFunc02(a int, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func main() {
	//有参无返回值函数的调用
	a := 666
	MyFunc01(a)
	MyFunc02(222, 777)
}
