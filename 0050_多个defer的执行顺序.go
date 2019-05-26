package main

import (
	"fmt"
)

func test(x int) {
	result := 100 / x

	fmt.Println("result = ", result)
}

func main() {
	defer fmt.Println("aaaaaaaaaaaaaa")
	defer fmt.Println("bbbbbbbbbbbbbb")

	//调用一个函数，导致内存出问题
	defer test(0)

	defer fmt.Println("cccccccccccccc")
	//运行顺序cccc、bbbbb、aaaa、test（）报错信息
	//defer执行顺序，写在前面的defer后执行
	//如果前面有一个defer定义的语句/函数出现错误，这些调用依旧会执行
}
