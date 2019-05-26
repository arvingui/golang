package main

import (
	"fmt"
)

func main() {
	//注意1：iota是在给常量赋值时使用的
	const (
		a = iota // a = 0
		b = iota // b = 1
		c = iota // c = 2
	)

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	//注意2：iota遇到const就会重置为0
	const d = iota
	fmt.Println("d = ", d) // d = 0

	//注意3：可以只写一个iota
	const (
		e = iota // e = 0
		f        // f = 1
		g        // g = 2
	)
	fmt.Printf("e = %d, f = %d, g = %d\n", e, f, g)

	//注意4：如果是同一行，值都一样
	const (
		a1         = iota             // a1 = 0
		a2, a3, a4 = iota, iota, iota // a2 = a3 = a4 = 1
		a5         = iota             // a5 = 2
	)
	fmt.Printf("a1 = %d, a2 = %d, a3 = %d, a4 = %d, a5 = %d\n", a1, a2, a3, a4, a5)
}
