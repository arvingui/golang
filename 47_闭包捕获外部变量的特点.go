package main

import (
	"fmt"
)

func main() {
	a := 10
	str := "mike"

	func() {
		//闭包以引用方式捕获外部变量
		a = 666
		str = "go"
		//内部修改外部变量会发生什么呢，内部和外部的打印是否一样，答案是是一样的
		fmt.Printf("内部: a = %d. str = %s\n", a, str)
	}() //（）代表直接调用

	fmt.Printf("外部：a = %d, str = %s\n", a, str)
}
