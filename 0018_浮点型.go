package main

import (
	"fmt"
)

func main() {
	var f1 float32
	f1 = 3.14
	fmt.Println("f1 = ", f1)
	fmt.Printf("f1 type is %T\n", f1)

	f2 := 3.14
	fmt.Printf("f2 type is %T\n", f2) //float64存储小数比float32更准
}
·