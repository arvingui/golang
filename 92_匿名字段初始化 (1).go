package main

type Person struct {
	name string //名字
	set  byte   //性别
	age  int    //年龄
}

type Student struct {
	Person //只有类别，没有名字，匿名字段，继承了Person成员
	id     int
	addr   string
}

func main() {

}
