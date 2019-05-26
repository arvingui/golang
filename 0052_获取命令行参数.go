package main

import "fmt"
import "os"

func main() {
	//接受用户参数传递的参数，都是以字符串方式传递
	list := os.Args

	n := len(list)
	fmt.Println("n = ", n)

	//执行go run 44_获取命令行参数.go 1 2
	//结果是n = 3
	//44_获取命令行参数.go 本身也算一个参数

	//打印命令行各个参数
	for i := 0; i < n; i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}

	//用range
	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}
}
