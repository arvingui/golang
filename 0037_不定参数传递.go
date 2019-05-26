package main

import (
	"fmt"
)

func MyFunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Printf("data = %d\n", data)
	}
}

func MyFunc2(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}

func test(args ...int) {
	//全部元素传递给mysfunc
	//MyFunc(args...)

	//只想把后2个参数传递给另外一个函数使用
	MyFunc2(args[:2]...) //从args[0]到args[2]（不包括args[2]）传递过去
	//MyFunc2(args[2:]...) //从args[2]开始（包括本身），把后面所有元素传递过去
}

func main() {
	test(1, 2, 3, 4)
}
