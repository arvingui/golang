package main

import (
	"fmt"
)

func main() {

	//if支持1个初始化语句，初始化语句和判断条件以分号分隔
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	}
}
