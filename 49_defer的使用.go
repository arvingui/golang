package main

import (
	"fmt"
)

func main() {

	//defer延迟调用，main函数结束前调用
	//defer只能使用在函数内部
	defer fmt.Println("bbbbbbbbbb")

	fmt.Println("aaaaaaaaaaaaaaa")
}
