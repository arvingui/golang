package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Printf("请输入一个整数:")

	//fmt.Scanf("%d", &a)
	fmt.Scan(&a) //这样也可以
	fmt.Println("a = ", a)

}
