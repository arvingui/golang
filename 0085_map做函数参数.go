package main

import "fmt"

func test(m map[int]string) {
	delete(m, 1) //这里使用的是同一个map
}

func main() {
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}
	fmt.Println("m = ", m)

	test(m)
	fmt.Println("m = ", m)
}
