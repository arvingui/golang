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
	s1 := Student{1, "mike", 'm', 18, "beijing"}
	s2 := Student{1, "mike", 'm', 18, "beijing"}
	s3 := Student{2, "mike", 'm', 18, "beijing"}
	fmt.Println("s1 == s2 ", s1 == s2)
	fmt.Println("s1 == s3 ", s1 == s3)

	//t同类型的2个结构体变量可以相互赋值
	var tmp Student
	tmp = s3
	fmt.Println("tmp = ", tmp)
}
