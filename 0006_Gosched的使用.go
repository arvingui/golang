package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {            //go+创建匿名函数，不要忘记最后的()
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()
	
	for i := 0; i < 2; i++ {
		//让出时间片，先让别的协程执行，再来执行此协程
		runtime.Gosched()
		fmt.Println("hello")
	}
}