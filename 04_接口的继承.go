package main

import (
	"fmt"
)
//定义接口的类型
type Humaner interface {  //子集
	//方法，只有声明，没有实现，由别的<自定义类型>实现
	sayhi()
}

type Personer interface {  //超集
	Humaner  //匿名字段，继承了sayhi()
	sing(lrc string)
}

type Student struct {
	name string
	id    int
}

//Strudent实现了sayhi()

func (tmp *Student) sayhi() {
	fmt.Printf("Strudent[%s, %d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Student)sing(lrc string) {
	fmt.Println("Strudent在唱着: ", lrc)
}

func main() {
	//定义一个接口类型的变量
	var i Personer
	s := &Student{"mike", 666}
	i = s
	
	i.sayhi()   //继承过来的方法
	i.sing("学生歌")
}


