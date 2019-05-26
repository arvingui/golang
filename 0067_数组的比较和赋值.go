package main

import "fmt"

func main() {
	//支持同类型数组比较与赋值，但比较只支持==或！=，比较的是每个元素和数字类型
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}
	fmt.Println("a == b", a == b) //true
	fmt.Println("a == c", a == c) //false

	var d [5]int
	d = a
	fmt.Println("d = ", d)
}
