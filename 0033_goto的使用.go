package main

import (
	"fmt"
)

func main() {
	//break, continue只能用在循环体中
	//而goto可以用在任何地方，但是不能跨函数使用
	fmt.Println("1111111111")

	goto End //goto是关键字，End是用户起的标签

	fmt.Println("22222222222")

End:
	fmt.Println("33333333333")
}
