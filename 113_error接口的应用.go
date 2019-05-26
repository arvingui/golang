/*
做异常处理，很常用,但是通常处理普通错误，但如果遇到一些数组访问越界、空指针引用等，
这些运行时错误会引起panic异常
*/
package main

import "fmt"
import "errors"

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return
}

func main() {
	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("result = ", result)
	}
}
