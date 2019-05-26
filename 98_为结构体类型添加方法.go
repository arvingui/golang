//注意1：只要接收者类型不一样，这个方法就算同名，也是不同方法，不会出现重复定义函数的错误
//注意2：参数receive类型可以是T也可以是*T，但是基类型T不能是接口或指针
//注意3：不支持重载方法，但是能定义名字相同但是不同参数的方法，同注意1
//参数4：receive可以是基础类型，也可以是结构体类型
package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//带有接收者的函数叫做方法
func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

//通过一个函数，给成员复制
func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}

func main() {
	//定义同时初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体变量
	var p2 Person
	(&p2).SetInfo("yoyo", 'f', 22)
	p2.PrintInfo()
}
