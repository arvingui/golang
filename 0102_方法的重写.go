package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别，字符类型
	age  int    //年龄
}

func (tmp *Person) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生，继承Person字段，成员和方法都继承了
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

//Student也实现了一个方法，这个方法和Person方法同名，这种方法叫重写
func (tmp *Student) PrintInfo() {
	fmt.Println("Student: tmp = ", tmp)
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "beijing"}
	s.PrintInfo() //问题来了，这里到底调用的是Person,还是Student
	//答案是就近原则：先找本作用域的方法，找不到再用继承的方法

	//如果就调用继承的方法怎么办呢，就用下面的显示调用继承的方法
	s.Person.PrintInfo()
}
