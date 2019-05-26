//go语言可见性规则
//如果想使用别的包的函数、结构体类型、结构体成员
//函数名，类型名，结构体成员变量名，首字母必须大写，可见
//如果首字母是小写，只能在同一个包里使用
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

//值传递，这样无法达到效果
func test01(s Student) {
	s.id = 666
	fmt.Println("test01: ", s)
}

func main01() {
	s := Student{1, "mike", 'm', 18, "beijing"}

	test01(s) //值传递，这样形参是无法改变实参的
	fmt.Println("main: ", s)
}

//地址传递，这样可以达到效果
func test02(p *Student) {
	p.id = 666
}

func main() {
	s2 := Student{1, "mike", 'm', 18, "beijing"}

	test02(&s2) //地址传递，这样形参是可以改变实参的
	fmt.Println("main: ", s2)
}
