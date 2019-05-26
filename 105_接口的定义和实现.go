/*
接口的定义
type Humaner interface {
	SayHi()
}

1、接口命名习惯以er结尾
2、接口只有方法声明，没有实现，没有数据字段
3、接口可以匿名嵌入其他接口，或嵌入到结构中

尽管go 语言中没有封装、继承、多态这些概念，但同样通过别的方式实现这些特性
封装：通过方法实现
继承：通过匿名字段实现
多态：通过接口实现
*/
package main

import "fmt"

type Humaner interface {
	//方法，只有声明，没有实现，由别的类型（自定义类型）实现
	sayhi() //如果写sayhi() string则表示sayhi()方法的返回值是字符串
}

type Student struct {
	name string
	id   int
}

//Student实现了此方法
func (tmp *Student) sayhi() { //注意：这边接收者一般都是指针传递，可以修改，否则普通值传递则不行
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

//Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %s] sayhi\n", tmp.addr, tmp.group)
}

type MyStr string

//MyStr实现了此方法
func (tmp *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi\n", *tmp)
}

func main() {
	//定义接口类型的变量
	var i Humaner

	//只要实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值
	s := &Student{"mike", 666}
	i = s
	i.sayhi()

	t := &Teacher{"beijing", "go"}
	i = t
	i.sayhi()

	var str MyStr = "hello mike"
	i = &str
	i.sayhi()
}
