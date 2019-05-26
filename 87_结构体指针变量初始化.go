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
	//按顺序初始化，每个成员必须初始化,别忘记&
	var p1 *Student = &Student{1, "mike", 'm', 18, "beijing"}
	fmt.Println("*p1 = ", *p1) //'m'打印出来是对应的ascii码

	//指定成员初始化，没有指定初始化的成员默认为0
	p2 := &Student{name: "mike", addr: "beijing"}
	fmt.Println("*p2 = ", *p2)

	//也可以打印没有*号，go语言的灵活之处
	p3 := &Student{name: "mike", addr: "beijing"}
	fmt.Println("p3 = ", p3)
	fmt.Printf("p3 type is %T\n", p3)
}
