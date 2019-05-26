package main

import "fmt"

func test1() {
	fmt.Println("11111111111111111111")
}

func test2(x int) {
	var a [10]int
	a[x] = 111 //当x为20的时候，导致数组越界，产生一个panic，导致程序崩溃
}

func test3() {
	fmt.Println("3333333333333333333")
}

func main() {
	test1()
	test2(20)
	test3()
}
