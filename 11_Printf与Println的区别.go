package main

import (
	"fmt"
)

func main() {
	a := 10
	//一段一段处理，自动换行
	fmt.Println("a = ", a)
	//格式化输出，换行需要自己加"\n"
	fmt.Printf("a = %d\n", a)

	b := 20
	c := 30
	fmt.Println("Println: a = ", a, ", b = ", b, ", c = ", c) //这就是所谓的Println是一段一段的处理，先处理a,再b,再c
	fmt.Printf("Printf: a = %d, b = %d, c = %d\n", a, b, c)   //这就是Println与Printf的区别
}
