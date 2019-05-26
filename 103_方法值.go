package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, p)
}

func main() {
	p := Person{"mike", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)

	p.SetInfoPointer() //传统调用方式

	//保存方法入口地址
	pFunc := p.SetInfoPointer //这个就是方法值，无需再传递接收，隐藏了接收者，注意后面不要加括号了
	pFunc()                   //等价于p.SetInfoPointer()

	vFunc := p.SetInfoValue
	vFunc() //等价于p.SetInfoValue（）
}
