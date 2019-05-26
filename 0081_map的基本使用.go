package main

import "fmt"

//map的键值要是唯一的
func main() {
	//定义一个变量，类型为map[int]string
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	//对于map只有len，没有cap
	fmt.Println("len = ", len(m1))

	//可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	//可以通过make创建，可以指定长度，但这只是指定了容量，实际上
	//里面还没有值，len（）求长度还是会等于0
	m3 := make(map[int]string, 10)
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))

	m4 := make(map[int]string, 2)
	m4[1] = "mike"
	m4[2] = "go"
	m4[3] = "c++" //虽然一开始分配了2个内存，但是后面赋值多了会自动扩容
	//map赋值后是无序的，所以打印结果不一定是123分布

	fmt.Println("m3 = ", m4)
	fmt.Println("len = ", len(m4))

	//直接定义加初始化
	m5 := map[int]string{1: "mike", 2: "go", 3: "c++", 1: "xxx"} //报错，键值要唯一！
	fmt.Println("m5 = ", m5)
	fmt.Println("len = ", len(m5))
}
