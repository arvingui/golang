package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id   int
	addr string
	name string //和Person同名了
}

func main() {
	var s Student
	//默认规则：就近原则，先赋值给本作用域
	s.name = "mike" //操作的是Student的name,还是Person的name呢? 答案是Student的name
	//如果就想指定赋值给Person，那就显示调用
	s.Person.name = "yoyo"

	s.sex = 'm'
	s.age = 18
	s.id = 01
	s.addr = "beijing"
	fmt.Printf("s = %+v\n", s)
}
