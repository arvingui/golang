package main

import (
	"fmt"
)

func main() {
	var str string
	str = "abc"
	fmt.Println("str = ", str)

	str2 := "mike"
	fmt.Printf("str2 type is %T\n", str2)

	//内建函数，测字符串长度，有多少个字符
	fmt.Println("len(str2) = ", len(str2))
}
