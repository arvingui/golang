package main

import (
	"fmt"
)

func main() {
	//range默认返回2个值，一个是元素的位置，一个是该元素位置上的值
	//传统写法
	str := "abc"

	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	//迭代打印
	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}

	//for i, _ : range str    这样也行，更直观
	for i := range str {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
}
