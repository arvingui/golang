package main

import (
	"fmt"
)

func test01() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

func test02(i int) int {
	if i == 1 {
		return 1
	}
	return i + test02(i-1)
}

func main() {
	sum := test01()
	fmt.Println("sum = ", sum)

	sum02 := test02(100)
	fmt.Println("sum02 = ", sum02)
}
