package main

import (
	"fmt"
)

func main() {
	var flag bool

	flag = false
	fmt.Printf("flag = %d\n", flag) //输出是这个：flag = %!d(bool=false)

	fmt.Printf("flag = %t\n", flag) //输出：flag = flase

	//fmt.Printf("flag = %d\n", int(flag)) //会报错
	//flag = (bool)1	//也会报错，因为bool类型与int型不能相互转化，这种类型叫做不兼容类型

	var ch byte
	ch = 'a'

	var t int
	t = int(ch) //强制类型转换，但是括号括的不是int，这个跟c语言不一样，注意
	fmt.Println("t = ", t)
}
