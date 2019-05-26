package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := a[0:3:5]
	fmt.Println("a = ", a)
	fmt.Println("len(a) = ", len(a)) //长度3-0
	fmt.Println("cap(a) = ", cap(a)) //容量5-0

	s = a[1:4:5]
	fmt.Println("s = ", s)           //从下标1开始，取4-1=3个元素
	fmt.Println("len(s) = ", len(s)) //长度4-1
	fmt.Println("cap(s) = ", cap(s)) //容量5-1
}
