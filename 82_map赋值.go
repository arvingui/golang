package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "yoyo"}
	fmt.Println("m1 = ", m1)

	//赋值，如果已经存在的key值，则修改内容
	m1[1] = "c++"
	fmt.Println("m1 = ", m1)

	//赋值，如果不存在则追加，且自动后面追加内存，和append类似
	m1[3] = "go"
	fmt.Println("m1 = ", m1)
}
