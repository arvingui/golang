package main

import "fmt"

type mystr string //自定义类型，给一个类型改名
type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //结构体匿名字段
	int    //基础类型匿名字段
	mystr
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "hehehe"}
	fmt.Printf("s = %+v\n", s)

	s.Person = Person{"go", 'm', 22}
	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr) //s.int,没有名字可以这样写

	fmt.Println(s.Person, s.int, s.mystr)
}
