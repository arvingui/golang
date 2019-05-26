package main

import (
	"fmt"
)

func main() {
	//变量的声明需要关键字var
	//常量的声明需要关键字const
	const a int = 10

	//a = 20     会报错，const常量不允许再次被修改
	fmt.Println("a = ", a)

	const b = 10 //const也可以支持自动推导类型
	//const b := 10   注意此时用:=是报错的， 常量的自动推导不需要:=
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)
}
