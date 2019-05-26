//package hello//错误，同一级目录包名必须一样
package main

import "fmt"

func test() {
	fmt.Println("this is a test func")
}

//要设置环境变量，变量名：GOPATH  值：src的绝对路径
