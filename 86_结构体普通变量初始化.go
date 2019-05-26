package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	//按顺序初始化，每个成员必须初始化
	var s1 Student = Student{1, "mike", 'm', 18, "beijing"}
	fmt.Println("s1 = ", s1) //'m'打印出来是对应的ascii码

	//指定成员初始化，没有指定初始化的成员默认为0
	s2 := Student{name: "mike", addr: "beijing"}
	fmt.Println("s2 = ", s2)
}
