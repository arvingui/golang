package main

import (
	"fmt"
	"time"
)

//break可用于for、switch、select， 而continue仅能用于for循环
func main() {
	i := 0
	for { //for循环后面不写任何东西，就是死循环
		i++
		time.Sleep(time.Second) //睡眠1s

		if i == 5 {
			//break //跳出最近循环
			continue //仅仅是跳出5这个循环
		}
		fmt.Println("i = ", i)

	}
}
